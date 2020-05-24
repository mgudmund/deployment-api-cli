package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
  Use:   "list",
  Short: "List deployments",
  Long:  `Use this to list deployments. Returns the a list of deployments`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Not implemented")
  },
}

