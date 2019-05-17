package models


type Config struct {
	DbDialect  string `yaml:"db_dialect"`
	DbPath  string `yaml:"db_path"`
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}