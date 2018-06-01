package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/palash25/foaas-cli/fucks"
	"github.com/spf13/cobra"
)

type operation struct {
	Name   string            `json:"name"`
	URL    string            `json:"url"`
	Fields []operationFields `json:"fields"`
}

type operationFields struct {
	Name  string `json:"name"`
	Field string `json:"field"`
}

var operationsCmd = &cobra.Command{
	Use:   "operations",
	Short: "Print the operations of FOaaS",
	Run: func(cmd *cobra.Command, args []string) {
		ops := fucks.GetOperations()
		var opsList []operation
		if err := json.Unmarshal([]byte(ops), &opsList); err != nil {
			fmt.Println("failed to parse list of fucks")
			os.Exit(1)
		}

		var out string
		for i := range opsList {
			var fields []string
			for b := range opsList[i].Fields {
				fields = append(fields, opsList[i].Fields[b].Field)
			}
			out = fmt.Sprintf("%s - %s", strings.ToLower(opsList[i].Name), strings.Join(fields, ", "))
			fmt.Println(out)
		}
	},
}
