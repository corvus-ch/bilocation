package tag

import (
	"fmt"
	"strings"

	"github.com/corvus-ch/bilocation/tag/internal"
	"github.com/golang/protobuf/proto"
)

// Set implements a set data structure for file tags.
// Inspired by https://gist.github.com/bgadrian/cb8b9344d9c66571ef331a14eb7a2e80.
type Set map[string]map[string]struct{}

// NewTagSet returns a new tag set populated with the data read from the protocol buffer wire format data.
func NewTagSet(data []byte) (Set, error) {
	tags := &internal.Set{}
	err := proto.Unmarshal(data, tags)
	if err != nil {
		return nil, err
	}

	s := Set{}

	for _, tag := range tags.Tags {
		s.Add(tag.Name, tag.Classifiers...)
	}

	return s, nil
}

// Has checks if name is present in the set.
func (s Set) Has(name string) bool {
	_, ok := s[name]
	return ok
}

// HasClassifier checks if the name and classifier pair is present in the set.
func (s Set) HasClassifier(name, classifier string) bool {
	_, ok := s[name][classifier]
	return ok
}

// Read returns the list of classifiers for name.
func (s Set) Get(name string) []string {
	t, ok := s[name]
	if !ok {
		return make([]string, 0)
	}

	r := make([]string, 0, len(t))

	for c, _ := range t {
		r = append(r, c)
	}

	return r
}

// Add adds name with an optional list of classifiers to the set.
func (s Set) Add(name string, classifiers ...string) {
	if _, ok := s[name]; !ok {
		s[name] = make(map[string]struct{}, len(classifiers))
	}

	for _, classifier := range classifiers {
		if _, ok := s[name][classifier]; !ok {
			s[name][classifier] = struct{}{}
		}
	}
}

// Del deletes name with all its classifiers.
func (s Set) Del(name string) {
	delete(s, name)
}

// DelClassifier removes the classifier from name.
func (s Set) DelClassifier(name, classifier string) {
	if data, ok := s[name]; ok {
		delete(data, classifier)
		if len(data) == 0 {
			delete(s, name)
		}
	}
}

// Size returns the number of tags in the set.
func (s Set) Size() int {
	return len(s)
}

// Bytes returns the set in the protocol buffer wire format.
func (s Set) Bytes() ([]byte, error) {
	set := &internal.Set{}
	set.Tags = make([]*internal.Tag, 0, s.Size())

	for name, classifiers := range s {
		tag := &internal.Tag{Name: name}
		tag.Classifiers = make([]string, 0, len(classifiers))
		for classifier, _ := range classifiers {
			if len(classifier) > 0 {
				tag.Classifiers = append(tag.Classifiers, classifier)
			}
		}
		set.Tags = append(set.Tags, tag)
	}
	return proto.Marshal(set)
}

func (s Set) String() string {
	var pairs []string
	for value, keys := range s {
		for key := range keys {
			pairs = append(pairs, fmt.Sprintf("%s='%s'", key, value))
		}
	}

	return strings.Join(pairs, " ")
}
