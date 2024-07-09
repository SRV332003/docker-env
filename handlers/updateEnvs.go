package handlers

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func UpdateEnvKeys(envfile string, Dockerfile string) error {
	// open envfile in read mode
	file, err := os.OpenFile(envfile, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}

	// read the content of envfile
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

			line := "ENV " + keyValue[0] + " {" + keyValue[0] + "}"

			envKeys = append(envKeys, line)
		}
	}

	// open Dockerfile in read & write mode
	dockerfile, err := os.OpenFile(Dockerfile, os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	// search for {{.ENV}} in Dockerfile
	fileinfo, err = dockerfile.Stat()
	content = make([]byte, fileinfo.Size())
	_, err = dockerfile.Read(content)
	if err != nil {
		return err
	}

	// prepare keys to replace {{.ENV}} in Dockerfile
	replaceText := strings.Join(envKeys, "\n")
	replaceText = replaceText + "\n"
	replaceText = "\n\n# {{ENV}}\n" + replaceText + "# {{END ENV}}\n\n"

	// search for {{.ENV}} in Dockerfile

	re := regexp.MustCompile(`\n*#\s*{{ENV}}(\D|\d)*#\s*{{END ENV}}\n*`)
	match := string(re.Find(content))

	if match == "" {
		log.Println("{{.ENV}} not used previously in Dockerfile")
		re = regexp.MustCompile(`\n*#\s*{{ENV}}`)
		match = string(re.Find(content))
		fmt.Println(match)
		if match == "" {
			return fmt.Errorf("{{.ENV}} not found in Dockerfile")
		}
	}

	// replace {{.ENV}} with key value pairs from envfile
	newContent := strings.Replace(string(content), match, replaceText, 1)

	// clear the contents of the Dockerfile
	err = dockerfile.Truncate(0)
	if err != nil {
		return err
	}

	// write the new content to Dockerfile
	_, err = dockerfile.WriteAt([]byte(newContent), 0)
	if err != nil {
		return err
	}

	return nil

}

func UpdateEnvVals(envfile string, Dockerfile string) error {
	// open envfile in read mode
	file, err := os.OpenFile(envfile, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}

	// read the content of envfile
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

			line := "ENV " + keyValue[0] + "=" + keyValue[1]

			envKeys = append(envKeys, line)
		}
	}

	// open Dockerfile in read & write mode
	dockerfile, err := os.OpenFile(Dockerfile, os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	// search for {{.ENV}} in Dockerfile
	fileinfo, err = dockerfile.Stat()
	content = make([]byte, fileinfo.Size())
	_, err = dockerfile.Read(content)
	if err != nil {
		return err
	}

	// prepare keys to replace {{.ENV}} in Dockerfile
	replaceText := strings.Join(envKeys, "\n")
	replaceText = replaceText + "\n"
	replaceText = "\n\n# {{ENV}}\n" + replaceText + "# {{END ENV}}\n\n"

	// search for {{.ENV}} in Dockerfile

	re := regexp.MustCompile(`\n*#\s*{{ENV}}(\D|\d)*#\s*{{END ENV}}\n*`)
	match := string(re.Find(content))

	if match == "" {
		log.Println("{{.ENV}} not used previously in Dockerfile")
		re = regexp.MustCompile(`\n*#\s*{{ENV}}`)
		match = string(re.Find(content))
		fmt.Println(match)
		if match == "" {
			return fmt.Errorf("{{.ENV}} not found in Dockerfile")
		}
	}

	// replace {{.ENV}} with key value pairs from envfile
	newContent := strings.Replace(string(content), match, replaceText, 1)

	// clear the contents of the Dockerfile
	err = dockerfile.Truncate(0)
	if err != nil {
		return err
	}

	// write the new content to Dockerfile
	_, err = dockerfile.WriteAt([]byte(newContent), 0)
	if err != nil {
		return err
	}

	return nil

}
