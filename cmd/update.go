package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
  Use:   "update [deployment_id] [status]",
  Short: "Update deployment",
  Long:  `Use this to update the deployment with a new status.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Placeholder: Updated deployment with id: ")
  },
}
