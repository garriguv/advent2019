advent2019:
	go build -o $@

test:
	go test -v ./...

.PHONY: advent2019 test
