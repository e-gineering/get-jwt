package cmd

import (
	"context"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	"github.com/atotto/clipboard"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// azureCmd represents the azure command
var azureCmd = &cobra.Command{
	Use:   "azure",
	Short: "Get a JWT from Azure AD",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Start a context
		ctx := context.Background()

		// Read values from the flags
		clientId, _ := cmd.Flags().GetString("client-id")
		tenantId, _ := cmd.Flags().GetString("tenant-id")
		scope, _ := cmd.Flags().GetString("scope")

		getAzureJwt(ctx, clientId, tenantId, scope)
	},
}

func init() {
	rootCmd.AddCommand(azureCmd)

	azureCmd.Flags().String("client-id", "", "Help message for client-id")
	azureCmd.Flags().String("tenant-id", "", "Help message for tenant-id")
	azureCmd.Flags().String("scope", "", "Help message for scope")
}

func getAzureJwt(ctx context.Context, clientId string, tenantId string, scope string) {
	log.Debug("Client ID: %s", clientId)
	log.Debug("Tenant ID: %s", tenantId)
	log.Debug("Scope: %s", scope)

	authority := "https://login.microsoftonline.com/" + tenantId

	scopes := []string{scope}

	// Initialize a public client
	publicClientApp, err := public.New(clientId, public.WithAuthority(authority))
	if err != nil {
		log.Error("Failed to initialize the public client", "error", err)
	}

	log.Info("Opening browser to login...")

	// Open browser to do the interactive login
	result, err := publicClientApp.AcquireTokenInteractive(context.Background(), scopes)
	if err != nil {
		log.Error("Failed to log in successfully", "error", err)
	}

	accessToken := result.AccessToken

	log.Info("Writing to clipboard 📋")
	clipboard.WriteAll(accessToken)
}
