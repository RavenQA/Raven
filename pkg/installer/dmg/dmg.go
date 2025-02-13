package dmg

type Dmg struct {
	ImagePath  string
	MountPoint string
	AppPath    string
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

func WithAppPath(path string) Option {
	return func(dmg *Dmg) {
		dmg.AppPath = path
	}
}

func NewDmg(opts ...Option) *Dmg {
	d := Dmg{}
	for _, opt := range opts {
		opt(&d)
	}
	return &d
}
