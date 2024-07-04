package main

import (
	"fmt"
	"os"

	"github.com/SRV332003/docker-env/handlers"
)

func main() {
	// get cli arguments
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Usage: docker-env <envfile> <Dockerfile>")
		fmt.Println("Default values are .env and Dockerfile")

		os.Exit(1)
	}
	if len(args) == 1 {
		args = append(args, ".env", "Dockerfile")
	}

	if len(args) == 2 {
		args = append(args, "Dockerfile")
	}

	fmt.Println("Creating Dockerfile template")
	err := handlers.UpdateEnvs(args[1], args[2])
	if err != nil {
		fmt.Println(err)
	}

}
