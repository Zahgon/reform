package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/reform.v1"
)

var (
	execFlags  = flag.NewFlagSet("exec", flag.ExitOnError)
	execSplitF = execFlags.Bool("split", false, "Split statements by semicolon; does not handles them in string literals")
)

func init() {
	execFlags.Usage = func() {
		fmt.Fprintf(os.Stderr, "`exec` command executes SQL queries from given files or stdin.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  %s [global flags] exec [exec flags] [file names]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Global flags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExec flags:\n")
		execFlags.PrintDefaults()

		// TODO mention -split flag
		fmt.Fprintf(os.Stderr, `
Each file's content is executed as a single query. If it contains multiple
statements, make sure SQL driver supports them. If file names are not given,
a query is read from stdin until EOF, then executed.
`)
	}
}

// readFiles reads queries from given files, or from stdin, if files are not given
func readFiles(files []string, split bool) (queries []string) {
	_ = "STUB: not implemented"
	// read stdin
	return nil
}

// read files

//nolint:gosec

// cmdExec implements exec command.
func cmdExec(db *reform.DB, files []string) { _ = "STUB: not implemented"; return }
