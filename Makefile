build-arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o modbusTestTool-arm main.go

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o modbusTestTool main.go

