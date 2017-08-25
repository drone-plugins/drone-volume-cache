package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/drone/drone-cache-lib/archive/tar"
	"github.com/drone/drone-cache-lib/cache"
	"github.com/drone/drone-cache-lib/storage"
)

type plugin struct {
	mount    []string
	path     string
	file     string
	fallback string
	rebuild  bool
	restore  bool
	flush    bool
	ttl      time.Duration

	storage storage.Storage
}

// Exec runs the plugin
func (p *plugin) Exec() error {
	encoder := tar.New()
	cacher := cache.New(p.storage, encoder)

	if p.rebuild {
		dir, _ := filepath.Split(p.file)
		os.MkdirAll(dir, 0700)

		logrus.Infof("Rebuilding cache at %s", p.file)
		if err := cacher.Rebuild(p.mount, p.file); err == nil {
			logrus.Infof("Cache rebuilt")
		} else {
			logrus.Warnf("Error rebuilding cache. %s", err)
		}
	}

	if p.restore {
		logrus.Infof("Restoring cache from %s", p.file)
		if err := cacher.Restore(p.file, p.fallback); err == nil {
			logrus.Info("Cache restored")
		} else {
			logrus.Warningf("Error restoring cache", err)
			p.storage.Delete(p.file)
			p.storage.Delete(p.fallback)
		}
	}

	if p.flush && p.ttl > 0 {
		logrus.Infof("Purging cached items older then %v", p.ttl)
		flusher := cache.NewFlusher(p.storage, testExpired(p.ttl))
		if ferr := flusher.Flush(p.path); ferr != nil {
			logrus.Warnf("Error purging cache. %s", ferr)
		}
	}

	return nil
}

// Check if older then x days (default 30 days)
func testExpired(ttl time.Duration) cache.DirtyFunc {
	return func(file storage.FileEntry) bool {
		return file.LastModified.Before(time.Now().Add(-ttl))
	}
}
