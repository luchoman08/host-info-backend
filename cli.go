package main

import (
	"./interfaces"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
	"os"
)

var app = cli.NewApp()

func commands(
	dbService interfaces.DatabaseService,
	domainService interfaces.DomainService,
	serverService interfaces.ServerService) {
	app.Commands = []cli.Command{
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "Migrate database",
			Action: func(c *cli.Context) {
				fmt.Println("Iniciando migraciones")
				dbService.Migrate()
				fmt.Println("Migraciones terminadas")
			},
		},
		{
			Name:    "drop-tables",
			Aliases: []string{"dt"},
			Usage:   "Drop tables",
			Action: func(c *cli.Context) {
				fmt.Println("Iniciando borrado de tablas")
				dbService.DropTables()
				fmt.Println("Borrado terminado")
			},
		},
		{
			Name:    "get-local-server",
			Aliases: []string{"g-l-s"},
			Usage:   "Get local server by ip",
			Action: func(c *cli.Context) {
				fmt.Println(serverService.GetServerByIP(c.Args()[0]))
			},
		},
		{
			Name:    "look-ip",
			Aliases: []string{"l-ip"},
			Usage:   "Look up for IP given a domain name",
			Action: func(c *cli.Context) {
				fmt.Println(net.LookupIP((c.Args()[0])))
			},
		},
		{
			Name:    "get-local-domain",
			Aliases: []string{"g-l-d"},
			Usage:   "Get local domain by host name",
			Action: func(c *cli.Context) {
				fmt.Println(domainService.GetDomain(c.Args()[0]))
			},
		},
	}
}

func info() {
	app.Name = "Host Info Cli"
	app.Usage = "Run all cli methods for host info cli app"
	app.Author = "Luis Gerardo Manrique Cardona <luis.manrique@correounivalle.edu.co"
	app.Version = "1.0.0"
}
func main() {
	info()
	commands(
		ServiceContainer().GetDatabaseService(),
		ServiceContainer().GetDomainService(),
		ServiceContainer().GetServerService())
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
