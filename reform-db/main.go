// Package reform-db implements reform-db command.
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/internal"
)

var (
	logger *internal.Logger

	debugF   = flag.Bool("debug", false, "Enable debug logging")
	driverF  = flag.String("db-driver", "", "Database driver (required)")
	sourceF  = flag.String("db-source", "", "Database connection string (required)")
	waitF    = flag.Duration("db-wait", 0, "Wait for database connection to be established, retrying every second")
	versionF = flag.Bool("version", false, "Print version and exit")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "reform-db - a better ORM tool. %s.\n\n", reform.Version)
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  %s [global flags] [command] [command flags] [arguments]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Global flags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nCommands (run reform-db [command] -h for more information):\n")
		fmt.Fprintf(os.Stderr, "  exec  - executes SQL queries from given files or stdin\n")
		fmt.Fprintf(os.Stderr, "  query - executes SQL queries from given files or stdin, and returns results\n")
		fmt.Fprintf(os.Stderr, "  init  - generates Go model files for existing database schema\n\n")
		fmt.Fprintf(os.Stderr, "Registered database drivers: %s.\n", strings.Join(sql.Drivers(), ", "))
	}
}

func getDB() *reform.DB { _ = "STUB: not implemented"; return nil }

// Use single connection so various session-related variables work.
// For example: "PRAGMA foreign_keys" for SQLite3, "SET IDENTITY_INSERT" for MS SQL, etc.

func main() {
	flag.Parse()

	if *versionF {
		fmt.Println(reform.Version)
		os.Exit(0)
	}

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	*driverF = strings.TrimSpace(*driverF)
	*sourceF = strings.TrimSpace(*sourceF)

	logger = internal.NewLogger("reform-db: ", *debugF)

	// flagsets used below are created with ExitOnError, so Parse() is not expected to return errors
	switch flag.Arg(0) {
	case "exec":
		if err := execFlags.Parse(flag.Args()[1:]); err != nil {
			panic(err)
		}
		cmdExec(getDB(), execFlags.Args())

	case "query":
		if err := queryFlags.Parse(flag.Args()[1:]); err != nil {
			panic(err)
		}
		cmdQuery(getDB(), queryFlags.Args())

	case "init":
		if err := initFlags.Parse(flag.Args()[1:]); err != nil {
			panic(err)
		}

		if initFlags.NArg() > 1 {
			logger.Fatalf("Expected zero or one argument for %q, got %d", "init", initFlags.NArg())
		}

		dir := initFlags.Arg(0)
		var err error
		if dir == "" {
			if dir, err = os.Getwd(); err != nil {
				logger.Fatalf("%s", err)
			}
		}
		if dir, err = filepath.Abs(dir); err != nil {
			logger.Fatalf("%s", err)
		}
		fi, err := os.Stat(dir)
		if os.IsNotExist(err) {
			logger.Fatalf("%q should be existing directory", dir)
		}
		if err != nil {
			logger.Fatalf("%s", err)
		}
		if !fi.IsDir() {
			logger.Fatalf("%q should be existing directory", dir)
		}

		cmdInit(getDB(), dir)

	default:
		flag.Usage()
		logger.Fatalf("Unexpected command %q", flag.Arg(0))
	}
}
