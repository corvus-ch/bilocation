package tag_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/corvus-ch/bilocation/tag"
	"github.com/pkg/xattr"
	"github.com/sebdah/goldie"
	"github.com/stretchr/testify/assert"
)

var tagTests = map[string][]string{
	"name":             {"test name"},
	"name_classifier":  {"first classifier=test name"},
	"name_classifiers": {"first classifier=test name", "second classifier=test name"},
}

func TestTag(t *testing.T) {
	for name, tags := range tagTests {
		t.Run(name, func(t *testing.T) {
			f, err := ioutil.TempFile("", "")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(f.Name())

			err = tag.Tag(&config{path: f.Name(), tags: tags})
			assert.NoError(t, err)

			data, err := xattr.Get(f.Name(), tag.XattrName)
			assert.NoError(t, err)
			goldie.Assert(t, name, data)
		})
	}
}

var untagTests = map[string]struct {
	golden string
	tags   []string
}{
	"name": {"empty", []string{"test name"}},
	"name_classifier": {"empty", []string{"first classifier=test name"}},
	"name_classifiers": {"name_classifier", []string{"second classifier=test name"}},
}

func TestUntag(t *testing.T) {
	for name, test := range untagTests {
		t.Run(name, func(t *testing.T) {
			f, err := ioutil.TempFile("", "")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(f.Name())

			data, err := ioutil.ReadFile(goldenFileName(name))
			assert.NoError(t, err)
			xattr.Set(f.Name(), tag.XattrName, data)

			err = tag.Untag(&config{path: f.Name(), tags: test.tags})
			assert.NoError(t, err)

			data, _ = xattr.Get(f.Name(), tag.XattrName)
			goldie.Assert(t, test.golden, data)
		})
	}
}

type config struct {
	path string
	tags []string
}

func (c *config) Path() string {
	return c.path
}

func (c *config) Tags() []string {
	return c.tags
}
