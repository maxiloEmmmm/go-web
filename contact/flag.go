package contact

import (
	"flag"
	"fmt"
	"os"
)

func InitFlag() {
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	}
}

var (
	help bool
)

func init() {
	flag.BoolVar(&help, "h", false, "help")
	flag.StringVar(&ConfigPath, "c", "./config.yaml", "config path")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `go-web
Usage: app [-c filename]

Options:
`)
	flag.PrintDefaults()
}
