// Package reform implements reform command.
package main

import (
	"flag"
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/internal"
)

var (
	logger *internal.Logger

	debugF   = flag.Bool("debug", false, "Enable debug logging")
	gofmtF   = flag.Bool("gofmt", true, "Format with gofmt")
	versionF = flag.Bool("version", false, "Print version and exit")
)

func processFile(path, file, pack string) error { _ = "STUB: not implemented"; return nil }

//nolint // real Close() error checked at the end if we reach it

// decide about view/table suffix

func gofmt(path string) { _ = "STUB: not implemented"; return }

func goformat(filePath string) { _ = "STUB: not implemented"; return }

//nolint:gosec

//nolint:gosec

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "reform - a better ORM generator. %s.\n\n", reform.Version)
		fmt.Fprintf(os.Stderr, "Usage:\n\n")
		fmt.Fprintf(os.Stderr, "  %s [flags] [packages or directories]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  go generate [flags] [packages or files] (with '//go:generate reform' in files)\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *versionF {
		fmt.Println(reform.Version)
		os.Exit(0)
	}

	logger = internal.NewLogger("reform: ", *debugF)

	logger.Debugf("Environment:")
	for _, pair := range os.Environ() {
		if strings.HasPrefix(pair, "GO") {
			logger.Debugf("\t%s", pair)
		}
	}

	wd, err := os.Getwd()
	if err != nil {
		logger.Fatalf("%s", err)
	}
	logger.Debugf("wd: %s", wd)
	logger.Debugf("args: %v", flag.Args())

	// process arguments
	for _, arg := range flag.Args() {
		// import arg as directory or package path
		var pack *build.Package
		s, err := os.Stat(arg)
		if err == nil && s.IsDir() {
			pack, err = build.ImportDir(arg, 0)
		}
		if os.IsNotExist(err) {
			err = nil
		}
		if pack == nil && err == nil {
			pack, err = build.Import(arg, wd, 0)
		}
		if err != nil {
			logger.Fatalf("%s: %s", arg, err)
		}

		logger.Debugf("%#v", pack)

		var changed bool
		for _, f := range pack.GoFiles {
			err = processFile(pack.Dir, f, pack.Name)
			if err != nil {
				logger.Fatalf("%s %s: %s", arg, f, err)
			}
			goformat(filepath.Join(pack.Dir, f))
			changed = true
		}

		if changed {
			gofmt(pack.Dir)
		}
	}

	// process go generate environment
	file := os.Getenv("GOFILE")
	pack := os.Getenv("GOPACKAGE")
	if file != "" && pack != "" {
		err := processFile(wd, file, pack)
		if err != nil {
			logger.Fatalf("%s", err)
		}
		goformat(file)
		gofmt(wd)
	}
}
