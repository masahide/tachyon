package main

import (
	"os"

	"github.com/masahide/tachyon"
	_ "github.com/masahide/tachyon/net"
	_ "github.com/masahide/tachyon/package"
	_ "github.com/masahide/tachyon/procmgmt"
)

var Release string
var version string

func main() {
	if Release != "" {
		tachyon.Release = Release
	}
	if version != "" {
		tachyon.Version = version
	}

	os.Exit(tachyon.Main(os.Args))
}
