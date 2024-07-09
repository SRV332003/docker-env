package main

import (
	"fmt"

	"github.com/SRV332003/envdaemon/handlers"
	"github.com/spf13/cobra"
)

func main() {
	// get cli arguments

	var rootCmd = &cobra.Command{Use: "envdaemon"}

	// take env file path as input with default value as .env
	var envFilePath string
	rootCmd.PersistentFlags().StringVarP(&envFilePath, "env", "e", ".env", "env file path")

	// take docker file path as input with default value as Dockerfile
	var dockerFilePath string
	rootCmd.PersistentFlags().StringVarP(&dockerFilePath, "docker", "d", "Dockerfile", "docker file path")

	// flage to create .env.example file
	var example bool
	rootCmd.PersistentFlags().BoolVarP(&example, "example", "x", false, "create .env.example file")

	// if --val is passed, set val to true
	var val bool
	rootCmd.PersistentFlags().BoolVarP(&val, "val", "v", false, "update env values")

	// help command
	rootCmd.AddCommand(&cobra.Command{Use: "help", Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: envdaemon [flags]")
		fmt.Println("Flags:")
		fmt.Println("  -e, --env string: env file path (default \".env\")")
		fmt.Println("  -d, --docker string: docker file path (default \"Dockerfile\")")
		fmt.Println("  -v, --val: update env values")
		fmt.Println("  -x, --example: create .env.example file")
	}})

	rootCmd.Execute()
	if val {
		fmt.Println("Updating Env values")
		err := handlers.UpdateEnvVals(envFilePath, dockerFilePath)
		if err != nil {
			panic(err)
		}
	} else if !example {
		fmt.Println("Updating Env keys")
		err := handlers.UpdateEnvKeys(envFilePath, dockerFilePath)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Creating .env.example file")
		err := handlers.CreateEnvExample(envFilePath)
		if err != nil {
			panic(err)
		}
	}
}
