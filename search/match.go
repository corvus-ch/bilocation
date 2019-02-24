package search

import (
	"os"
	"syscall"

	"github.com/corvus-ch/bilocation/tag"
)

// Match represents a file matching a query during a search.
type Match interface {
	Info() os.FileInfo
	Path() string
	Tags() tag.Set
	Stats() *syscall.Stat_t
}

type match struct {
	info  os.FileInfo
	path  string
	tags  tag.Set
	stats *syscall.Stat_t
}

func (m *match) Info() os.FileInfo {
	return m.info
}

func (m *match) Path() string {
	return m.path
}

func (m *match) Tags() tag.Set {
	return m.tags
}

func (m *match) Stats() *syscall.Stat_t {
	return m.stats
}
