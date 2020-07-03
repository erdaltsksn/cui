package main

import (
	"fmt"
	"os"

	"github.com/erdaltsksn/cui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cui.VersionCmd)
}
