package walker

import (
	"io/fs"
	"path/filepath"
)

type set[T comparable] = map[T]struct{}

type WalkerOptions struct {
	Extensions  set[string]
	IgnoredDirs set[string]
}

// walks and retrieves files
func WalkThisWay(dir string, ext set[string], ignores set[string]) (files []string, err error) {

	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if d.IsDir() {
			if _, ok := ignores[d.Name()]; ok {
				return filepath.SkipDir
			}
		} else {
			e := filepath.Ext(path)
			if _, ok := ext[e]; ok {
				return nil
			}
			a, _ := filepath.Abs(path)
			files = append(files, a)
		}
		return nil
	})
	return
}
