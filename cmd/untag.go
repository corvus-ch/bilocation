package cmd

import (
	"github.com/corvus-ch/bilocation/tag"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Untag registers the untag sub command to the application.
func Untag(app *kingpin.Application, cfg *config) {
	c := app.Command("untag", "removes a tag from a file")
	c.Action(func(_ *kingpin.ParseContext) error {
		return tag.Untag(cfg)
	})
	c.Arg("file", "path to the file").
		Required().
		ExistingFileVar(&cfg.path)
	c.Arg("tag", "a list of tag names or `=` delimited key value pairs").
		Required().
		StringsVar(&cfg.tags)
}
