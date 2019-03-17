package cmd

import (
	"strings"

	"github.com/corvus-ch/bilocation/tag"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Inspect registers the inspect sub command to the application.
func Inspect(app *kingpin.Application, cfg *config) {
	c := app.Command("inspect", "show the tags of a file")
	c.Action(func(_ *kingpin.ParseContext) error {
		set, err := tag.Read(cfg.Path())
		if err != nil {
			return err
		}

		for name := range set {
			c := set.Get(name)
			if len(c) > 0 {
				cfg.Log().Infof("%s (%s)", name, strings.Join(c, ", "))
			} else {
				cfg.Log().Info(name)
			}
		}
		return nil
	})
	c.Arg("file", "path to the file").
		Required().
		ExistingFileVar(&cfg.path)
}
