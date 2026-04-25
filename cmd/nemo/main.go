package main

import (
	"fmt"
	"os"

	"github.com/hymkor/nemo"
)

func main() {
	if err := nemo.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
