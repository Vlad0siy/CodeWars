.PHONY: build
build:
	go build -v -work ./main.go

.PHONY: run
run:
		go run BitCounting.go
