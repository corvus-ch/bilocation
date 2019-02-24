package tag

import (
	"github.com/pkg/xattr"
	"strings"
)

// XattrName holds the extended attribute name the tags are stored with.
const XattrName = "bilocation.corvus-ch.name/v1/tags"

// Config provides the input for tag operations.
type Config interface {
	Tags() []string
	Path() string
}

// Tag adds a tag to a file.
func Tag(cfg Config) error {
	s, err := Read(cfg.Path())
	if err != nil {
		return err
	}
	for _, arg := range cfg.Tags() {
		key, value, _ := nameAndClassifier(arg)
		s.Add(key, value)
	}

	return Write(cfg.Path(), s)
}

// Untag removes a tag from a file.
func Untag(cfg Config) error {
	s, err := Read(cfg.Path())
	if err != nil {
		return nil
	}
	for _, arg := range cfg.Tags() {
		name, classifier, hasClassifier := nameAndClassifier(arg)
		if hasClassifier {
			s.DelClassifier(name, classifier)
		} else {
			s.Del(name)
		}
	}

	return Write(cfg.Path(), s)
}

// Read reads the a tag set from the files extended attributes.
func Read(file string) (Set, error) {
	data, err := xattr.Get(file, XattrName)
	if err != nil && !isAttributeNotFoundError(err) {
		return nil, err
	}

	return NewTagSet(data)
}

// Read writes the tag set to the files extended attributes.
func Write(file string, set Set) error {
	if set.Size() < 1 {
		err := xattr.Remove(file, XattrName)
		if err != nil && isAttributeNotFoundError(err) {
			return nil
		}

		return err
	}

	data, err := set.Bytes()
	if err != nil {
		return err
	}

	return xattr.Set(file, XattrName, data)
}

func isAttributeNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "attribute not found")
}

func nameAndClassifier(s string) (string, string, bool) {
	key := s
	value := ""
	hasValue := strings.Contains(s, "=")
	if hasValue {
		parts := strings.SplitN(s, "=", 2)
		key = parts[1]
		value = parts[0]
	}

	return key, value, hasValue
}
