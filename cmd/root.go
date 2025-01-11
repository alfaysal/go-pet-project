/*
Copyright © 2025 Md All Faysal
*/
package cmd

import (
	"fmt"
	"github.com/alfaysal/go-pet-project/internal/config"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-pet-project",
	Short: "This is just a demo project to familiar with GO",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	err := config.InitialiseConfig()
	if err != nil {
		fmt.Println(err)
	}
}
