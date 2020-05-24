package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var User string
var Password string
var Url string

func init() {
	rootCmd.AddCommand(signupCmd)
	signupCmd.Flags().StringVarP(&User, "user", "u", "", "Username")
	signupCmd.Flags().StringVarP(&Password, "password", "p", "", "Password")
	signupCmd.Flags().StringVarP(&Url, "url", "U", "", "URL to the API")

	signupCmd.MarkFlagRequired("user")
	signupCmd.MarkFlagRequired("password")
	signupCmd.MarkFlagRequired("url")

}

var signupCmd = &cobra.Command{
	Use:   "sign-up -u [user] -p [password]",
	Short: "Sign-up for deployment tracking",
	Long:  `UUse this to signup first time. Use your network id as you username and a password of your choosing.`,
	Run: func(cmd *cobra.Command, args []string) {
		pwd := signUp()
		token := getToken(pwd)
		fmt.Println("")
		fmt.Println("------------------------------------------------------------------------------------")
		fmt.Println("")
		fmt.Println("Token: " + token)
		fmt.Println("")
		fmt.Println("- Treat this token as a secret, if used in pipelines")
		fmt.Println("- The cli expects this token in an environment variable called DEPLOYMENT_API_TOKEN")
		fmt.Println("")
		fmt.Println("------------------------------------------------------------------------------------")
		fmt.Println("")

	},
}

func signUp() (pwd string) {

	type signUpRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	client := resty.New()
	if Verbose {
		client.SetDebug(true)
	}
	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(signUpRequest{
			Username: User,
			Password: Password,
		}).
		Post(Url + "/auth/sign-up")

	if err != nil {
		fmt.Println("Call to API failed" + err.Error())
		os.Exit(1)
	}
	if resp.IsError() {
		fmt.Println("HTTP Return Code: " + strconv.Itoa(resp.StatusCode()))
		fmt.Println(resp.String())
		os.Exit(1)
	}
	var signupreq signUpRequest
	json.Unmarshal([]byte(resp.String()), &signupreq)
	return signupreq.Password

}

func getToken(pwd string) (token string) {
	type tokenRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type Token struct {
		Token string `json:"token"`
	}
	client := resty.New()
	if Verbose {
		client.SetDebug(true)
	}
	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(tokenRequest{
			Username: User,
			Password: pwd,
		}).
		Post(Url + "/auth/generate-token")

	if err != nil {
		fmt.Println("Call to API failed" + err.Error())
		os.Exit(1)
	}
	if resp.IsError() {
		fmt.Println("HTTP Return Code: " + strconv.Itoa(resp.StatusCode()))
		fmt.Println(resp.String())
		os.Exit(1)
	}
	var tokenreq Token
	json.Unmarshal([]byte(resp.String()), &tokenreq)
	return tokenreq.Token

}
