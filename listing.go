package embedtool

import (
	"embed"
)


func ListEmbedFiles(f embed.FS, src string) (files []string, err error) {
	data, err := f.ReadDir(src)
	if err != nil {
		return nil, err
	}
	for _, file := range data {
		if !file.IsDir() {
			files = append(files, file.Name())
		}
	}
	return
}

func ListEmbedDir(f embed.FS, src string) (dirs []string, err error) {
	data, err := f.ReadDir(src)
	if err != nil {
		return nil, err
	}
	for _, file := range data {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}
	return
}

func ListEMbed(f embed.FS, src string) (dirs, files []string, err error) {
	data, err := f.ReadDir(src)
	if err != nil {
		return nil, nil, err
	}
	for _, file := range data {
		if !file.IsDir() {
			files = append(files, file.Name())
		} else {
			dirs = append(dirs, file.Name())
		}
	}
	return
}
