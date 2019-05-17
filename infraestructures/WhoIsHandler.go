package infraestructures

import (
	"github.com/likexian/whois-go"
	"github.com/likexian/whois-parser-go"
	"strings"
)

type WhoIsHandler struct {
}

func (handler *WhoIsHandler) GetWhoIsRaw(ipAddress string) (text string, err error) {
	text, err = whois.Whois(ipAddress)
	return
}

func (handler *WhoIsHandler) ParseWhoIsText(text string) map[string]string {
	whoisText := strings.Replace(text, "\r", "", -1)
	whoisText = whoisparser.TextReplacer.ReplaceAllString(whoisText, "\n$1: $2")
	var keyValue = make(map[string]string)
	whoisLines := strings.Split(text, "\n")
	for i := 0; i < len(whoisLines); i++ {
		if strings.Contains(whoisLines[i], ":") {
			var split = strings.Split(whoisLines[i], ":")
			keyValue[split[0]] = strings.Trim(split[1], " ")
		}
	}
	return keyValue
}
func (handler *WhoIsHandler) GetWhoIsParsed(ipAddress string) (out map[string]string, err error) {
	out = make(map[string]string)
	text, err := handler.GetWhoIsRaw(ipAddress)
	if err != nil {
		return out, err
	}
	out = handler.ParseWhoIsText(text)
	return out, err
}
