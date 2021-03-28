package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	putCmd.PersistentFlags().StringVarP(&Region, "region", "r", "us-west-1", "AWS region Default: us-west-1")
	rootCmd.AddCommand(putCmd)
}

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Upload a secret to AWS Secrets Manager",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}
