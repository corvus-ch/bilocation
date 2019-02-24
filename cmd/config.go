package cmd

import "github.com/bketelsen/logr"

type config struct {
	root  string
	path  string
	query string
	tags  []string
	log   logr.Logger
}

func NewConfig(log logr.Logger) *config {
	return &config{log: log}
}

func (c *config) Query() string {
	return c.query
}

func (c config) Root() string {
	return c.root
}

func (c config) Path() string {
	return c.path
}

func (c config) Tags() []string {
	return c.tags
}

func (c config) Log() logr.Logger {
	return c.log
}
