//go:build run

package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-colorable"

	"github.com/nyaosorg/go-ttyadapter/tty8pe"

	"github.com/hymkor/go-generics-list"
	"github.com/hymkor/nemo/pager"
)

// TextElement represents one line in the pager
type TextElement struct {
	Text string
}

// Display is called by pager to render each line
func (t TextElement) Display(screenWidth int) string {
	return t.Text
}

func main() {
	// Create a linked list of lines
	lines := list.New[TextElement]()
	for i := 0; i < 2000; i++ {
		lines.PushBack(TextElement{Text: fmt.Sprintf("<%d>", i)})
	}
	// Create pager
	pg := &pager.Pager[TextElement]{
		Status: func(session *pager.Session[TextElement]) string {
			return "\x1B[7mTest pager\x1B[27m"
		},
	}
	// Run pager event loop
	err := pg.EventLoop(
		&tty8pe.Tty{},                  // terminal input
		lines,                          // data source
		colorable.NewColorableStdout()) // output

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
