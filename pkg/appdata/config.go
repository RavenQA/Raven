package appdata

import (
	"os"
	"path/filepath"
)

type Config struct {
	Dir    string
	TmpDir string
}

func NewConfig() (*Config, error) {
	d, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	d = filepath.Join(d, `Raven`)
	err = os.MkdirAll(d, 0700)
	if err != nil {
		return nil, err
	}
	return &Config{
		Dir:    d,
		TmpDir: os.TempDir(),
	}, nil
}
