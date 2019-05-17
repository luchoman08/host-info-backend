package infraestructures

import (
	"github.com/luchoman08/ssllabs"
)

type SSLLabsHandler struct {
	Client ssllabs.Client
}

func ( handler *SSLLabsHandler) GetDetailedReport(route string) (ssllabs.Host, error) {
	return handler.Client.GetDetailedReport(route)
}

func (handler *SSLLabsHandler) ReadyState() string {
	return "READY"
}