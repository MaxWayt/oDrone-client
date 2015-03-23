package commands

import (
	"flag"
	"fmt"
	"github.com/wayt/odrone-client/query"
	"os"
)

func searchUsage() {

	fmt.Printf("Usage: odrone-client search [keyword]\n")
	os.Exit(1)
}

func Search() {

	flag.Parse()

	if flag.NArg() < 2 {

		searchUsage()
	}

	result := []packageInfo{}
	if err := query.Get("/packages/"+flag.Arg(1), &result); err != nil {

		fmt.Println("error:", err)
		os.Exit(1)
	}

	if len(result) == 0 {
		fmt.Println("No match :(")
		return
	}

	fmt.Println("Result:")
	for _, pkg := range result {

		fmt.Println(pkg.Name)
		if len(pkg.Summary) != 0 {
			fmt.Println("\t" + pkg.Summary)
		}
	}

	fmt.Println("\nTotal:", len(result), "package(s) found(s).")
}
