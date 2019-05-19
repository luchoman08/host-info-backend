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
// GetMinorSSLGradeFromList returns the minor ssl grade from list of grades
// if empty list is given, empty string is returned
func GetMinorSSLGradeFromList(grades []string) (minorGrade string) {
	minorGrade = ""
	for _, grade := range grades {
		minorGrade = GetMinorSSLGrade(minorGrade, grade)
	}
	return
}

// GetMinorSSLGrade compares two SSL grades and return the major grade,
// for example, if g1 = A and g2 is B, return B, also works with a grade
// modifiers like + or -, if gi = A+ and g2 = A- ,  A+ is returned
func GetMinorSSLGrade(g1 string, g2 string) (string) {
	if g1 == "" && g2 != "" {
		return g2
	}
	if g2 == "" && g1 != "" {
		return g1
	}
	if g1 == g2 {
		return g2
	}
	g1Runes := []rune(g1)
	g2Runes := []rune(g2)
	g1Grade := g1Runes[0]
	g2Grade := g2Runes[0]
	var g1Modifier, g2Modifier rune
	if len(g1Runes) > 1 {
		g1Modifier = g1Runes[0]
	} else {
		g1Modifier = 0
	}
	if len(g2Runes) > 1 {
		g2Modifier = g2Runes[0]
	} else {
		g2Modifier = 0
	}
	// The major symbols are indexed before in the alphabet
	if g1Grade > g2Grade {
		return  g1
	}
	if g1Grade  == g2Grade {
		if g1Modifier > g2Modifier {
			return g1
		}
		return g2
	}
	return g2
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
