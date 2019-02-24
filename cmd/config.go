package cmd

type config struct {
	root  string
	path  string
	query string
	tags  []string
}

func NewConfig() *config {
	return &config{}
}

func (c *config) Query() string {
	return c.query
}

func (c config) Root() string {
	return c.root
}

func (c config) Path() string {
	return c.path
}

func (c config) Tags() []string {
	return c.tags
}
