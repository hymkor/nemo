nemo
====

A terminal pager for Go, usable both as a CLI and as a library.

Overview
--------

`nemo` is a lightweight pager designed for terminal applications.

It can be used as:

- A CLI pager (`cmd/nemo`)
- A reusable library for building TUI applications

The library is split into three layers:

- `pager` — core pager working on a linked list of elements
- `asyncpager` — background-loading extension of `pager`
- `nemo` — CLI-oriented wrapper with a simple entry point

Features
--------

- Works with both files and standard input
- Handles input text containing ANSI escape sequences (e.g. colored output) correctly
- Horizontal scrolling for long lines (no automatic wrapping)
- Designed for embedding into Go applications
- Optional background loading for large inputs

Installation
------------

### CLI

```sh
go install github.com/hymkor/nemo/cmd/nemo@latest
```

Usage (CLI)
-----------

```sh
nemo file.txt
cat file.txt | nemo
```

Library Usage
--------------

### Basic Idea

To use the pager, implement a type with a `Display(width int) string` method
and provide a list of elements.

Example
-------

```example.go
package main

import (
    "fmt"
    "os"

    "github.com/mattn/go-colorable"

    "github.com/nyaosorg/go-ttyadapter/fav"

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
        new(fav.Tty),                   // terminal input
        lines,                          // data source
        colorable.NewColorableStdout()) // output

    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }
}
```

Async Pager
-----------

Use `asyncpager` when your data source is loaded incrementally or in the background.

```go
import "github.com/hymkor/nemo/asyncpager"
```

It provides the same interface as `pager`, with additional support for background loading.

Package Structure
-----------------

```
nemo/
├── cmd/nemo        # CLI entry point
├── pager           # core pager
├── asyncpager      # async extension
└── (root)          # Run() entry
```

Design Notes
------------

- `pager` is minimal and synchronous
- `asyncpager` adds concurrency without changing the core interface
- `Display(width)` gives full control over rendering
- `nemo.Run()` provides a simple integration point for applications

License
-------

MIT License
