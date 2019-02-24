package cmd

import (
	"github.com/corvus-ch/bilocation/search"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Summary registers the summary sub command to the application.
func Summary(app *kingpin.Application, cfg *config) {
	c := app.Command("summary", "prints a summary of existing tags for a given search query")
	c.Action(func(_ *kingpin.ParseContext) error {
		return search.Summary(cfg)
	})
	c.Arg("query", "the search query").
		Required().
		StringVar(&cfg.query)
	c.Flag("path", "the root dir to start the search with").
		Short('C').
		Default(".").
		ExistingDirVar(&cfg.root)
}
