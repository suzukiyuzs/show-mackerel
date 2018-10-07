package main

import (
	"os"

	showmackerel "github.com/suzukiyuzs/show-mackerel/lib"
)

func main() {
	os.Exit(showmackerel.Show(os.Args))
}
