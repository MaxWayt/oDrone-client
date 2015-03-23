package commands

import (
	"flag"
	"fmt"
	"github.com/wayt/odrone-client/query"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func installUsage() {

	fmt.Printf("Usage: odrone-client install [package]\n")
	os.Exit(1)
}

func Install() {

	flag.Parse()

	if flag.NArg() < 2 {

		installUsage()
	}

	pkg := packageInfo{}
	if err := query.Get("/package/"+flag.Arg(1), &pkg); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Println("Downloading", pkg.Name, " ...")

	filename := fmt.Sprintf("/tmp/%s", filepath.Base(pkg.FileUrl))

	out, err := os.Create(filename)
	resp, err := http.Get(pkg.FileUrl)

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	resp.Body.Close()
	out.Close()

	fmt.Println("Waiting for Android device...")
	if err := exec.Command("adb", "wait-for-device").Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Println("Installing ...")
	if err := exec.Command("adb", "install", filename).Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
