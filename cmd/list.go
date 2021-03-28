package cmd

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/spf13/cobra"
)

func init() {
	listCmd.PersistentFlags().StringVarP(&Region, "region", "r", "us-west-1", "AWS region Default: us-west-1")
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all secrets of a region.",
	Long:  "List all secrets of a regions.",
	Run: func(cmd *cobra.Command, args []string) {
		listSecrets()
	},
}

func listSecrets() ([]string, error) {
	//Create a Secrets Manager client
	session := session.Must(session.NewSession())
	svc := secretsmanager.New(session,
		aws.NewConfig().WithRegion(Region))

	input := &secretsmanager.ListSecretsInput{}

	result, err := svc.ListSecrets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidNextTokenException:
				fmt.Println(secretsmanager.ErrCodeInvalidNextTokenException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil, err
	}

	if result.SecretList != nil {
		list := result.SecretList
		for _, s := range list {
			fmt.Println(*s.Name)
		}
		return nil, nil
	} else {
		return nil, err
	}

}
