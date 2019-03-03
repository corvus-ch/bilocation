package matcher_test

import (
	"testing"

	"github.com/corvus-ch/bilocation/search/internal/matcher"
	"github.com/corvus-ch/bilocation/tag"
	"github.com/stretchr/testify/assert"
)

var sets = map[string]tag.Set{
	"empty":            {},
	"name":             {"test name": make(map[string]struct{})},
	"name_classifier":  {"test name": {"first classifier": struct{}{}}},
	"name_classifiers": {"test name": {"first classifier": struct{}{}, "second classifier": struct{}{}}},
}

func TestNil(t *testing.T) {
	for name, set := range sets {
		t.Run(name, func(t *testing.T) {
			assert.False(t, matcher.NewNil().Match(set))
		})
	}
}

func TestAny(t *testing.T) {
	for name, set := range sets {
		t.Run(name, func(t *testing.T) {
			assert.True(t, matcher.NewAny().Match(set))
		})
	}
}

func TestNot(t *testing.T) {
	for name, set := range sets {
		t.Run(name, func(t *testing.T) {
			assert.False(t, matcher.NewNot(matcher.NewAny()).Match(set))
			assert.True(t, matcher.NewNot(matcher.NewNil()).Match(set))
		})
	}
}

var noneTests = map[string]bool{
	"empty":            true,
	"name":             false,
	"name_classifier":  false,
	"name_classifiers": false,
}

func TestNone(t *testing.T) {
	for name, set := range sets {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, noneTests[name], matcher.NewNone().Match(set))
		})
	}
}

var nameTests = map[string]bool{
	"empty":            false,
	"name":             true,
	"name_classifier":  true,
	"name_classifiers": true,
}

func TestName(t *testing.T) {
	for name, set := range sets {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, nameTests[name], matcher.NewName("test name").Match(set))
		})
	}
}

var classifierTests = map[string]bool{
	"empty":            false,
	"name":             false,
	"name_classifier":  true,
	"name_classifiers": true,
}

func TestClassifier(t *testing.T) {
	for name, set := range sets {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, classifierTests[name], matcher.NewClassifier("first classifier").Match(set))
		})
	}
}

var nameAndClassifierTests = map[string]bool{
	"empty":            false,
	"name":             false,
	"name_classifier":  false,
	"name_classifiers": true,
}

func TestNameAndClassifier(t *testing.T) {
	for name, set := range sets {
		t.Run(name, func(t *testing.T) {
			m := matcher.NewNameAndClassifier("test name", "second classifier")
			assert.Equal(t, nameAndClassifierTests[name], m.Match(set))
		})
	}
}

var booleanTests = map[string][]matcher.Matcher{
	"none":   {matcher.NewNil(), matcher.NewNil()},
	"first":  {matcher.NewAny(), matcher.NewNil()},
	"second": {matcher.NewNil(), matcher.NewAny()},
	"both":   {matcher.NewAny(), matcher.NewAny()},
}

var andTest = map[string]bool{
	"none":   false,
	"first":  false,
	"second": false,
	"both":   true,
}

func TestAnd(t *testing.T) {
	for name, expect := range andTest {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, expect, matcher.NewAnd(booleanTests[name]).Match(tag.Set{}))
		})
	}
}

var orTest = map[string]bool{
	"none":   false,
	"first":  true,
	"second": true,
	"both":   true,
}

func TestOr(t *testing.T) {
	for name, expect := range orTest {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, expect, matcher.NewOr(booleanTests[name]).Match(tag.Set{}))
		})
	}
}
