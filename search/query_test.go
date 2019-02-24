package search_test

import (
	"testing"

	"github.com/corvus-ch/bilocation/search"
	"github.com/corvus-ch/bilocation/tag"
	"github.com/stretchr/testify/assert"
)

func TestMatcher(t *testing.T) {
	tests := map[string][]bool{
		"":      {false, false, false, false, false},
		"*":     {true, true, true, true, true},
		"a":     {true, false, true, true, false},
		"b":     {false, true, true, false, false},
		"c":     {false, false, false, true, false},
		"d=*":   {false, false, false, true, true},
		"d=c":   {false, false, false, true, false},
		"a/b":   {false, false, true, false, false},
		"a,b":   {true, true, true, true, false},
		"a/b/c": {false, false, false, false, false},
		"a,b/c": {false, false, false, true, false},
	}

	data := []tag.Set{
		{"a": map[string]struct{}{}},
		{"b": map[string]struct{}{}},
		{"a": map[string]struct{}{}, "b": map[string]struct{}{}},
		{"a": map[string]struct{}{}, "c": map[string]struct{}{"d": {}}},
		{"d": map[string]struct{}{"d": {}}},
	}

	for q, test := range tests {
		t.Run(q, func(t *testing.T) {
			q := search.NewQuery(q)
			for i, set := range data {
				assert.Equal(t, test[i], q.Match(set), set)
			}
		})
	}
}
