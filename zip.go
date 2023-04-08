package embedtool

import (
	"archive/zip"
	"bytes"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Zip struct {
	Reader *zip.Reader
	Writer *zip.Writer
}

type WalkFunc func(path string, info fs.FileInfo, err error) error

func (z *Zip) GetBytes(e embed.FS, src string) (b []byte, err error) {
	return e.ReadFile(src)
}

func (z *Zip) GetReader(body []byte) (err error) {
	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return err
	}
	z.Reader = zipReader
	return nil
}

func (z *Zip) Walk(f embed.FS, src string, fn WalkFunc, rec bool) {
	dires, files, err := ListEMbed(f, src)
	for _, val := range files {
		fn(src+"/"+val, nil, err)
	}
	for _, val := range dires {
		fmt.Println("ezez")
		fn(src+"/"+val, nil, err)
	}
	if rec {
		for _, val := range dires {
			z.Walk(f, src+"/"+val, fn, rec)
		}
	}
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

func (z *Zip) Make(f embed.FS, src, dest string, rec bool) (err error) {
	node, err := f.Open(src)
	if err != nil {
		return err
	}
	isDir, err := FsFileIsDir(node)
	if err != nil {
		return err
	}

	zipf, err := os.Create(dest)
	if err != nil {
		panic(err)
	}
	defer zipf.Close()
	zipw := zip.NewWriter(zipf)
	defer zipw.Close()
	if isDir {

		z.Walk(f, src, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			isDir, err := IsDir(path, f)
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Println(path, isDir)
			relpath, err := filepath.Rel(src, path)
			if err != nil {
				return err
			}
			h := &zip.FileHeader{
				Name:   relpath,
				Method: zip.Deflate,
			}
			if isDir {
				h.Name += "/"
				h.Method = zip.Store
				_, err = zipw.CreateHeader(h)
				if err != nil {
					return err
				}
				return nil
			}
			w, err := zipw.CreateHeader(h)
			if err != nil {
				return err
			}
			if !isDir {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()

				_, err = io.Copy(w, file)
				if err != nil {
					return err
				}
			}
			return nil
		}, rec)
	} else {
		f.ReadFile(src)
		h := &zip.FileHeader{
			Name:   filepath.Base(src),
			Method: zip.Deflate,
		}
		w, err := zipw.CreateHeader(h)
		if err != nil {
			panic(err)
		}
		fn, err := f.ReadFile(src)
		if err != nil {
			return err
		}
		_, err = w.Write(fn)
		if err != nil {
			panic(err)
		}
	}
	return nil
}
