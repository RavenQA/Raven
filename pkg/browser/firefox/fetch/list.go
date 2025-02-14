package fetch

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/soikes/raven/pkg/browser"
)

const productDetailsUrl = `https://product-details.mozilla.org/1.0/firefox.json`

func FetchBrowserList() ([]browser.Browser, error) {
	rl, err := fetchVersions()
	if err != nil {
		return nil, err
	}
	var bl []browser.Browser
	for _, release := range rl.Releases {
		if release.Category == "major" || release.Category == "stability" {
			if isValidVersion(release.Version) {
				d, err := parseReleaseDate(release.Date)
				if err != nil {
					log.Printf("failed to parse firefox %s release date: %s", release.Version, release.Date)
					continue
				}
				b := browser.Firefox()
				b.Version = release.Version
				b.ReleaseDate = d
				bl = append(bl, b)
			}
		}
	}
	return bl, nil
}

type releaseList struct {
	Releases map[string]productDetails
}

type productDetails struct {
	BuildNumber int
	Category    string
	Date        string
	Version     string
}

func fetchVersions() (releaseList, error) {
	var rl releaseList
	resp, err := http.Get(productDetailsUrl)
	if err != nil {
		return rl, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&rl)
	if err != nil {
		return rl, err
	}
	return rl, nil
}

func parseReleaseDate(in string) (time.Time, error) {
	return time.Parse(time.DateOnly, in)
}

const pattern = `^(\d*\.){0,3}\d*$`

var verRxp = regexp.MustCompile(pattern)

func isValidVersion(in string) bool {
	return verRxp.MatchString(in)
}
