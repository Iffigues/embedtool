package main

import (
	"embed"
	"fmt"
)

type Archive interface {
	GetBytes(e embed.FS, src string) ([]byte, error)
	GetReader(r []byte, size int64) ( error)
	ParseReader() ([]byte, error)
}


func GetArchive(arch Archive, e embed.FS, src string) ([]byte, error) {
	data, err := arch.GetBytes(e, src)
	if err != nil {
		return nil, err
	}
	err = arch.GetReader(data, int64(len(data)))
	if err != nil {
		return nil, err
	}
	return arch.ParseReader()
}

func tete(e embed.FS, a string) {
	zz := &Zip{}
	rrr, zzz := GetArchive(zz, e, a)
	fmt.Println(string(rrr), zzz)
}
