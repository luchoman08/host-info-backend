package infraestructures

import (
	"github.com/badoux/goscraper"
	"net/url"
)

type GoScraperHandler struct {
}

func (scraper *GoScraperHandler) Scrape(url url.URL, maxRedirect int) (*goscraper.Document, error) {
	return goscraper.Scrape(url.String(), maxRedirect)
}
