package cmd

import (
	"github.com/corvus-ch/bilocation/search"
	"gopkg.in/alecthomas/kingpin.v2"
	"strings"
)

// Summary registers the summary sub command to the application.
func Summary(app *kingpin.Application, cfg *config) {
	c := app.Command("summary", "prints a summary of existing tags for a given search query")
	c.Action(func(_ *kingpin.ParseContext) error {
		set, err := search.Summary(cfg)
		if err != nil {
			return err
		}

		for name, _ := range set {
			c := set.Get(name)
			if len(c) > 0 {
				cfg.Log().Infof("%s (%s)", name, strings.Join(c, ", "))
			} else {
				cfg.Log().Info(name)
			}
		}

		return nil
	})
	c.Arg("query", "the search query").
		Required().
		StringVar(&cfg.query)
	c.Flag("path", "the root dir to start the search with").
		Short('C').
		Default(".").
		ExistingDirVar(&cfg.root)
}
