package main

import (
	"fmt"
	"github.com/wayt/odrone-client/commands"
	"os"
)

func help() {

	fmt.Printf("Usage: %s command [args]\n", os.Args[0])
}

func main() {

	if len(os.Args) < 2 {

		help()
		return
	}

	switch os.Args[1] {

	case "search":
		commands.Search()
	case "info":
		commands.Info()
	case "package":
		commands.Package()
	case "install":
		commands.Install()
	}
}
