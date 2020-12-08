go env -w CGO_ENABLED=0
go build -o build/dplatform.exe github.com/33cn/dplatform/cmd/dplatform
go build -o build/dplatform-cli.exe github.com/33cn/dplatform/cmd/cli
