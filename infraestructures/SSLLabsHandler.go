package infraestructures

import (
	"github.com/luchoman08/ssllabs"
)

// SSLLabsHandler implements all methods for access the ssllabs API
type SSLLabsHandler struct {
	Client ssllabs.Client
}

// GetDetailedReport returns the info consulted from ssllabs API if this info exists,
// otherwise a error is returned
func (handler *SSLLabsHandler) GetDetailedReport(route string) (ssllabs.Host, error) {
	return handler.Client.GetDetailedReport(route)
}

// ReadyState returns the value that is understood as the server info is available.
func (handler *SSLLabsHandler) ReadyState() string {
	return "READY"
}
