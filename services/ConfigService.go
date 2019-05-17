package services

import (
	"../models"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)
const configFileName = "config.yml";
type ConfigService struct {
	config models.Config
}

func (service *ConfigService ) ReadConfig() (err error) {
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
func (service *ConfigService) GetConfig() models.Config {
	return service.config
}
