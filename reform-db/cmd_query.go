package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/reform.v1"
)

var (
	queryFlags  = flag.NewFlagSet("query", flag.ExitOnError)
	querySplitF = queryFlags.Bool("split", false, "Split statements by semicolon; does not handles them in string literals")
)

func init() {
	queryFlags.Usage = func() {
		fmt.Fprintf(os.Stderr, "`query` command executes SQL queries from given files or stdin, and returns results.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  %s [global flags] query [query flags] [file names]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Global flags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nQuery flags:\n")
		queryFlags.PrintDefaults()

		// TODO mention -split flag
		fmt.Fprintf(os.Stderr, `
Each file's content is executed as a single query. If it contains multiple
statements, make sure SQL driver supports them. If file names are not given,
a query is read from stdin until EOF, then executed.
`)
	}
}

// cmdQuery implements query command.
func cmdQuery(db *reform.DB, files []string) { _ = "STUB: not implemented"; return }

// write table header

// read all rows, scan each field to []byte
