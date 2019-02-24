package tag_test

import (
	"fmt"
	"github.com/corvus-ch/bilocation/tag"
	"github.com/sebdah/goldie"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"sort"
	"testing"
)

func TestSet_Has(t *testing.T) {
	for _, name := range []string{"name", "name_classifier", "name_classifiers"} {
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(goldenFileName(name))
			assert.NoError(t, err)
			set, err := tag.NewTagSet(data)
			assert.NoError(t, err)
			assert.True(t, set.Has("test name"))
			assert.False(t, set.Has("another name"))
		})
	}
}

var hasClassifierTests = map[string][]bool{
	"empty":            {false, false},
	"name":             {false, false},
	"name_classifier":  {true, false},
	"name_classifiers": {true, true},
}

func TestSet_HasClassifier(t *testing.T) {
	for name, test := range hasClassifierTests {
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(goldenFileName(name))
			assert.NoError(t, err)
			set, err := tag.NewTagSet(data)
			assert.NoError(t, err)
			assert.Equal(t, test[0], set.HasClassifier("test name", "first classifier"))
			assert.Equal(t, test[1], set.HasClassifier("test name", "second classifier"))
			assert.False(t, set.HasClassifier("another name", "first classifier"))
			assert.False(t, set.HasClassifier("another name", "second classifier"))
		})
	}
}

var getTests = map[string][]string{
	"name":             {},
	"name_classifier":  {"first classifier"},
	"name_classifiers": {"first classifier", "second classifier"},
}

func TestSet_Get(t *testing.T) {
	for name, test := range getTests {
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(goldenFileName(name))
			assert.NoError(t, err)
			set, err := tag.NewTagSet(data)
			assert.NoError(t, err)
			classifiers := set.Get("test name")
			sort.Sort(sort.StringSlice(classifiers))
			assert.Equal(t, test, classifiers)
			assert.Empty(t, set.Get("another name"))
		})
	}
}

var addTests = map[string]func(*testing.T, tag.Set){
	"name": func(t *testing.T, set tag.Set) {
		set.Add("test name")
	},
	"name_classifier": func(t *testing.T, set tag.Set) {
		set.Add("test name", "first classifier")
	},
	"name_classifiers": func(t *testing.T, set tag.Set) {
		set.Add("test name", "first classifier", "second classifier")
	},
}

func TestSet_Add(t *testing.T) {
	for name, test := range addTests {
		t.Run(name, func(t *testing.T) {
			set := tag.Set{}
			test(t, set)
			bytes, err := set.Bytes()
			assert.NoError(t, err)
			goldie.Assert(t, name, bytes)
		})
	}
}

func TestSet_Del(t *testing.T) {
	for _, name := range []string{"name", "name_classifier", "name_classifiers"} {
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(goldenFileName(name))
			assert.NoError(t, err)
			set, err := tag.NewTagSet(data)
			assert.NoError(t, err)
			set.Del("test name")
			bytes, err := set.Bytes()
			assert.NoError(t, err)
			goldie.Assert(t, "empty", bytes)
		})
	}
}

var delClassifierTests = map[string]struct {
	classifier string
	fixture    string
}{
	"name_classifiers": {"second classifier", "name_classifier"},
	"name_classifier":  {"first classifier", "empty"},
}

func TestSet_DelClassifier(t *testing.T) {
	for name, test := range delClassifierTests {
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(goldenFileName(name))
			assert.NoError(t, err)
			set, err := tag.NewTagSet(data)
			assert.NoError(t, err)
			set.DelClassifier("test name", test.classifier)
			bytes, err := set.Bytes()
			assert.NoError(t, err)
			goldie.Assert(t, test.fixture, bytes)
		})
	}
}

func TestSet_Size(t *testing.T) {
	for _, size := range []int{1, 3, 5, 8, 13, 21} {
		t.Run(fmt.Sprint(size), func(t *testing.T) {
			set := tag.Set{}
			for i := 0; i < size; i++ {
				set.Add(fmt.Sprintf("%s-%d", t.Name(), i))
			}
			assert.Equal(t, size, set.Size())
		})
	}
}

func goldenFileName(name string) string {
	return filepath.Join(goldie.FixtureDir, fmt.Sprintf("%s%s", name, goldie.FileNameSuffix))
}
