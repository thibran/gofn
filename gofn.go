package main

import (
	"flag"
	"os"

	"github.com/thibran/gofn/bindir"
	"github.com/thibran/gofn/run"
)

var noFlags = false

func main() {
	flag.Usage = printHelp
	// print help if no arguments have been passed
	if len(os.Args) == 1 {
		noFlags = true
		os.Args = append(os.Args, "-h")
	}
	// parse
	f := run.ParseAndCreateFunction()
	if f.Debug {
		run.PrintDebug(f)
		os.Exit(0)
	}
	// bild and or run gofn-function
	bindir, err := bindir.Path()
	if err != nil {
		panic(err)
	}
	r := run.NewRun(bindir)
	if err = r.Exec(f); err != nil {
		os.Exit(1)
	}
}
