/*
cli is a command line client for interacting with a skycoin cx-tracker
*/
package main

import (
	"fmt"
	"os"

	"github.com/sn-srdjan/cx-tracker-cli/src/cli"
)

func main() {
	cfg, err := cli.LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cxTrackerCLI, err := cli.NewCLI(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := cxTrackerCLI.Execute(); err != nil {
		os.Exit(1)
	}
}
