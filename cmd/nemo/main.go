package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hymkor/nemo"
)

var (
	flagShowControl = flag.Bool("show-control", false, "display control characters as \\xNN")
	flagStripCr     = flag.Bool("strip-cr", false, "remove carriage returns")
)

func main() {
	flag.Parse()

	app := &nemo.Application{
		ShowControl: *flagShowControl,
		StripCr:     *flagStripCr,
	}
	if err := app.Run(flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
