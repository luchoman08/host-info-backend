package app

import (
	"fmt"
	"github.com/likexian/whois-go"
	"github.com/likexian/whois-parser-go"
	"github.com/luchoman08/goscraper"
	"github.com/luchoman08/ssllabs"
	"net/url"
	"strings"
)

type whoIsOutput struct {
	Country string
	OrgName string
}
type WebPageInfo struct {
	Title   string
	IconUrl string
}

func parseWhoIsText(text string) map[string]string {
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
func getWhoIsOutput(ipAddress string) whoIsOutput {
	var whoIsOut = whoIsOutput{}
	var text, _ = whois.Whois(ipAddress)
	var whoIsKeyValue = parseWhoIsText(text)
	whoIsOut.Country = whoIsKeyValue["Country"]
	whoIsOut.OrgName = whoIsKeyValue["OrgName"]
	return whoIsOut
}

func extractServerInfo(endpoint ssllabs.Endpoint) ServerInfo {
	var serverInfo = ServerInfo{}
	var whoIs = getWhoIsOutput(endpoint.IPAddress)
	serverInfo.SslGrade = endpoint.Grade
	serverInfo.Address = endpoint.ServerName
	serverInfo.Country = whoIs.Country
	serverInfo.Owner = whoIs.OrgName
	return serverInfo
}

func normalizePageIcoUrl(route string, mainUrl url.URL) string {
	var iconUrl, _ = url.Parse(route)
	if iconUrl.Scheme == "" {
		iconUrl.Scheme = mainUrl.Scheme
	}
	if iconUrl.Host == "" {
		iconUrl.Host = mainUrl.Host
	}
	fmt.Println(iconUrl.Scheme)
	return iconUrl.String()
}

func extractWebPageInfo(url url.URL) (WebPageInfo, error) {
	var webPageInfo = WebPageInfo{}
	s, err := goscraper.Scrape(url.String(), 5)

	if err != nil {
		fmt.Println(err)
		return webPageInfo, err
	}
	webPageInfo.IconUrl = normalizePageIcoUrl(s.Preview.Icon, url)
	webPageInfo.Title = s.Preview.Title
	return webPageInfo, err
}

func normalizeUrl(url *url.URL) {
	if url.Host == "" {
		url.Host = strings.Trim(url.Path, "/")
		url.Path = ""
	}
}

// When a string url is parsed without scheme (protocol), the parsed Host route is empty
// and the url Path is equal to the input string, but this is wrong, if this is
// the case, the Path, Host and Scheme are corrected with this method
func normalizeUrlWithScheme(url *url.URL, scheme string) {
	if url.Scheme == "" {
		url.Scheme = scheme
	}
	normalizeUrl(url)

}
func ExtractDomainInfo(route string, c *ssllabs.Client) (domainInfo DomainInfo, err error) {
	domainInfo = DomainInfo{}
	var url, urlError = url.Parse(route)

	normalizeUrl(url)
	if urlError != nil {
		fmt.Println("Url format error")
		err = ErrUrlMalformed
		return
	}

	hostInfo, ssll_err := c.GetDetailedReport(url.Host)
	if ssll_err != nil {
		err = ssll_err
		return
	}
	normalizeUrlWithScheme(url, hostInfo.Protocol)
	var info, webError = extractWebPageInfo(*url)
	if webError == nil {
		domainInfo.Title = info.Title
		domainInfo.Logo = info.IconUrl
	}
	if hostInfo.Endpoints != nil {
		domainInfo.SslGrade = hostInfo.Endpoints[0].Grade
	}
	domainInfo.IsDown = hostInfo.Status != "READY"
	var servers []ServerInfo
	for i := 0; i < len(hostInfo.Endpoints); i++ {
		servers = append(servers, extractServerInfo(hostInfo.Endpoints[i]))
	}
	domainInfo.Servers = servers
	return
}
