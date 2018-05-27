package cmd

import (
  "fmt"
  "github.com/palash25/foaas-cli/fuck"
  "github.com/spf13/cobra"
)

var operationsCmd = &cobra.Command{
  Use:   "operations",
  Short: "Print the operations of FOaaS",
  Run: func(cmd *cobra.Command, args []string) {
    ops := fuck.GetOperations()
    fmt.Println(ops)
  },
}
