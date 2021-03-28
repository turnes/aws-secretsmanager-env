package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	deleteCmd.PersistentFlags().StringVarP(&Region, "region", "r", "us-west-1", "AWS region Default: us-west-1")
	rootCmd.AddCommand(deleteCmd)

}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a secret.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete command")
	},
}
