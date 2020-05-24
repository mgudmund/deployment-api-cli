package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
  Use:   "search",
  Short: "Search for deployments",
  Long:  `Use this to search for deployments. Returns search results`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Not implemented")
  },
}

