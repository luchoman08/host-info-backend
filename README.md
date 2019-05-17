# Host info (backend)
Web service que se encarga de ofrecer funcionalidades para la visualización de información sobre hosts web

Este repositorio fue altamente influenciado por el repositorio 
[service-pattern-go](https://github.com/irahardianto/service-pattern-go), el cual
muestra una implementación usando inversión de control con injección de dependencias
en busca de seguir los principios SOLID.

Como ejecutar: `go run main.go router.go servicecontainer.go`

Como iniciar las tablas de la base de datos:  `go run cli.go servicecontainer.go migrate`

Como borrar las tablas de la base de datos: `go run cli.go servicecontainer.go drop-tables`
