package main

import (
	"fmt"
	"os"

	"github.com/thaibui2308/anime-cli/cli"
)

func main() {
	cli.AddCommands()

	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
