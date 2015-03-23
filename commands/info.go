package commands

import (
	"flag"
	"fmt"
	"github.com/wayt/odrone-client/query"
	"os"
)

func infoUsage() {

	fmt.Printf("Usage: odrone-client search [keyword]\n")
	os.Exit(1)
}

func Info() {

	flag.Parse()

	if flag.NArg() < 2 {

		infoUsage()
	}

	pkg := packageInfo{}
	if err := query.Get("/package/"+flag.Arg(1), &pkg); err != nil {

		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Println("Package:\t" + pkg.Name)
	fmt.Println("\tauthor:", pkg.AuthorEmail)
	fmt.Println("\trevision:", pkg.Revision)
	if len(pkg.Summary) != 0 {
		fmt.Println("\tsummary:", pkg.Summary)
	}
	if len(pkg.Dependencies) != 0 {
		fmt.Println("\tdependencies:", pkg.Dependencies)
	}
	fmt.Println("\tfile:", pkg.FileUrl)
}
