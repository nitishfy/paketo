/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "paketo",
	Short: "Paketo: facilitates interaction with the OpenBuildServices platform to manage Kubernetes packages for new releases",
	Long:  "Paketo: facilitates interaction with the OpenBuildServices platform to manage Kubernetes packages for new releases",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Kubernetes OBS ClI Tool")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(Reconcile())
}
