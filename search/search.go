package search

import (
	"context"
	"github.com/corvus-ch/bilocation/tag"
	"os"
	"path/filepath"
	"syscall"

	"github.com/corvus-ch/logr/log"
	"golang.org/x/sync/errgroup"
)

// Search recursively walks the file tree rooted at root and searches for files matching query.
// The context can be used to to cancel the search early.
// Search walks the filesystem looking for files matching a given query.
func Search(cfg Config) ([]Match, error) {
	q := NewQuery(cfg.Query())

	g, ctx := errgroup.WithContext(context.Background())
	candidates := make(chan *match, 100)
	g.Go(func() error {
		defer close(candidates)

		return filepath.Walk(cfg.Root(), func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}

			select {
			case candidates <- &match{info: info, path: path}:
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})

	})
	matches := make(chan *match, 100)

	for c := range candidates {
		candidate := c
		g.Go(func() error {
			var err error
			candidate.tags, err = tag.Read(candidate.path)
			if err != nil {
				log.Infof("Failed to read attributes from %s: %v.", candidate.Path, err)
				return nil
			}
			if !q.Match(candidate.tags) {
				return nil
			}
			candidate.stats = candidate.info.Sys().(*syscall.Stat_t)
			if candidate.stats == nil {
				log.Infof("Failed to read stats from %s.", candidate.Path)
				return nil
			}
			select {
			case matches <- candidate:
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})
	}

	go func() {
		g.Wait()
		close(matches)
	}()

	var result []Match
	for m := range matches {
		result = append(result, m)
	}
	return result, g.Wait()
}
