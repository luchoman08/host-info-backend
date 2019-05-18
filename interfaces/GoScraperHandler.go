package interfaces

import (
	"github.com/badoux/goscraper"
	"net/url"
)

type GoScraperHandler interface {
	Scrape(url url.URL, maxRedirect int) (*goscraper.Document, error)
}
