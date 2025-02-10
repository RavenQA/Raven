package progress

import (
	"io"
)

type ProgressFunc func(int, int)

type ProgressReadCloser struct {
	io.Reader
	io.Closer
	total        int
	progress     int
	progressFunc ProgressFunc
}

func NewProgressReadCloser(r io.ReadCloser, total int, f ProgressFunc) *ProgressReadCloser {
	pr := ProgressReadCloser{
		Reader:       r,
		Closer:       r,
		total:        total,
		progressFunc: f,
	}
	return &pr
}

func (r *ProgressReadCloser) Read(p []byte) (int, error) {
	n, err := r.Reader.Read(p)
	if err != nil {
		return n, err
	}
	r.progress += n
	if r.progressFunc != nil {
		r.progressFunc(r.progress, r.total)
	}
	return n, nil
}

func (r *ProgressReadCloser) Close() error {
	return r.Closer.Close()
}
