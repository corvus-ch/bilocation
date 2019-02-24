package main

import (
	"testing"

	"github.com/corvus-ch/logr/buffered"
	"github.com/sebdah/goldie"
	"github.com/stretchr/testify/assert"
)

var tests = map[string][]string{
	"help":    {"help"},
	"tag":     {"help", "tag"},
	"untag":   {"help", "untag"},
	"search":  {"help", "search"},
	"summary": {"help", "summary"},
}

func TestApp(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			log := buffered.New(0)
			a := App(log)
			a.Terminate(func(i int) {
				assert.Equal(t, 0, i)
			})
			_, err := a.Parse(test)
			assert.NoError(t, err)
			goldie.Assert(t, name, log.Buf().Bytes())
		})
	}
}
