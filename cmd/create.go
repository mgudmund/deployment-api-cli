package cmd

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new deployment",
	Long:  `Use this to create a new deployment. Retuns the deployment id`,
	Run: func(cmd *cobra.Command, args []string) {
		location := createDeployment()
		fmt.Println("Placeholder: created deployment with id: " + location)
	},
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
	/*d := Deployment{
		name:           "mgtest1",
		version:        "v0.0.1",
		statusPage:     "none",
		repositoryUrl:  "none",
		commitSha:      "none",
		environment:    "Prod",
		status:         "created",
		productName:    "mgtest1-prod",
		capabilityName: "techarch",
	}*/
	/*if err != nil {
		fmt.Println("Error creating json object")
		os.Exit(-1)
	}*/

	client := resty.New()
	//client.SetDebug(true)
	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthScheme("Bearer").
		SetAuthToken(token).
		SetBody(Deployment{
			Name:           "mgtest1",
			Version:        "v0.0.1",
			StatusPage:     "none",
			RepositoryUrl:  "none",
			CommitSha:      "none",
			Environment:    "production",
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
		fmt.Println(resp.StatusCode())
	}

	return resp.Header().Get("Location")
}
