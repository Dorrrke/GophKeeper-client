build: download
	go mod tidy
	go env -w CGO_ENABLED=1
	go env -w GOOS=linux
	go build -ldflags "-X cmd.buildVersion=v1.0.1 -X cmd.buildCommit=8d00bcc -X 'cmd.buildDate=$(date +'%Y/%m/%d %H:%M:%S')'" -o gophkeeper ./cmd/gophkeeper/main.go

build-win: download
	go mod tidy
	go env -w CGO_ENABLED=1
	go build -ldflags "-X cmd.buildVersion=v1.0.1 -X cmd.buildCommit=8d00bcc -X 'cmd.buildDate=$(date +'%Y/%m/%d %H:%M:%S')'" -o gophkeeper.exe ./cmd/gophkeeper/main.go

download:
	go mod download
	go mod verify

run: download
	go run ./cmd/gophkeeper/main.go -help

