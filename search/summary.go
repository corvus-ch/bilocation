package search

import (
	"github.com/corvus-ch/bilocation/tag"
)

// Summary walks the filesystem collecting tags of files matching a given query.
func Summary(cfg Config) (tag.Set, error) {
	matches, err := Search(cfg)
	if err != nil {
		return nil, err
	}

	set := tag.Set{}

	for _, m := range matches {
		for name, _ := range m.Tags() {
			set.Add(name, m.Tags().Get(name)...)
		}
	}

	return set, nil
}
