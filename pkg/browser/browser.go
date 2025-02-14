package browser

import (
	"fmt"
	"os"
	"time"
)

type Browser struct {
	Product     string
	Version     string
	ReleaseDate time.Time // TODO: Create custom unmarshaler for time.Time. Also how to use a time type in the frontend?
	InstallPath string
}

func (b *Browser) Identifier() string {
	return fmt.Sprintf("%s-%s", b.Product, b.Version)
}

func Identifier(product, version string) string {
	return fmt.Sprintf("%s-%s", product, version)
}

// IsAvailable checks if the Browser has been downloaded to the local filesystem.
func (b *Browser) IsAvailable() (bool, error) {
	if b.InstallPath == "" {
		return false, nil
	}
	_, err := os.Stat(b.InstallPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Firefox() Browser {
	return Browser{
		Product: ProductFirefox,
	}
}

const ProductFirefox = `Firefox`
