package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Verbose bool

var rootCmd = &cobra.Command{
	Use:   "deployment-api-cli",
	Short: "deployment-api-cli is a cli for the deloyment API",
	Long: `
To track our deployments and improve our DevOps practices we provide an API with a cli. 
The cli can be included in deployment pipelines to track metrics.

The cli expects two environment variables to exist, 
- DEPLOYMENT_API_TOKEN that contains the JWT you genereted when singing up 
- DEPLOYMENT_API_URL that contains the URL to the API

`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("main")
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "V", false, "verbose output")

}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
