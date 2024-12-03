package main

import (
	"fmt"
	"os"

	"github.com/brandond/bundle2jwks/app"
)

func main() {
	if err := app.New().Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
