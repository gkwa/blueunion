package main

import (
	"os"

	"github.com/taylormonacelli/blueunion"
)

func main() {
	code := blueunion.Execute()
	os.Exit(code)
}
