package cmd

import (
  "fmt"
  "github.com/palash25/foaas-cli/fuck"
  "github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Hugo",
  Long:  `All software has versions. This is Hugo's`,
  Run: func(cmd *cobra.Command, args []string) {
    version := fuck.GetVersion()
    fmt.Println(version)
  },
}
