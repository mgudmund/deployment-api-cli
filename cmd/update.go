package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var DeploymentId string
var Status string

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&DeploymentId, "deploymentid", "d", "", "Deployment ID")
	updateCmd.Flags().StringVarP(&Status, "status", "s", "", "Deployment status")

	updateCmd.MarkFlagRequired("deploymentid")
	updateCmd.MarkFlagRequired("status")

}

var updateCmd = &cobra.Command{
	Use:   "update [deployment_id] [status]",
	Short: "Update deployment",
	Long:  `Use this to update the deployment with a new status.`,
	Run: func(cmd *cobra.Command, args []string) {
		updateDeployment()
		fmt.Println("Updated deployment with id: " + DeploymentId)

	},
}

func updateDeployment() {
	//Get env vars
	url := os.Getenv("DEPLOYMENT_API_URL")
	token := os.Getenv("DEPLOYMENT_API_TOKEN")

	client := resty.New()
	if Verbose {
		client.SetDebug(true)
	}
	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthScheme("Bearer").
		SetAuthToken(token).
		Patch(url + "/deployments/" + DeploymentId + "?status=" + Status)

	if err != nil {
		fmt.Println("Call to API failed" + err.Error())
		os.Exit(1)
	}
	if resp.IsError() {
		fmt.Println("HTTP Return Code: " + strconv.Itoa(resp.StatusCode()))
		fmt.Println(resp.String())
		os.Exit(1)
	}

}
