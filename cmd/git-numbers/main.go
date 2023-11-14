package main

import (
	"os"

	gitnumbers "github.com/dyson/git-numbers/internal/git-numbers"
)

func main() {
	os.Exit(gitnumbers.Run(os.Args))
}
