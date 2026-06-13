package nemo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"

	"github.com/nyaosorg/go-ttyadapter/fav"

	"github.com/hymkor/go-generics-list"

	"github.com/hymkor/nemo/asyncpager"
	"github.com/hymkor/nemo/internal/ansi"
	"github.com/hymkor/nemo/pager"
)

type textElement string

func (t textElement) Display(w int) string {
	s := string(t)
	for {
		i := strings.IndexByte(s, '\t')
		if i < 0 {
			return s + "\x1B[0m"
		}
		s = s[:i] + "    "[i%4:] + s[i+1:]
	}
}

var rxSymbol = regexp.MustCompile("[\x00-\x08\x0A-\x1A\x1C-\x1F]")

func (app *Application) main1(source io.Reader, title string) error {
	lines := list.New[textElement]()

	pg := &asyncpager.Pager[textElement]{
		Status: func(session *pager.Session[textElement]) string {
			var b strings.Builder
			if title != "" {
				b.WriteString(ansi.Reverse)
				b.WriteString(title)
				b.WriteString(ansi.Inverse)
			}
			L := lines.Len()
			start := session.WinPos
			end := session.WinPos + session.Pager.ContentHeight - 1
			if end+1 > L {
				end = L - 1
			}
			fmt.Fprintf(&b, " %d-%d / %d", start+1, end+1, L)
			return b.String()
		},
	}

	br := bufio.NewReader(source)

	getter := func() (textElement, error) {
		text, err := br.ReadString('\n')
		if err == nil || err == io.EOF {
			if app.ShowControl {
				var buffer strings.Builder
				for _, c := range text {
					if c < 0x20 {
						fmt.Fprintf(&buffer, "\\x%02X", c)
					} else {
						buffer.WriteRune(c)
					}
				}
				text = buffer.String()
			} else if app.StripCr {
				text = strings.ReplaceAll(text, "\n", "")
				text = strings.ReplaceAll(text, "\r", "")
			} else {
				text = rxSymbol.ReplaceAllStringFunc(text, func(s string) string {
					return string(rune(0x2400 + int(s[0])))
				})
			}
			return textElement(text), err
		}
		return "", err
	}

	store := func(obj textElement, err error) bool {
		if err != nil {
			return false
		}
		//if obj != nil {
		lines.PushBack(obj)
		//}
		return true
	}
	c := colorable.EnableColorsStdout(nil)
	defer c()

	return pg.EventLoop(
		new(fav.Tty),
		getter,
		store,
		lines,
		colorable.NewColorableStdout())
}

type Application struct {
	ShowControl bool `flag:"show-control,display control characters as \\xNN"`
	StripCr     bool `flag:"strip-cr,remove carriage returns"`
}

func (app *Application) Run(args []string) error {
	if len(args) < 1 {
		if isatty.IsTerminal(os.Stdin.Fd()) {
			return fmt.Errorf("Nemo %s-%s-%s", version, runtime.GOOS, runtime.GOARCH)
		}
		return app.main1(os.Stdin, "<STDIN>")
	}
	fd, err := os.Open(args[0])
	if err != nil {
		return err
	}
	defer fd.Close()
	return app.main1(fd, args[0])
}
