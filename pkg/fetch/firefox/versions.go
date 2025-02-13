package firefox

import (
	"encoding/json"
	"net/http"
	"regexp"
	"time"

	"github.com/soikes/raven/pkg/browser"
)

const productDetailsUrl = `https://product-details.mozilla.org/1.0/firefox.json`

type ReleaseList struct {
	Releases map[string]productDetails
}

type productDetails struct {
	BuildNumber int
	Category    string
	Date        string
	Version     string
}

func GetVersions() ([]browser.Version, error) {
	var versions []browser.Version
	rl, err := fetchVersions()
	if err != nil {
		return nil, err
	}
	for _, release := range rl.Releases {
		if release.Category == "major" || release.Category == "stability" {
			if isValidVersion(release.Version) {
				v, err := browser.VersionFromString(release.Version)
				if err != nil {
					return nil, err
				}
				t, err := time.Parse(time.DateOnly, release.Date)
				if err != nil {
					return nil, err
				}
				v.ReleaseDate = t
				versions = append(versions, v)
			}
		}
	}
	return versions, nil
}

const pattern = `^(\d*\.){0,3}\d*$`

var verRxp = regexp.MustCompile(pattern)

func isValidVersion(in string) bool {
	return verRxp.MatchString(in)
}

func fetchVersions() (ReleaseList, error) {
	var rl ReleaseList
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
