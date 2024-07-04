package handlers

import "os"

func CreateTemplate(dockerfile string) error {
	// open dockerfile in read mode
	file, err := os.OpenFile(dockerfile, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// create a new file in write mode
	newFile, err := os.Create("Docker.template")
	if err != nil {
		return err
	}
	defer newFile.Close()

	// copy the content of dockerfile to newFile
	_, err = file.WriteTo(newFile)
	if err != nil {
		return err
	}

	return nil

}
