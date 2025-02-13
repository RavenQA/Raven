package browser

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Version struct {
	Id          int
	Major       int
	Minor       *int
	Build       *int
	Patch       *int
	ReleaseDate time.Time
}

func VersionFromString(ver string) (Version, error) {
	parts := strings.Split(ver, ".")
	var v Version
	if len(parts) > 0 {
		major, err := strconv.Atoi(parts[0])
		if err != nil {
			return v, err
		}
		v.Major = major
	}
	if len(parts) > 1 {
		minor, err := strconv.Atoi(parts[1])
		if err != nil {
			return v, err
		}
		v.Minor = &minor
	}
	if len(parts) > 2 {
		build, err := strconv.Atoi(parts[2])
		if err != nil {
			return v, err
		}
		v.Build = &build
	}
	if len(parts) > 3 {
		patch, err := strconv.Atoi(parts[3])
		if err != nil {
			return v, err
		}
		v.Patch = &patch
	}
	return v, nil
}

type Browser struct {
	Version
	Name      string
	Path      string
	Available bool
}

func (b *Browser) Identifier() string {
	return fmt.Sprintf("%s-%s", b.Name, b.VersionString())
}

func (b *Browser) VersionString() string {
	v := fmt.Sprint(b.Major)
	if b.Minor != nil {
		v += fmt.Sprintf(".%d", *b.Minor)
	}
	if b.Build != nil {
		v += fmt.Sprintf(".%d", *b.Build)
	}
	if b.Patch != nil {
		v += fmt.Sprintf(".%d", *b.Patch)
	}
	return v
}

// IsAvailable checks if the Browser has been downloaded to the local filesystem.
func (b *Browser) IsAvailable() (bool, error) {
	if b.Path == "" {
		return false, nil
	}
	_, err := os.Stat(b.Path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
