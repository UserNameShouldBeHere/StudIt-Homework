lint:
	golangci-lint run

run-1:
	go run cmd/task_1/main.go

test-1:
	go test -v cmd/task_1/main_test.go cmd/task_1/main.go

run-2:
	go run cmd/task_2/main.go

test-2:
	go test -v cmd/task_2/main_test.go cmd/task_2/main.go

run-3:
	go run cmd/task_3/main.go

test-3:
	go test -v cmd/task_3/main_test.go cmd/task_3/main.go