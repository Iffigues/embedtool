package main

import (
	"bytes"
	"embed"
	"io"
	"archive/tar"
	"io/ioutil"
)

type Tar struct {
	Reader *tar.Reader
}

func (z *Tar)GetBytes(e embed.FS, src string) (b []byte, err error) {
	return e.ReadFile(src)
}


func (z *Tar)GetReader(body []byte) (err error) {
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


func (z *Tar) ParseReader() (files map[string][]byte, err error) {
	files = make(map[string][]byte)
	for {
		header, err := z.Reader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		info := header.FileInfo()
		if info.IsDir() {
			continue
		}
		bs, err  := ioutil.ReadAll(z.Reader)
		if err != nil {
			return nil, err
		}
		files[header.Name] = bs
	}
	return
}
