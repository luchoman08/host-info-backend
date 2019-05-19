package infraestructures

import (
	"github.com/likexian/whois-go"
	"github.com/likexian/whois-parser-go"
	"strings"
)

// WhoIsHandler implements all required methods for access the
// who is info
type WhoIsHandler struct {
}

// GetWhoIsRaw returns the info from who is as simple string, exactly than who is
// server was returned the info
func (handler *WhoIsHandler) GetWhoIsRaw(ipAddress string) (text string, err error) {
	text, err = whois.Whois(ipAddress)
	return
}

// ParseWhoIsText return a key value struct with all the who is info retrieved
// extracted form a raw who is response
func (handler *WhoIsHandler) ParseWhoIsText(text string) map[string]string {
	whoisText := strings.Replace(text, "\r", "", -1)
	whoisText = whoisparser.TextReplacer.ReplaceAllString(whoisText, "\n$1: $2")
	var keyValue = make(map[string]string)
	whoisLines := strings.Split(text, "\n")
	for i := 0; i < len(whoisLines); i++ {
		if strings.Contains(whoisLines[i], ":") {
			var split = strings.Split(whoisLines[i], ":")
			keyValue[strings.ToLower(split[0])] = strings.Trim(split[1], " ")
		}
	}
	return keyValue
}

// GetWhoIsParsed query a who is server and return its response parsed in a key value struct
func (handler *WhoIsHandler) GetWhoIsParsed(ipAddress string) (out map[string]string, err error) {
	out = make(map[string]string)
	text, err := handler.GetWhoIsRaw(ipAddress)
	if err != nil {
		return out, err
	}
	out = handler.ParseWhoIsText(text)
	return out, err
}
