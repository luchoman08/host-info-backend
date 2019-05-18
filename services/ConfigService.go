package services

import (
	"../models"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const configFileName = "config.yml";
// ConfigService implements the methods for access and load the
// app config values
type ConfigService struct {
	config models.Config
}
// ReadConfig load the initial app config than is saved in a
// yml file
func (service *ConfigService) ReadConfig() (err error) {
	yamlFile, err := ioutil.ReadFile(configFileName)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &service.config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return
}
// GetConfig returns the app config than exists at the moment
func (service *ConfigService) GetConfig() *models.Config {
	return &service.config
}
