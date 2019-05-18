package interfaces

import "github.com/luchoman08/ssllabs"

// SSLabsHandler provide the methods for access the ssl labs API
type SSLabsHandler interface {
	GetDetailedReport(route string) (ssllabs.Host, error)

	ReadyState() string
}
