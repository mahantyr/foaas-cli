package cmd

import (
  "fmt"
  "github.com/palash25/foaas-cli/fuck"
  "github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of FOaaS",
  Run: func(cmd *cobra.Command, args []string) {
    version := fuck.GetVersion()
    fmt.Println(version)
  },
}
