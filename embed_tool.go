package main

import (
	"embed"
	"io"
	"io/fs"
	"os"
)

func isDir( src string, f embed.FS) (dir bool, err error) {
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

func fsFileIsDir(f fs.File) (dir bool, err error){
	state, err := f.Stat()
	if err != nil {
		return false, err
	}
	return state.IsDir(), nil
}

func fsGetFileMod(f fs.File) (m fs.FileMode, err error) {
	state, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return state.Mode(), nil
}

func createFile(file  fs.File, dest string) (err error) {
	destination, err := os.Create(dest)
	if err != nil {
		return err
        }
	buf := make([]byte, 3000)
        for {
                n, err := file.Read(buf)
                if err != nil && err != io.EOF {
                        return err
                }
                if n == 0 {
                        break
                }
                if _, err := destination.Write(buf[:n]); err != nil {
                        return err
                }
        }
	return nil
}

func createFileWithPerm(file  fs.File, dest string, mode fs.FileMode) (err error) {
	destination, err := os.Create(dest)
	if err != nil {
		return err
        }
	err = os.Chmod(dest, mode)
	if err != nil {
		return err
	}
	buf := make([]byte, 3000)
        for {
                n, err := file.Read(buf)
                if err != nil && err != io.EOF {
                        return err
                }
                if n == 0 {
                        break
                }
                if _, err := destination.Write(buf[:n]); err != nil {
                        return err
                }
        }
	return nil
}

func copyFile(dest, src string, overwrite bool, f embed.FS) (err error) {
	if _, err := os.Stat(dest); err == nil {
		if !overwrite {
			return err
		}
	}
	file, err := f.Open(src)
	if err != nil {
		return err
	}
	if isDir, err := fsFileIsDir(file); (err != nil) || (isDir) {
		return err
	}
	return createFile(file, dest)
}

func copyFileSetPerm(dest, src string, mode fs.FileMode, overwrite bool, f embed.FS) (err error) {
	if _, err := os.Stat(dest); err == nil {
		if !overwrite {
			return err
		}
	}
	file, err := f.Open(src)
	if err != nil {
		return err
	}
	if isDir, err := fsFileIsDir(file); (err != nil) || (isDir) {
		return err
	}
	return createFileWithPerm(file, dest, mode)
}
