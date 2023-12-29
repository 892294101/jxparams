package jxparams

type Config struct {
	must         bool
	defaultValue string
}

func NewConfig() *Config {
	return new(Config)
}

func (c *Config) getMust() bool { return c.must }

func (c *Config) getDefault() string { return c.defaultValue }

func (c *Config) SetMust(m bool) *Config {
	c.must = m
	return c
}

func (c *Config) SetDefault(s string) *Config {
	c.defaultValue = s
	return c
}
