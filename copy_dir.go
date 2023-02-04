package embedtool


import (
	"embed"
	"os"
	"errors"
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
	return
}

func CopyEmbededDir(f embed.FS, src, dest string, recursive bool) (err error) {

	if isDir, err := DirExist(dest); err != nil || !isDir {
		return err
	}
	isDir, err := IsDir(src, f)

	if err != nil {
		return err
	}

	if  !isDir {
		return errors.New("src is not a directory")
	}

	return copyEmbededDir(f, src, dest, recursive)
}
