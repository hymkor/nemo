package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hymkor/struct2flag"

	"github.com/hymkor/nemo"
)

func main() {
	var app nemo.Application

	struct2flag.BindDefault(&app)
	flag.Parse()
	if err := app.Run(flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
