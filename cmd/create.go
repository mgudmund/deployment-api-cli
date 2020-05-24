package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var Environment string
var CommitSha string
var Version string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new deployment",
	Long:  `Use this to create a new deployment. Retuns the deployment id`,
	Run: func(cmd *cobra.Command, args []string) {
		location := createDeployment()
		if location == "" {
			fmt.Println("No deployment created.")
		} else {
			fmt.Println("Deployment ID: " + location)
		}
	},
}

func init() {

	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&Environment, "environment", "e", "", "Deployment Environment")
	createCmd.Flags().StringVarP(&CommitSha, "commitsha", "c", "", "Git Commit SHA")
	createCmd.Flags().StringVarP(&Version, "version", "v", "", "Product version as semver")
	createCmd.MarkFlagRequired("environment")
	createCmd.MarkFlagRequired("commitsha")
	createCmd.MarkFlagRequired("version")

}

func createDeployment() (loc string) {

	//Get env vars
	url := os.Getenv("DEPLOYMENT_API_URL")
	token := os.Getenv("DEPLOYMENT_API_TOKEN")

	type Deployment struct {
		Name           string `json:"name"`
		Version        string `json:"version"`
		StatusPage     string `json:"statusPage"`
		RepositoryUrl  string `json:"repositoryUrl"`
		CommitSha      string `json:"commitSha"`
		Environment    string `json:"environment"`
		Status         string `json:"status"`
		ProductName    string `json:"productName"`
		CapabilityName string `json:"capabilityName"`
	}

	client := resty.New()

	client.SetDebug(true)

	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthScheme("Bearer").
		SetAuthToken(token).
		SetBody(Deployment{
			Name:           "mgtest1",
			Version:        Version,
			StatusPage:     "none",
			RepositoryUrl:  "none",
			CommitSha:      CommitSha,
			Environment:    Environment,
			Status:         "created",
			ProductName:    "mgtest1-prod",
			CapabilityName: "SP",
		}).
		Post(url + "/deployments/")

	if err != nil {
		fmt.Println("Call to API failed" + err.Error())
		os.Exit(-1)
	}
	if resp.IsError() {
		fmt.Println("HTTP Return Code: " + strconv.Itoa(resp.StatusCode()))
		fmt.Println(resp.String())
		os.Exit(-1)
	}
	return resp.Header().Get("Location")
}
