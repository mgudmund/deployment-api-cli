package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
  Use:   "delete",
  Short: "Delete a deployment",
  Long:  `Use this to delete a deployment. Retuns the deployment id that was deleted`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Not implemented")
  },
}

