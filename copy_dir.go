package embedtool

import (
	"embed"
	"os"
	"path/filepath"
)

func DirExist(path string) (bool, error) {
	files, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return files.IsDir(), nil
}

func copyEmbededDir(f embed.FS, src, dest string, recursive bool) (err error) {
	dires, files, err := ListEMbed(f, src)
	if err != nil {
		return err
	}

	for _, val := range files {
		CopyFile(filepath.Join(dest, "/", val), filepath.Join(src, "/", val), false, f)
	}

	for _, val := range dires {
		os.Mkdir(filepath.Join(dest, "/", val), 777)
	}
	if recursive {
		for _, val := range dires {
			copyEmbededDir(f, filepath.Join(src, "/", val), filepath.Join(dest, "/", val), true)
		}
	}
	return
}

func CopyEmbededDir(f embed.FS, src, dest string, recursive bool) (err error) {

	if isDir, err := DirExist(dest); err != nil || !isDir {
		return err
	}
	return copyEmbededDir(f, src, dest, recursive)
}
