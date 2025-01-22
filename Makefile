
gimb:
	go mod tidy
	go build

release:
	go mod tidy
	go build
	mv ./gimb ~/gimb
