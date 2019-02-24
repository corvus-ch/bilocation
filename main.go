package main

import (
	"fmt"
	"github.com/corvus-ch/bilocation/cmd"
	stdLog "log"
	"os"

	"github.com/bketelsen/logr"
	"github.com/corvus-ch/logr/std"
	"github.com/corvus-ch/logr/writer_adapter"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func App(log logr.Logger) *kingpin.Application {
	w := writer_adapter.NewBufferedErrorWriter(log)
	app := kingpin.New("bilocation", "file management with tags")
	app.ErrorWriter(w)
	app.UsageWriter(w)
	app.Version(fmt.Sprintf("%v, commit %v, built at %v", version, commit, date))

	cfg := cmd.NewConfig(log)
	cmd.Tag(app, cfg)
	cmd.Untag(app, cfg)
	cmd.Search(app, cfg)
	cmd.Summary(app, cfg)

	return app
}

func main() {
	log := std.New(0, stdLog.New(os.Stderr, "", 0), stdLog.New(os.Stdout, "", 0))
	if _, err := App(log).Parse(os.Args[1:]); err != nil {
		log.Errorf("%s, try --help", err)
		os.Exit(1)
	}
}
