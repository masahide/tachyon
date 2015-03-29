package main

import (
	"github.com/masahide/tachyon"
	_ "github.com/masahide/tachyon/net"
	_ "github.com/masahide/tachyon/package"
	_ "github.com/masahide/tachyon/procmgmt"
	"os"
)

var Release string

func main() {
	if Release != "" {
		tachyon.Release = Release
	}

	os.Exit(tachyon.Main(os.Args))
}
