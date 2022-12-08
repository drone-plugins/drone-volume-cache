package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/drone/drone-cache-lib/storage"
)

type localCache struct {
}

func (s *localCache) Get(path string, dst io.Writer) error {
	src, err := os.Open(path)
	if err != nil {
		return err
	}
	defer src.Close()
	_, err = io.Copy(dst, src)
	return err
}

func (s *localCache) Put(path string, src io.Reader) error {
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	return err
}

func (s *localCache) List(path string) ([]storage.FileEntry, error) {
	var files []storage.FileEntry
	walker := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, storage.FileEntry{
				Path:         path,
				Size:         info.Size(),
				LastModified: info.ModTime(),
			})
		}
		return nil
	}
	_ = filepath.Walk(path, walker)
	return files, nil
}

func (s *localCache) Delete(path string) error {
	return os.Remove(path)
}
