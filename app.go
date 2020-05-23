package main

import (
    "fmt"
    "os"
    "github.com/go-resty/resty/v2"
)

func main() {
    deployment_url := os.Getenv("DEPLOYMENT_URL")
    token := os.Getenv("TOKEN")
    action := os.Getenv("ACTION")

    // Create a Resty Client
    client := resty.New()
    resp, err := client.R().
      SetHeader("Content-Type", "application/json").
      SetBody(`{"username":"testuser", "password":"testpass"}`).
      //SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
      Post("https://myapp.com/login") 

    fmt.Println(resp)
    fmt.Println(err)   
    fmt.Println(deployment_url)
    fmt.Println(token)
    fmt.Println(action)
}
