package embedtool

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"embed"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Tar struct {
	Reader *tar.Reader
	Writer *tar.Writer
}

func (z *Tar) GetBytes(e embed.FS, src string) (b []byte, err error) {
	return e.ReadFile(src)
}

func (z *Tar) GetReader(body []byte) (err error) {
	tarReader := tar.NewReader(bytes.NewReader(body))
	z.Reader = tarReader
	return nil
}

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
		bs, err := ioutil.ReadAll(z.Reader)
		if err != nil {
			return nil, err
		}
		files[header.Name] = bs
	}
	return
}

func (z *Tar) Walk(f embed.FS, src string, fn WalkFunc, rec bool) {
	dires, files, err := ListEMbed(f, src)
	for _, val := range files {
		fn(src+"/"+val, nil, err)
	}
	for _, val := range dires {
		fn(src+"/"+val, nil, err)
	}
	if rec {
		for _, val := range dires {
			z.Walk(f, src+"/"+val, fn, rec)
		}
	}
}

func (z *Tar) IsDir(f embed.FS, src, dest string, rec bool, zipw *zip.Writer) error {
	z.Walk(f, src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		isDir, err := IsDir(path, f)
		if err != nil {
			return err
		}

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
	return nil
}
func (z *Tar) IsFile(f embed.FS, src, dest string, rec bool, zipw *zip.Writer) error {
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
	return nil
}

func (z *Tar) Make(f embed.FS, src, dest string, rec bool) (err error) {
	node, err := f.Open(src)
	if err != nil {
		return err
	}
	defer node.Close()

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
		return z.IsDir(f, src, dest, rec, zipw)
	} else {
		return z.IsFile(f, src, dest, rec, zipw)
	}
}
