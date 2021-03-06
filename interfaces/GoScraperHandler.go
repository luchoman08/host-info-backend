package interfaces

import (
	"github.com/luchoman08/goscraper"
	"net/url"
)

// GoScraperHandler provide the methods for access the go scraper library
type GoScraperHandler interface {
	Scrape(url url.URL, maxRedirect int) (*goscraper.Document, error)
}
