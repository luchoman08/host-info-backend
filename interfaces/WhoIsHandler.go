package interfaces
// WhoIsHandler provide the methods for access the who is external info
type WhoIsHandler interface {
	ParseWhoIsText(text string) map[string]string
	GetWhoIsRaw(ipAddress string) (text string, err error)
	GetWhoIsParsed(ipAddress string) (out map[string]string, err error)
}
