package browser

import (
	"fmt"
	"os"
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
