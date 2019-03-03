package matcher

import "github.com/corvus-ch/bilocation/tag"

// Matcher defines the interface for matching a tag set against a set of criteria.
type Matcher interface {
	// Match checks if the tag set matches the criteria.
	Match(tag.Set) bool
}

type any struct{}

// NewAny create a matcher which always matches.
func NewAny() Matcher {
	return &any{}
}

func (m *any) Match(_ tag.Set) bool {
	return true
}

type and struct {
	matchers []Matcher
}

// NewAnd returns a matcher which requires all of the provided matchers to match.
func NewAnd(m []Matcher) Matcher {
	return &and{matchers: m}
}

func (m *and) Match(set tag.Set) bool {
	r := len(m.matchers) > 0
	for _, matcher := range m.matchers {
		r = r && matcher.Match(set)
	}

	return r
}

type or struct {
	matchers []Matcher
}

// NewOr returns a matcher which requires any of the provided matchers to match.
func NewOr(m []Matcher) Matcher {
	return &or{matchers: m}
}

func (m *or) Match(set tag.Set) bool {
	r := false
	for _, matcher := range m.matchers {
		r = r || matcher.Match(set)
	}

	return r
}

type tagName struct {
	name string
}

// NewName creates a matcher which requires the tag to have the given name.
func NewName(name string) Matcher {
	return &tagName{name: name}
}

func (m *tagName) Match(set tag.Set) bool {
	return set.Has(m.name)
}

type tagClassifier struct {
	classifier string
}

// NewClassifier creates a matcher which requires the tag to have the given classifier.
func NewClassifier(classifier string) Matcher {
	return &tagClassifier{classifier: classifier}
}
func (m *tagClassifier) Match(set tag.Set) bool {
	for name, _ := range set {
		if set.HasClassifier(name, m.classifier) {
			return true
		}
	}

	return false
}

type nameAndClassifier struct {
	name       string
	classifier string
}

// NewNameAndClassifier creates a matcher which requires the tag to have the given name and classifier pair.
func NewNameAndClassifier(name, classifier string) Matcher {
	return &nameAndClassifier{name: name, classifier: classifier}
}

func (m *nameAndClassifier) Match(set tag.Set) bool {
	return set.HasClassifier(m.name, m.classifier)
}

type none struct{}

// NewNone creates a matcher which matches empty tag sets.
func NewNone() Matcher {
	return &none{}
}

func (m *none) Match(set tag.Set) bool {
	return len(set) == 0
}

type not struct {
	matcher Matcher
}

// NewNot creates a matcher which never matches.
func NewNot(m Matcher) Matcher {
	return &not{matcher: m}
}

func (m *not) Match(set tag.Set) bool {
	return !m.matcher.Match(set)
}

type null struct{}

// NewNil creates a matcher which never matches.
func NewNil() Matcher {
	return &null{}
}

func (m *null) Match(set tag.Set) bool {
	return false
}
