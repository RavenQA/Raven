package fetch

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/soikes/raven/pkg/platform"
	"github.com/soikes/raven/pkg/progress"
	"golang.org/x/text/language"
)

// The download API must be used to fetch the latest version.
const releaseApiUrl = "https://download.mozilla.org/"
const productKey = `product`
const platformKey = `os`
const langKey = `lang`
const platformWindows = `win64`
const platformMac = `osx`
const platformLinux64 = `linux64`
const platformLinux = `linux`
const productFirefoxLatest = `firefox-latest`
const macInstallerExt = `.dmg`

var platforms = map[platform.Platform]string{
	platform.PlatformWindows: platformWindows,
	platform.PlatformMac:     platformMac,
	platform.PlatformLinux64: platformLinux64,
	platform.PlatformLinux:   platformLinux,
}

// For other versions, the FTP directory can be used.
const releaseFtpUrl = `https://ftp.mozilla.org/pub/firefox/releases/`
const platformDirWindows = `win32`
const platformDirMac = `mac`
const platformDirLinux64 = `linux-i686`
const platformDirLinux = `linux-x86_64`
const versionDirLatest = `latest`

var platformDirs = map[platform.Platform]string{
	platform.PlatformWindows: platformDirWindows,
	platform.PlatformMac:     platformDirMac,
	platform.PlatformLinux64: platformDirLinux64,
	platform.PlatformLinux:   platformDirLinux,
}

type Config struct {
	Platform     platform.Platform
	ProgressFunc progress.ProgressFunc
	DownloadPath string
}

func (c *Config) Fetch(version string, lang language.Tag) error {
	var u string
	var err error
	switch version {
	case `latest`:
		u, err = buildApiUrl(c.Platform, lang)
		if err != nil {
			return err
		}
	default:
		u, err = buildFtpUrl(version, c.Platform, lang)
		if err != nil {
			return err
		}
	}
	return fetchRelease(u, c.DownloadPath, c.ProgressFunc)
}

func buildApiUrl(platform platform.Platform, lang language.Tag) (string, error) {
	var u *url.URL
	u, err := url.Parse(releaseApiUrl)
	if err != nil {
		return "", err
	}
	p, ok := platforms[platform]
	if !ok {
		return "", fmt.Errorf("unsupported platform %T", platform)
	}
	params := url.Values{}
	params.Add(productKey, productFirefoxLatest)
	params.Add(platformKey, p)
	params.Add(langKey, lang.String())
	u.RawQuery = params.Encode()
	return u.String(), nil
}

func buildFtpUrl(version string, platform platform.Platform, lang language.Tag) (string, error) {
	var u *url.URL
	u, err := url.Parse(releaseFtpUrl)
	if err != nil {
		return "", err
	}
	name := fmt.Sprintf("%s %s%s", `Firefox`, version, macInstallerExt)
	u = u.JoinPath(version, platformDirs[platform], lang.String(), name)
	return u.String(), nil
}

func fetchRelease(u string, outpath string, progressFunc progress.ProgressFunc) error {
	f, err := os.Create(outpath)
	if err != nil {
		return err
	}
	fmt.Println("getting: ", u)
	rsp, err := http.Get(u)
	if err != nil {
		return err
	}
	var r io.ReadCloser
	if progressFunc != nil {
		r = progress.NewProgressReadCloser(rsp.Body, int(rsp.ContentLength), progressFunc)
	} else {
		r = rsp.Body
	}
	defer r.Close()
	_, err = io.Copy(f, r)
	if err != nil {
		return err
	}
	return nil
}
