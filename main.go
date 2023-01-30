package main

import (
	"embed"
	"fmt"
)
//go:embed hello
var f embed.FS

func main() {
	tete(f, "hello/a.zip")
	fmt.Println("end")
}
