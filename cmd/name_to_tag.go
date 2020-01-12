package cmd

import (
	"context"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/corvus-ch/bilocation/tag"
	"golang.org/x/sync/errgroup"
	"gopkg.in/alecthomas/kingpin.v2"
)

var test bool

// NameToTag registers the name to tag sub command to the application.
func NameToTag(app *kingpin.Application, cfg *config) {
	c := app.Command("ntt", "tags files based on their names")
	c.Action(func(_ *kingpin.ParseContext) error {
		log := cfg.Log()
		for path := range findFiles(cfg) {
			tags, err := tag.NewTagSet([]byte{})
			if err != nil {
				return err
			}
			for _, rule := range cfg.Tags() {
				parts := strings.SplitN(rule, "=", 2)
				re, err := regexp.Compile(parts[1])
				if err != nil {
					return err
				}
				matches := re.FindStringSubmatch(path)
				if len(matches) < 1 {
					continue
				}
				tags.Add(matches[1], parts[0])
			}
			if test {
				log.Infof("%s %v", path, tags)
			} else if tags.Size() > 0 {
				if err := tag.Write(path, tags); err != nil {
					return err
				}
			}
		}
		return nil
	})

	c.Flag("path", "path to the directory containing the files to be tagged").
		Short('C').
		Default(".").
		ExistingDirVar(&cfg.root)
	c.Flag("dry-run", "print the result and does not touch the files").
		Short('n').
		BoolVar(&test)
	c.Arg("rule", "a list of tag names and their regexp").
		Required().
		StringsVar(&cfg.tags)
}

func findFiles(cfg *config) chan string {
	g, ctx := errgroup.WithContext(context.Background())
	ch := make(chan string, 100)
	g.Go(func() error {
		defer close(ch)

		return filepath.Walk(cfg.Root(), func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}

			select {
			case ch <- path:
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})

	})

	return ch
}
