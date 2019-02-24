package cmd

import (
	"github.com/corvus-ch/bilocation/tag"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Tag registers the tag sub command to the application.
func Tag(app *kingpin.Application, cfg *config) {
	c := app.Command("tag", "adds a tag to a file")
	c.Action(func(_ *kingpin.ParseContext) error {
		return tag.Tag(cfg)
	})
	c.Arg("file", "path to the file").
		Required().
		ExistingFileVar(&cfg.path)
	c.Arg("tag", "a list of tag names or `=` delimited key value pairs").
		Required().
		StringsVar(&cfg.tags)
}
