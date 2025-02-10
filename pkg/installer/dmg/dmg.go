package dmg

import (
	"fmt"
	"os/exec"
)

type Dmg struct {
	ImagePath  string
	MountPoint string
}

type Option func(dmg *Dmg)

func WithMountPoint(path string) Option {
	return func(dmg *Dmg) {
		dmg.MountPoint = path
	}
}

func WithImagePath(path string) Option {
	return func(dmg *Dmg) {
		dmg.ImagePath = path
	}
}

func NewDmg(opts ...Option) *Dmg {
	d := Dmg{}
	for _, opt := range opts {
		opt(&d)
	}
	return &d
}

func (d Dmg) Attach(outpath string) error {
	c := exec.Command(dmgCmd, append(d.attachArgs(), outpath)...)
	err := c.Run()
	if err != nil {
		if xerr, ok := err.(*exec.ExitError); ok {
			fmt.Printf("%s\n", xerr.Stderr)
		}
		return err
	}
	return nil
}

func (d Dmg) Detach() error {
	c := exec.Command(dmgCmd, d.detachArgs()...)
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
