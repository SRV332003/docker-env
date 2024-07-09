package handlers

import (
	"fmt"
	"os"
	"strings"
)

// CreateEnvExample creates .env.example file
func CreateEnvExample(envFilePath string) error {
	fmt.Println("Creating .env.example file from ", envFilePath)
	// read the content of envfile
	file, err := os.OpenFile(envFilePath, os.O_RDONLY, 0644)
	fileinfo, err := file.Stat()
	content := make([]byte, fileinfo.Size())
	_, err = file.Read(content)
	if err != nil {
		return err
	}

	// parse key value pairs from envfile
	envs := strings.Split(string(content), "\n")
	envKeys := make([]string, 0)
	for _, env := range envs {
		if env != "" && env != "\n" && env[0] != '#' {
			keyValue := strings.Split(env, "=")
			if len(keyValue) != 2 {
				continue
			}

			line := keyValue[0]

			envKeys = append(envKeys, line)
		}
	}
	// create .env.example file
	exampleFile, err := os.Create(".env.example")
	if err != nil {
		return err
	}
	defer exampleFile.Close()

	fmt.Println("Env keys: ", envKeys)
	// write env keys to .env.example file
	for _, key := range envKeys {
		_, err := exampleFile.WriteString(fmt.Sprintf("%s=\n", key))
		if err != nil {
			return err
		}
	}

	return nil
}
