package main

import (
	"github.com/jinzhu/gorm"
	"github.com/luchoman08/ssllabs"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"sync"
	"./controllers"
	"./infraestructures"
	"./repositories"
	"./services"
	"./models"
	"./interfaces"
)

type IServiceContainer interface {
	Inject()
	GetModels() []interface{}
	GetDomainController() controllers.DomainController
	GetDatabaseService() interfaces.DatabaseService
	GetDomainRepository() interfaces.IDomainRepository
}

type kernel struct {
	DomainController controllers.DomainController
	DatabaseService interfaces.DatabaseService
	DomainRepository interfaces.IDomainRepository
}
func (k *kernel) GetDatabaseService() interfaces.DatabaseService {
	return k.DatabaseService
}
func (k *kernel) GetDomainController() controllers.DomainController {
	return k.DomainController
}
func (k *kernel) GetDomainRepository() interfaces.IDomainRepository{
	return k.DomainRepository
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
		&sslLabsHandler,
		&scraper,
		gH,
		serverService}
	domainService := &services.DomainService{domainRepository}
	domainController := controllers.DomainController{domainService}
	k.DomainRepository = domainRepository
	k.DomainController = domainController

}
func (k *kernel) injectDatabaseService(gH *infraestructures.GORMHandler, models []interface{}) {
	k.DatabaseService = &services.DatabaseService{gH, models}
}
func (k *kernel) Inject() {
	var client, _ = ssllabs.NewClient()
	db, err := gorm.Open("postgres", "postgresql://host_info:host_info@localhost:5432/host_info?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	gormHandler := infraestructures.GORMHandler{db}
	var whoIsHandler = infraestructures.WhoIsHandler{}
	serverRepo := repositories.ServerRepository{&gormHandler,&whoIsHandler }
	serverService := services.ServerService{&serverRepo}
	k.injectDomainController(&serverService, *client, db, &gormHandler)
	k.injectDatabaseService(&gormHandler, k.GetModels())
}
var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
			k.Inject()
		})
	}
	return k
}
