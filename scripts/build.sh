CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server src/main.go
chmod 777 server
