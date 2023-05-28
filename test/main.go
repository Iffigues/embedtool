package main

import (
	"embed"
	"embedtool"
	"fmt"
)

//go:embed hello
var content embed.FS

func main() {
	fmt.Println("ezeez")
	fmt.Println(embedtool.CopyEmbededDir(content, "hello", "./lml", true))
}
