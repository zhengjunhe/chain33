go env -w CGO_ENABLED=0
go build -o build/dplatformos.exe github.com/33cn/dplatformos/cmd/dplatformos
go build -o build/dplatformos-cli.exe github.com/33cn/dplatformos/cmd/cli
