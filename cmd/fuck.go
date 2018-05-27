package cmd

import (
  "fmt"
  "github.com/palash25/foaas-cli/fucks"
  "github.com/spf13/cobra"
)

func init() {
  fuckCmd.Flags().StringVarP(&from, "from", "f", "", "Who the fuck is this fuck from?")
}

//  name, lang, company, reference, reaction, thing, behaviour, tool, do, something
var from string

var fuckCmd = &cobra.Command{
  Use:   "fuck",
  Short: "Request various fucks from FOaaS",
  Run: func(cmd *cobra.Command, args []string) {
    ops := fucks.Asshole(from)
    fmt.Println(ops)
  },
}
