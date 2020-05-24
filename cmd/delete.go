package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var DeleteId string

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&DeleteId, "deleteid", "d", "", "Deployment ID to be deleted")
	deleteCmd.MarkFlagRequired("deleteid")
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a deployment",
	Long:  `Use this to delete a deployment. Retuns the deployment id that was deleted`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteDeployment()
		fmt.Println("Deployment deleted: " + DeleteId)
	},
}

func deleteDeployment() {
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
		Delete(url + "/deployments/" + DeleteId)

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
