package models

type Config struct {
	DbDialect   string `yaml:"db_dialect"`
	DbPath      string `yaml:"db_path"`
	Port        string `yaml:"port"`
	Host        string `yaml:"host"`
	SsllRetries int    `yaml:"ssll_retries"`
	SsllTimeout int    `yaml:"ssll_timeout"`
}
