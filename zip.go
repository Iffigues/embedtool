package main

import (
	"bytes"
	"embed"
	"archive/zip"
	"io/ioutil"
)

type Zip struct {
	Reader *zip.Reader
}

func (z *Zip)GetBytes(e embed.FS, src string) (b []byte, err error) {
	return e.ReadFile(src)
}


func (z *Zip)GetReader(body []byte, size int64) (err error) {
	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return err
	}
	z.Reader = zipReader
	return nil
}


func (z *Zip) ReadZipFile(zf *zip.File) ([]byte, error) {
    f, err := zf.Open()
    if err != nil {
        return nil, err
    }
    defer f.Close()
    return ioutil.ReadAll(f)
}


func (z *Zip) ParseReader() (array []byte, err error) {
    for _, zipFile := range z.Reader.File {
        unzippedFileBytes, err := z.ReadZipFile(zipFile)
        if err != nil {
            continue
        }
	array = append(array, unzippedFileBytes...)
        _ = unzippedFileBytes // this is unzipped file bytes
    }
    return array, nil
}
