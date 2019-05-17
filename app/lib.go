package app

import (
	"fmt"
	"net/url"
	"strings"
)

func NormalizePageIcoUrl(route string, mainUrl url.URL) string {
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
// When a string url is parsed without scheme (protocol), the parsed Host route is empty
// and the url Path is equal to the input string, but this is wrong, if this is
// the case, the Host is corrected with this method
// *Example*
// When a string 'google.com' is parsed to url, url.Host is emtpy and url.Path is equal
// to 'google.com', after `NormalizeUrl`, url.Host is equal to google.com and
// ur.Path is empty
func NormalizeUrl(u *url.URL) {
	if u.Host == "" {
		u.Host = strings.Trim(u.Path, "/")
		u.Path = ""
	}
}

// When a string url is parsed without scheme (protocol), the parsed Host route is empty
// and the url Path is equal to the input string, but this is wrong, if this is
// the case, the Path, Host and Scheme are corrected with this method
func NormalizeUrlWithScheme(url *url.URL, scheme string) {
	if url.Scheme == "" {
		url.Scheme = scheme
	}
	NormalizeUrl(url)
}