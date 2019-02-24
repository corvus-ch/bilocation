package cmd

import (
	"github.com/corvus-ch/bilocation/search"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Search registers the search sub command to the application.
func Search(app *kingpin.Application, cfg *config) {
	c := app.Command("search", "search for files wit a given tag")
	c.Action(func(_ *kingpin.ParseContext) error {
		matches, err := search.Search(cfg)
		if err != nil {
			return err
		}

		for _, match := range matches {
			cfg.Log().Info(match.Path())
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
