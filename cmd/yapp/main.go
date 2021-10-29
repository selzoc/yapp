package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/reidmit/yapp/internal/config"
	"github.com/reidmit/yapp/internal/server"
	"github.com/spf13/cobra"
)

const appName = "yapp"

// expected to be supplied at build time:
var appVersion string
var appCommit string

func main() {
	rootCmd := &cobra.Command{
		Use:     appName + " <path/to/yapp.yml>",
		Version: getAppVersion(),
		Short:   "Run your app",
		Long:    "Run your app using the given yapp.yml file",
		Args:    cobra.ExactArgs(1),
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd:   true,
			DisableDescriptions: true,
			DisableNoDescFlag:   true,
		},
		Run: func(cmd *cobra.Command, args []string) {
			configPath := args[0]
			port, _ := cmd.Flags().GetInt64("port")

			appConfig, err := config.Load(configPath)
			if err != nil {
				fmt.Printf("error reading config: %v", err)
				os.Exit(1)
			}

			server.Serve(appConfig, port)
		},
	}

	rootCmd.Flags().Int64(
		"port",
		getDefaultPort(),
		"port to listen on (can also set $PORT env var)",
	)

	rootCmd.Execute()
}

func getAppVersion() string {
	var v string

	if appVersion != "" {
		v = appVersion
	} else {
		v = "0.0.0"
	}

	if appCommit != "" {
		v += "-" + appCommit
	} else {
		v += "-dev"
	}

	return v
}

func getDefaultPort() int64 {
	if portVar, isSet := os.LookupEnv("PORT"); isSet {
		port, _ := strconv.ParseInt(portVar, 10, 64)
		return port
	}

	return 7000
}
