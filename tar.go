package main

import (
	"bytes"
	"embed"
	"io"
	"archive/tar"
	//"io/ioutil"
)

type Tar struct {
	Reader *tar.Reader
}

func (z *Tar)GetBytes(e embed.FS, src string) (b []byte, err error) {
	return e.ReadFile(src)
}


func (z *Tar)GetReader(body []byte, size int64) (err error) {
	tarReader := tar.NewReader(bytes.NewReader(body))
	z.Reader = tarReader
	return nil
}

/*
func (z *Tar) ReadZipFile(zf *tar.File) ([]byte, error) {
    f, err := zf.Open()
    if err != nil {
        return nil, err
    }
    defer f.Close()
    return ioutil.ReadAll(f)
}
*/


func (z *Tar) ParseReader() (array []byte, err error) {
	for {
		_, err := z.Reader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		/*path := filepath.Join(target, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return err
			}
			continue
		}*/
	}
	return nil, nil
}
