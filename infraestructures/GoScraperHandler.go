package infraestructures

import (
	"github.com/badoux/goscraper"
	"net/url"
)
// GoScraperHandler implements all methods for scrape an url
type GoScraperHandler struct {
}
// Scrape return all the info scraped from a given URL
func (scraper *GoScraperHandler) Scrape(url url.URL, maxRedirect int) (*goscraper.Document, error) {
	return goscraper.Scrape(url.String(), maxRedirect)
}
