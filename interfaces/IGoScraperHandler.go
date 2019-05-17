package interfaces

import (
	"github.com/badoux/goscraper"
	"net/url"
)

type IGoScraperHandler interface {
	Scrape(url url.URL, maxRedirect int) (*goscraper.Document, error)
}

