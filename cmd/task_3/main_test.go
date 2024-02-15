package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestWithoutFilter(t *testing.T) {
	files, err := FilterFiles("../../cmd", "")
	if err != nil {
		t.Error(err)
	}

	for i, _ := range files {
		files[i] = strings.ReplaceAll(files[i], "\\", "/")
	}

	result := []string{
		"../../cmd/task_1/main.go",
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
	files, err := FilterFiles("../../", "*.go")
	if err != nil {
		t.Error(err)
	}

	for i, _ := range files {
		files[i] = strings.ReplaceAll(files[i], "\\", "/")
	}

	result := []string{
		"../../cmd/task_1/main.go",
		"../../cmd/task_2/main.go",
		"../../cmd/task_2/main_test.go",
		"../../cmd/task_3/main.go",
		"../../cmd/task_3/main_test.go",
		"../../internal/api/api.go",
	}

	if !reflect.DeepEqual(files, result) {
		t.Error("Incorrect result")
	}
}

func TestWithNameFilter(t *testing.T) {
	files, err := FilterFiles("../../", "main")
	if err != nil {
		t.Error(err)
	}

	for i, _ := range files {
		files[i] = strings.ReplaceAll(files[i], "\\", "/")
	}

	result := []string{
		"../../cmd/task_1/main.go",
		"../../cmd/task_2/main.go",
		"../../cmd/task_2/main_test.go",
		"../../cmd/task_3/main.go",
		"../../cmd/task_3/main_test.go",
	}

	if !reflect.DeepEqual(files, result) {
		t.Error("Incorrect result")
	}
}
