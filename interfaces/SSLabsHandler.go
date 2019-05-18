package interfaces

import "github.com/luchoman08/ssllabs"

type SSLabsHandler interface {
	GetDetailedReport(route string) (ssllabs.Host, error)

	ReadyState() string
}
