name: Default

test:
	cd calculator
	go test -v ./...

buid:
	cd cmd/calculator
	go build -o calculator