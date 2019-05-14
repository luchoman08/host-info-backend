package main

import (
	"fmt"
	"github.com/luchoman08/ssllabs"
	"./app"
)

func main() {
	c, _ := ssllabs.NewClient()
	var domainInfo, _ = app.ExtractDomainInfo("https://google.com/", c)
	fmt.Println(domainInfo)
}
