package main

import (
	"os"

	"github.com/lrettig/smeshbeat/cmd"

	_ "github.com/lrettig/smeshbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
