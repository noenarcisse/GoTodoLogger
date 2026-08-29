package walker

import (
	"io/fs"
	"path/filepath"
	"slices"
)

type set[T comparable] map[T]struct{}

type WalkerOptions struct {
	Extensions  set[string]
	IgnoredDirs set[string]
}

// walks and retrieves files
func WalkThisWay(dir string, ext []string, ignores []string) (files []string, err error) {

	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if d.IsDir() {
			if slices.Contains(ignores, d.Name()) { //todo passer en set / map pour O(1) le contains
				return filepath.SkipDir
			}
		} else {
			e := filepath.Ext(path)
			if !(slices.Contains(ext, e)) { //todo passer en set / map pour O(1) le contains
				return nil
			}

			a, _ := filepath.Abs(path)
			files = append(files, a)
		}
		return nil
	})
	return
}
