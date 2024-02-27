package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Реализовать функцию получения всех файлов из директории
func getFiles(dir string) ([]string, error) {
	data, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	allFiles := make([]string, 0)
	for _, entry := range data {
		if !entry.IsDir() {
			allFiles = append(allFiles, filepath.Join(dir, entry.Name()))
		} else {
			newFiles, err := getFiles(filepath.Join(dir, entry.Name()))
			if err != nil {
				return nil, err
			}

			allFiles = append(allFiles, newFiles...)
		}
	}

	return allFiles, nil
}

// Реализовать функцию фильтрации переданных названий файлов
func filterFiles(files []string, filter string) ([]string, error) {
	if filter == "" {
		return files, nil
	}

	resultFiles := make([]string, 0)
	if filter[0:2] == "*." {
		for _, file := range files {
			if filepath.Ext(file) == filter[1:] {
				resultFiles = append(resultFiles, file)
			}
		}
	} else {
		for _, file := range files {
			if strings.Contains(file, filter) {
				resultFiles = append(resultFiles, file)
			}
		}
	}

	return resultFiles, nil
}

// Реализовать функцию сохранения в файл
func saveToFile(fileNames []string, outFile string) error {
	file, err := os.Create(outFile)
	if err != nil {
		return err
	}

	defer file.Close()

	for _, fileName := range fileNames {
		_, err = file.WriteString(fileName + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	files, err := getFiles("./")
	if err != nil {
		log.Fatal(err)
	}

	filteredFiles, err := filterFiles(files, "*.png")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range filteredFiles {
		fmt.Println(file)
	}

	err = saveToFile(filteredFiles, "aboba.txt")
	if err != nil {
		log.Fatal(err)
	}
}
