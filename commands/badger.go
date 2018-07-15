package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var databaseDir string

// BadgerCommand holds the primary root command for the utility
var BadgerCommand = &cobra.Command{
	Use: "Badger",
}

// Execute is the entrypoint for command execution
func Execute() {
	if err := BadgerCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
