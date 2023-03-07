package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"

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

		var clientId string
		var tenantId string
		var scope string

		// Look up values from environment variables
		clientIdEnv, clientIdEnvSet := os.LookupEnv("GET_JWT_AZURE_AD_CLIENT_ID")
		tenantIdEnv, tenantIdEnvSet := os.LookupEnv("GET_JWT_AZURE_AD_TENANT_ID")
		scopeEnv, scopeEnvSet := os.LookupEnv("GET_JWT_AZURE_AD_SCOPE")
		copyEnv, _ := os.LookupEnv("GET_JWT_COPY")
		copyEnvBool, _ := strconv.ParseBool(copyEnv)

		// Read values from the flags
		clientIdFlag, _ := cmd.Flags().GetString("client-id")
		tenantIdFlag, _ := cmd.Flags().GetString("tenant-id")
		scopeFlag, _ := cmd.Flags().GetString("scope")
		copyFlagBool, _ := cmd.Flags().GetBool("copy")

		// Set the variables, giving precendence to the env var value if it's set
		if clientIdEnvSet {
			clientId = clientIdEnv
		} else if cmd.Flag("client-id").Changed {
			clientId = clientIdFlag
		} else {
			log.Fatal("Please set either the `--client-id` flag or the environment variable GET_JWT_AZURE_AD_CLIENT_ID.")
		}

		if tenantIdEnvSet {
			tenantId = tenantIdEnv
		} else if cmd.Flag("tenant-id").Changed {
			tenantId = tenantIdFlag
		} else {
			log.Fatal("Please set either the `--tenant-id` flag or the environment variable GET_JWT_AZURE_AD_TENANT_ID.")
		}

		if scopeEnvSet {
			scope = scopeEnv
		} else if cmd.Flag("scope").Changed {
			scope = scopeFlag
		} else {
			log.Fatal("Please set either the `--scope` flag or the environment variable GET_JWT_AZURE_AD_SCOPE.")
		}

		accessToken := getAzureJwt(ctx, clientId, tenantId, scope)

		if copyEnvBool || copyFlagBool {
			log.Info("Writing JWT to clipboard 📋")
			clipboard.WriteAll(accessToken)
		} else {
			log.Info("Printing JWT below 📜")
			fmt.Println(accessToken)
		}
	},
}

func init() {
	rootCmd.AddCommand(azureCmd)

	azureCmd.Flags().String("client-id", "", "Help message for client-id")
	azureCmd.Flags().String("tenant-id", "", "Help message for tenant-id")
	azureCmd.Flags().String("scope", "", "Help message for scope")
	azureCmd.Flags().Bool("copy", false, "Copy to clipboard")
}

func getAzureJwt(ctx context.Context, clientId string, tenantId string, scope string) string {
	log.Debug("Client ID: %s", clientId)
	log.Debug("Tenant ID: %s", tenantId)
	log.Debug("Scope: %s", scope)

	authority := "https://login.microsoftonline.com/" + tenantId

	scopes := []string{scope}

	// Initialize a public client
	publicClientApp, err := public.New(clientId, public.WithAuthority(authority))

	// TODO: Handle some known errors
	//
	// Error:
	//   AADSTS9002327: Tokens issued for the 'Single-Page Application' client-type may only be redeemed via cross-origin requests.
	// Fix:
	//   Authentication > enable Web > redirect url of http://localhost
	//
	// Error:
	//   AADSTS7000218: The request body must contain the following parameter: 'client_assertion' or 'client_secret'.
	// Fix:
	//   Authentication > Advanced settings > Allow public client flows := true
	if err != nil {
		log.Fatal("Failed to log in successfully", "error", err)
	}

	log.Info("Opening browser to login...")

	// Open browser to do the interactive login
	result, err := publicClientApp.AcquireTokenInteractive(context.Background(), scopes)
	if err != nil {
		log.Fatal("Failed to log in successfully", "error", err)
	}

	return result.AccessToken
}
