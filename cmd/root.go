package root

import (
  "fmt"
  "os"
  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "foaas",
  Short: "FOaaS CLI",
  Long: `This is a terminal client to interact with the FOaaS
                (Fuck Off as a Service) API built with love by palash25
                in an effort to learn Go`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("fasfaf")
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println("aaa", err)
    os.Exit(1)
  }
}
