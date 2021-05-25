.PHONY: build
build:
	go build -v -work ./project/cmd/main.go

.PHONY: run
run:
		go run BitCounting.go
