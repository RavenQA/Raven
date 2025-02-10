package dmg

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func (d *Dmg) ExtractApp(image, out string) (err error) {
	err = d.Attach(image)
	if err != nil {
		return err
	}
	defer func() {
		derr := d.Detach()
		if derr != nil {
			err = derr
		}
	}()
	var app string
	entries, err := os.ReadDir(d.MountPoint)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() && filepath.Ext(entry.Name()) == ".app" {
			app = filepath.Join(d.MountPoint, entry.Name())
			break
		}
	}
	if app == "" {
		return errors.New("no .app in mounted .dmg")
	}
	e, err := exists(out)
	if err != nil {
		return err
	}
	if e {
		return fmt.Errorf("%s already exists", out)
	}
	err = os.CopyFS(out, os.DirFS(app))
	if err != nil {
		return err
	}
	return err
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
