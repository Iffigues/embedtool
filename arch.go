package embedtool

import (
	"embed"
	"fmt"
)

type Archive interface {
	GetBytes(e embed.FS, src string) ([]byte, error)
	GetReader(r []byte) error
	ParseReader() (map[string][]byte, error)
}

type MakeArchive interface {
	make(e embed.FS, src, dest string)
}

func GetArchive(arch Archive, e embed.FS, src string) (map[string][]byte, error) {
	data, err := arch.GetBytes(e, src)
	if err != nil {
		return nil, err
	}
	err = arch.GetReader(data)
	if err != nil {
		return nil, err
	}
	return arch.ParseReader()
}

func tete(e embed.FS, a string) {
	zz := &Tar{}
	rrr, zzz := GetArchive(zz, e, a)
	fmt.Println(zzz)
	for ee, i := range rrr {
		fmt.Println(ee, "\n", string(i))
	}
}
