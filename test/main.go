package main

import (
	"embed"
	"fmt"
	"embedtool"
)


//go:embed hello
var content embed.FS

func main() {
	fmt.Println(embedtool.CopyEmbededDir(content, "hello", "lml", true))
}
