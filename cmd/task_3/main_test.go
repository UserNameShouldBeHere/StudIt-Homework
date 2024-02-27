package main

import (
	"bufio"
	"os"
	"reflect"
	"strings"
	"testing"
)

var testData []string = []string{
	"test.go",
	"test.png",
	"main.py",
	"main_test.go",
	"main.go",
	"readme.md",
	"icon.png",
	"index.html",
	"index.js",
	"login.html",
	"api.go",
	"index.py",
}

func TestGetFiles(t *testing.T) {
	files, err := getFiles("../../cmd")
	if err != nil {
		t.Error(err)
	}

	for i := range files {
		files[i] = strings.ReplaceAll(files[i], "\\", "/")
	}

	result := []string{
		"../../cmd/task_1/main.go",
		"../../cmd/task_1/main_test.go",
		"../../cmd/task_2/main.go",
		"../../cmd/task_2/main_test.go",
		"../../cmd/task_3/main.go",
		"../../cmd/task_3/main_test.go",
	}

	if !reflect.DeepEqual(files, result) {
		t.Error("Incorrect result")
	}
}

func TestWithExtFilter(t *testing.T) {
	cases := []struct {
		testName string
		filter   string
		expected []string
	}{
		{
			"filter by name",
			"test",
			[]string{
				"test.go",
				"test.png",
				"main_test.go",
			},
		},
		{
			"filter by name",
			"index",
			[]string{
				"index.html",
				"index.js",
				"index.py",
			},
		},
		{
			"filter by name",
			"result",
			[]string{},
		},
		{
			"filter by extention",
			"*.go",
			[]string{
				"test.go",
				"main_test.go",
				"main.go",
				"api.go",
			},
		},
		{
			"filter by extention",
			"*.html",
			[]string{
				"index.html",
				"login.html",
			},
		},
		{
			"filter by extention",
			"*.css",
			[]string{},
		},
	}

	for _, currentCase := range cases {
		t.Run(currentCase.testName, func(t *testing.T) {
			files, err := filterFiles(testData, currentCase.filter)
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(files, currentCase.expected) {
				t.Errorf("Incorrect result: expected %v, got %v", currentCase.expected, files)
			}
		})
	}
}

func TestSaveToFile(t *testing.T) {
	err := saveToFile(testData, "test.txt")
	if err != nil {
		t.Error(err)
	}

	file, err := os.Open("test.txt")
	if err != nil {
		t.Error(err)
	}

	recievedData := make([]string, 0)
	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		recievedData = append(recievedData, line)
	}

	file.Close()

	if !reflect.DeepEqual(recievedData, testData) {
		t.Errorf("Incorrect result: expected %v, got %v", testData, recievedData)
	}

	err = os.Remove("./test.txt")
	if err != nil {
		t.Error(err)
	}
}
