package embedtool

import (
	"os"
	"io"
	"io/fs"
)


func FsFileIsDir(f fs.File) (dir bool, err error){
	state, err := f.Stat()
	if err != nil {
		return false, err
	}
	return state.IsDir(), nil
}

func FsGetFileMod(f fs.File) (m fs.FileMode, err error) {
	state, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return state.Mode(), nil
}

func CreateFile(file  fs.File, dest string) (err error) {
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

func CreateFileWithPerm(file  fs.File, dest string, mode fs.FileMode) (err error) {
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
