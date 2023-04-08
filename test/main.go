package main

import (
	"embed"
	"embedtool"
	"fmt"
)

//go:embed hello
var content embed.FS

func main() {
	z := &embedtool.Tar{}
	fmt.Println(embedtool.IsDir("hello/hello", content))
	fmt.Println(z.Make(content, "hello", "dir.tar", false))
	//fmt.Println(z.Make(content, "hello", "ici"))
	//fmt.Println(embedtool.CopyEmbededDir(content, "hello", "lml", true))
}
