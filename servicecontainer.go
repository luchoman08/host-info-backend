package main

import (
	"./controllers"
	"./infraestructures"
	"./interfaces"
	"./models"
	"./repositories"
	"./services"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/luchoman08/ssllabs"
	"log"
	"sync"
)

// IServiceContainer provide all methods than are needed to be public
// across the app
type IServiceContainer interface {
	Inject()
	GetModels() []interface{}
	GetDomainController() controllers.DomainController
	GetServerService() interfaces.ServerService
	GetDomainService() interfaces.DomainService
	GetDatabaseService() interfaces.DatabaseService
	GetDomainRepository() interfaces.DomainRepository
	GetConfigService() interfaces.ConfigService
}

type kernel struct {
	DomainController controllers.DomainController
	DomainService    interfaces.DomainService
	ServerService    interfaces.ServerService
	DatabaseService  interfaces.DatabaseService
	ConfigService    interfaces.ConfigService
	DomainRepository interfaces.DomainRepository
}

func (k *kernel) GetConfigService() interfaces.ConfigService {
	return k.ConfigService
}

func (k *kernel) GetDatabaseService() interfaces.DatabaseService {
	return k.DatabaseService
}
func (k *kernel) GetDomainController() controllers.DomainController {
	return k.DomainController
}
func (k *kernel) GetDomainRepository() interfaces.DomainRepository {
	return k.DomainRepository
}
func (k *kernel) GetDomainService() interfaces.DomainService {
	return k.DomainService
}
func (k *kernel) GetServerService() interfaces.ServerService {
	return k.ServerService
}

func (k *kernel) GetModels() []interface{} {
	var mdls []interface{}
	mdls = append(mdls, &models.DomainModel{})
	mdls = append(mdls, &models.ServerModel{})
	return mdls
}
func (k *kernel) injectDomainController(serverService interfaces.ServerService, client ssllabs.Client, db *gorm.DB, gH *infraestructures.GORMHandler) {

	sslLabsHandler := infraestructures.SSLLabsHandler{}
	sslLabsHandler.Client = client
	scraper := infraestructures.GoScraperHandler{}
	domainRepository := &repositories.DomainRepository{
		SSLabsHandler:    &sslLabsHandler,
		GoScraperHandler: &scraper,
		GORMHandler:      gH,
		ServerService:    serverService}
	domainService := &services.DomainService{DomainRepository: domainRepository, ServerService: serverService}
	k.DomainService = domainService
	domainController := controllers.DomainController{DomainService: domainService}
	k.DomainRepository = domainRepository
	k.DomainController = domainController

}
func (k *kernel) injectDatabaseService(gH *infraestructures.GORMHandler, interfaces []interface{}) {
	k.DatabaseService = &services.DatabaseService{GORMHandler: gH, Interfaces: interfaces}
}
func (k *kernel) Inject() {
	configService := &services.ConfigService{}
	err := configService.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	k.ConfigService = configService
	config := k.ConfigService.GetConfig()
	var client, _ = ssllabs.NewClient(ssllabs.Config{Retries: config.SsllRetries, Timeout: config.SsllTimeout})
	db, err := gorm.Open(config.DbDialect, config.DbPath)
	if err != nil {
		log.Fatal(err)
	}
	gormHandler := infraestructures.GORMHandler{Db: db}
	var whoIsHandler = infraestructures.WhoIsHandler{}
	serverRepo := repositories.ServerRepository{GORMHandler: &gormHandler, WhoIsHandler: &whoIsHandler}
	serverService := services.ServerService{ServerRepository: &serverRepo}
	k.ServerService = &serverService
	k.injectDomainController(&serverService, *client, db, &gormHandler)
	k.injectDatabaseService(&gormHandler, k.GetModels())
}

var (
	k             *kernel
	containerOnce sync.Once
)

// ServiceContainer provide all app kernel functions and init only one time
func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
			k.Inject()
		})
	}
	return k
}
