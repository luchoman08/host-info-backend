package models

// Config represents the struct of all app configuration values
type Config struct {
	DbDialect        string `yaml:"db_dialect"`
	DbPath           string `yaml:"db_path"`
	Port             string `yaml:"port"`
	Host             string `yaml:"host"`
	DefaultPageLimit int    `yaml:"default_page_limit"`
	SsllRetries      int    `yaml:"ssll_retries"`
	SsllTimeout      int    `yaml:"ssll_timeout"`
}
