package archive

import (
	"io"

	"github.com/Sirupsen/logrus"
	"github.com/drone/drone-cache-lib/archive"
	"github.com/replicon/fast-archiver/falib"
)

type fastArchiver struct{}

// New returns a new archiver based on fast-archiver.
func New() archive.Archive {
	return new(fastArchiver)
}

func (fa *fastArchiver) Pack(dirs []string, w io.Writer) error {
	a := falib.NewArchiver(w)
	a.BlockSize = 4096
	a.DirScanQueueSize = 128
	a.FileReadQueueSize = 128
	a.BlockQueueSize = 128
	a.DirReaderCount = 16
	a.FileReaderCount = 16

	a.Logger = &logger{}
	for _, dir := range dirs {
		a.AddDir(dir)
	}
	return a.Run()
}

func (fa *fastArchiver) Unpack(dst string, r io.Reader) error {
	a := falib.NewUnarchiver(r)
	a.Logger = &logger{}
	return a.Run()
}

type logger struct{}

func (l *logger) Verbose(v ...interface{}) {
	logrus.Debug(v...)
}

func (l *logger) Warning(v ...interface{}) {
	logrus.Warn(v...)
}
