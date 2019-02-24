package search

// Config provides the input to the Search operations.
type Config interface {
	Query() string
	Root() string
}
