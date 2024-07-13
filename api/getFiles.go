package api

import (
	"log"
	"os"
)

func GetFiles(r string) []string {
	var fileNames []string

	files, err := os.ReadDir(r)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames
}
