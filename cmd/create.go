package cmd

import (
  "fmt"

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
    fmt.Println("Placeholder: created deployment with id: ")
  },
}
