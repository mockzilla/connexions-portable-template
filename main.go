package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/mockzilla/mockzilla/v2/pkg/portable"
)

var version = "dev"

//go:embed openapi static
var content embed.FS

func main() {
	if len(os.Args) == 2 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Println(version)
		return
	}
	os.Exit(portable.RunFS(content, os.Args[1:]))
}
