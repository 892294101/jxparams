package jxparams

type Config struct {
	must         bool
	defaultValue string
	prefix       bool
	suffix       bool
}

func NewConfig() *Config {
	return new(Config)
}

func (c *Config) getMust() bool { return c.must }

func (c *Config) getDefault() string { return c.defaultValue }

func (c *Config) SetMust() *Config {
	c.must = true
	return c
}

func (c *Config) SetDefault(s string) *Config {
	c.defaultValue = s
	return c
}

func (c *Config) SetPrefix() *Config {
	c.prefix = true
	c.suffix = false
	//c.must = false
	return c
}

func (c *Config) SetSuffix() *Config {
	c.suffix = true
	c.prefix = false
	//c.must = false
	return c
}
