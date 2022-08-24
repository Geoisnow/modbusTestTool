build-arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o modbusTestTool-arm main.go

build-amd:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o modbusTestTool-amd main.go

build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o modbusTestTool-win main.go


