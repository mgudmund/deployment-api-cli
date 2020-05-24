package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var Environment string
var CommitSha string
var Version string

type YamlSpec struct {
	Product struct {
		Name           string `yaml:"name"`
		StatusPage     string `yaml:"statusPage"`
		RepositoryUrl  string `yaml:"repositoryUrl"`
		ProductName    string `yaml:"productName"`
		CapabilityName string `yaml:"capabilityName"`
	}
}

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

	//Get env vars
	url := os.Getenv("DEPLOYMENT_API_URL")
	token := os.Getenv("DEPLOYMENT_API_TOKEN")
	productSpec := getProductSpec()
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
		SetBody(Deployment{
			Name:           productSpec.Product.Name,
			Version:        Version,
			StatusPage:     productSpec.Product.StatusPage,
			RepositoryUrl:  productSpec.Product.RepositoryUrl,
			CommitSha:      CommitSha,
			Environment:    Environment,
			Status:         "created",
			ProductName:    productSpec.Product.ProductName,
			CapabilityName: productSpec.Product.CapabilityName,
		}).
		Post(url + "/deployments/")

	if err != nil {
		fmt.Println("Call to API failed" + err.Error())
		os.Exit(-1)
	}
	if resp.IsError() {
		fmt.Println("HTTP Return Code: " + strconv.Itoa(resp.StatusCode()))
		fmt.Println(resp.String())
		os.Exit(1)
	}
	return resp.Header().Get("Location")
}

func getProductSpec() (productSpec YamlSpec) {

	specFile := "product-spec.yaml"
	// Check for product-spec.yaml
	if _, err := os.Stat(specFile); err != nil {
		fmt.Printf("The file product-spec.yaml does not exist, but expected by cli.\n")
		os.Exit(1)
	}
	// Get values from product-spec.yaml
	yamlBuf, err := ioutil.ReadFile(specFile)
	if err != nil {
		fmt.Println("Error reading "+specFile+": #%v", err)
	}
	var yamlSpec YamlSpec
	err = yaml.Unmarshal(yamlBuf, &yamlSpec)
	if err != nil {
		fmt.Println("Error unmarshalling:" + err.Error())
	}
	return yamlSpec

}
