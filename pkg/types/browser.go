package types

import (
	"time"

	"github.com/soikes/raven/pkg/browser"
)

type BrowserListItem struct {
	Version     string
	Name        string
	ReleaseDate string
	Path        string
	Available   bool
}

func NewBrowserListItem(b browser.Browser) BrowserListItem {
	li := BrowserListItem{
		Version:     b.VersionString(),
		Name:        b.Name,
		ReleaseDate: b.ReleaseDate.Format(time.DateOnly),
		Path:        b.Path,
		Available:   b.Available,
	}
	return li
}
