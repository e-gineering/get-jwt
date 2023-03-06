package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// azureCmd represents the azure command
var azureCmd = &cobra.Command{
	Use:   "azure",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("azure called")

		clientId, _ := cmd.Flags().GetString("client-id")
		tenantId, _ := cmd.Flags().GetString("tenant-id")
		scope, _ := cmd.Flags().GetString("scope")

		getAzureJwt(clientId, tenantId, scope)

	},
}

func init() {
	rootCmd.AddCommand(azureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// azureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	azureCmd.Flags().String("client-id", "", "Help message for client-id")
	azureCmd.Flags().String("tenant-id", "", "Help message for tenant-id")
	azureCmd.Flags().String("scope", "", "Help message for scope")
}

func getAzureJwt(clientId string, tenantId string, scope string) {

	fmt.Println("hi")
	fmt.Println(clientId)
	fmt.Println(tenantId)
	fmt.Println(scope)

	authority := "https://login.microsoftonline.com/" + tenantId

	scopes := []string{scope}

	// Initialize a public client
	publicClientApp, err := public.New(clientId, public.WithAuthority(authority))
	if err != nil {
		log.Fatal(err)
	}

	// Open browser to do the interactive login
	result, err := publicClientApp.AcquireTokenInteractive(context.Background(), scopes)
	if err != nil {
		log.Fatal(err)
	}

	accessToken := result.AccessToken
	log.Default().Println("Writing to clipboard 📋")
	clipboard.WriteAll(accessToken)
}
