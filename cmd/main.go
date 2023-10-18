package main

import (
	"os"

	"github.com/taylormonacelli/greenpeace"
)

func main() {
	code := greenpeace.Main()
	os.Exit(code)
}
