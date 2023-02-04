package embedtool

import (
	"io/fs"
	"os"
	"embed"
)

func IsDir( src string, f embed.FS) (dir bool, err error) {
	node, err := f.Open(src)
	if err != nil {
		return false, err
	}
	state, err := node.Stat()
	if err != nil {
		return false, err
	}
	return state.IsDir(), nil
}

func CopyFile(dest, src  string, overwrite bool, f embed.FS) (err error) {
	if _, err := os.Stat(dest); err == nil {
		if !overwrite {
			return err
		}
	}
	file, err := f.Open(src)
	if err != nil {
		return err
	}
	if isDir, err := FsFileIsDir(file); (err != nil) || (isDir) {
		println(isDir)
		return err
	}
	return CreateFile(file, dest)
}

func copyFileSetPerm(dest, src  string, mode fs.FileMode, overwrite bool, f embed.FS) (err error) {
	if _, err := os.Stat(dest); err == nil {
		if !overwrite {
			return err
		}
	}
	file, err := f.Open(src)
	if err != nil {
		return err
	}
	if isDir, err := FsFileIsDir(file); (err != nil) || (isDir) {
		return err
	}
	return CreateFileWithPerm(file, dest, mode)
}
