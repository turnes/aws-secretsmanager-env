package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "aws-secretsmanager-env",
	Short: "Aws-secretsmanager-env is a tool to manage and to parse AWS Secrets Manager.",
	Long: `It's a tool to manage and to parse AWS Secrets Manager.
You can list, retrieve, delete and upload `,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

}

var Secret string
var Verbose bool
var Source string
var Region string

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
