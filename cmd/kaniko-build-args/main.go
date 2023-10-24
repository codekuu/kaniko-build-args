package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var jsonString string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "kaniko-build-args",
	Short: "Takes a json key, value object and returns a string that can be used in kaniko build",
	Long:  `Takes a json key, value object and returns a string that can be used in kaniko build`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.BindPFlag("json", cmd.Flags().Lookup("json")); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		buildParse()
	},
}

func main() {
	// Takes a json string
	RootCmd.PersistentFlags().StringVarP(&jsonString, "json", "j", "", "JSON string to parse")
	RootCmd.MarkPersistentFlagRequired("json")

	// Execute the command
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func buildParse() {
	var buildArgs string
	// Parse and loop through the json string, should add to the buildArgs as "--build-arg "key=value"
	for k, v := range viper.GetStringMapString("json") {
		buildArgs += fmt.Sprintf("--build-arg %s=%s ", k, v)
	}

	fmt.Println(buildArgs)
}
