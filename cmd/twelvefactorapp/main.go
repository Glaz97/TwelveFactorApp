package main

import (
	"fmt"
	"os"

	"github.com/Glaz97/twelvefactorapp/internal/command"
)

func main() {
	if err := command.GetRootCmd().Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
