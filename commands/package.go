package commands

import (
	"flag"
	"fmt"
	"github.com/maxwayt/odrone-client/query"
	"os"
)

var (
	createPkg = flag.Bool("c", false, "Create a new package")
	updatePkg = flag.Bool("u", false, "Update a package")
	deletePkg = flag.Bool("d", false, "Delete package")

	pkgName = flag.String("n", "", "Package name")

	pkgAuthor = flag.String("a", "", "Package author email")
	pkgDeps   = flag.String("deps", "", "Package dependencies")
	pkgSum    = flag.String("s", "", "Package summary")
)

func Package() {

	flag.CommandLine.Parse(os.Args[2:])

	if *createPkg {

		params := map[string]string{
			"name":         *pkgName,
			"author_email": *pkgAuthor,
		}
		if len(*pkgDeps) != 0 {
			params["dependencies"] = *pkgDeps
		}
		if len(*pkgSum) != 0 {
			params["summary"] = *pkgSum
		}

		f, err := os.Open(flag.Arg(0))
		if err != nil {

			fmt.Println("error:", err)
			os.Exit(1)
		}

		if err := query.Post("/package", params, f, nil); err != nil {

			fmt.Println("error:", err)
			os.Exit(1)
		}

	} else if *updatePkg {

		params := map[string]string{
			"name": *pkgName,
		}
		if len(*pkgDeps) != 0 {
			params["dependencies"] = *pkgDeps
		}
		if len(*pkgSum) != 0 {
			params["summary"] = *pkgSum
		}

		f, err := os.Open(flag.Arg(0))
		if err != nil {

			fmt.Println("error:", err)
			os.Exit(1)
		}

		if err := query.Put("/package", params, f, nil); err != nil {

			fmt.Println("error:", err)
			os.Exit(1)
		}

	} else if *deletePkg {

		if err := query.Delete("/package/"+*pkgName, nil); err != nil {

			fmt.Println("error:", err)
			os.Exit(1)
		}

	} else {

		fmt.Fprintf(os.Stderr, "Usage: %s package [OPTIONS] [FILE]\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
}
