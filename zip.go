package embedtool

import (
	"bytes"
	"embed"
	"archive/zip"
	"io/ioutil"
	"strings"
)

type Zip struct {
	Reader *zip.Reader
}

func (z *Zip)GetBytes(e embed.FS, src string) (b []byte, err error) {
	return e.ReadFile(src)
}


func (z *Zip)GetReader(body []byte) (err error) {
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


func (z *Zip) ParseReader() (files map[string][]byte, err error) {
    files = make(map[string][]byte)
    for _, zipFile := range z.Reader.File {
	unzippedFileBytes, err := z.ReadZipFile(zipFile)
        if err != nil {
            continue
        }
	if !strings.HasSuffix(zipFile.Name, "/") {
		files[zipFile.Name] = unzippedFileBytes
	}

    }
    return
}
