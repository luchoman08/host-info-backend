package app

import (
	"net/url"
	"strings"
)

// NormalizePageIcoURL process a page icon URL using the main url of the page,
// for example, `/favicon.ico`  this is converted in a absolute url. If the route given is
// absolute the same url is returned
func NormalizePageIcoURL(route string, mainURL url.URL) string {
	var iconURL, _ = url.Parse(route)
	if iconURL.Scheme == "" {
		iconURL.Scheme = mainURL.Scheme
	}
	if iconURL.Host == "" {
		iconURL.Host = mainURL.Host
	}
	return iconURL.String()
}

// NormalizeURL when a string url is parsed without scheme (protocol), the parsed Host route is empty
// and the url Path is equal to the input string, but this is wrong, if this is
// the case, the Host is corrected with this method
// *Example*
// When a string 'google.com' is parsed to url, url.Host is emtpy and url.Path is equal
// to 'google.com', after `NormalizeURL`, url.Host is equal to google.com and
// ur.Path is empty
func NormalizeURL(u *url.URL) {
	if u.Host == "" {
		u.Host = strings.Trim(u.Path, "/")
		u.Path = ""
	}
}

// NormalizeURLWithScheme When a string url is parsed without scheme (protocol), the parsed Host route is empty
// and the url Path is equal to the input string, but this is wrong, if this is
// the case, the Path, Host and Scheme are corrected with this method
func NormalizeURLWithScheme(url *url.URL, scheme string) {
	if url.Scheme == "" {
		url.Scheme = scheme
	}
	NormalizeURL(url)
}
