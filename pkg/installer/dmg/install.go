package dmg

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func (d *Dmg) Install(ctx context.Context) (err error) {
	err = d.attach(ctx, d.ImagePath)
	if err != nil {
		return err
	}
	defer func() {
		derr := d.detach(ctx)
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
	e, err := exists(d.AppPath)
	if err != nil {
		return err
	}
	if e {
		return fmt.Errorf("%s already exists", d.AppPath)
	}
	err = os.CopyFS(d.AppPath, os.DirFS(app))
	if err != nil {
		return err
	}
	return err
}

func (d Dmg) attach(ctx context.Context, outpath string) error {
	c := exec.CommandContext(ctx, dmgCmd, append(d.attachArgs(), outpath)...)
	err := c.Run()
	if err != nil {
		if xerr, ok := err.(*exec.ExitError); ok {
			fmt.Printf("%s\n", xerr.Stderr)
		}
		return err
	}
	return nil
}

func (d Dmg) detach(ctx context.Context) error {
	c := exec.CommandContext(ctx, dmgCmd, d.detachArgs()...)
	err := c.Run()
	if err != nil {
		if xerr, ok := err.(*exec.ExitError); ok {
			fmt.Printf("%s\n", xerr.Stderr)
		}
		return err
	}
	return nil
}

const dmgCmd = `hdiutil`

func (d Dmg) attachArgs() []string {
	return append([]string{
		`attach`,
		`-readonly`,
		`-noautofsck`,
		`-noautoopen`,
		// `-nobrowse`,
	},
		d.mountPointArg()...,
	)
}

func (d Dmg) mountPointArg() []string {
	return []string{
		`-mountpoint`,
		d.MountPoint,
	}
}

func (d Dmg) detachArgs() []string {
	return []string{
		`detach`,
		d.MountPoint,
	}
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
