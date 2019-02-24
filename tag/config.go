package tag

// Config provides the input for tag operations.
type Config interface {
	Tags() []string
	Path() string
}
