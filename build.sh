# /bin/bash

export GOOS=windows
export GOARCH=386
go build -o bin/wings-of-liberty.exe main.go


export GOOS=linux
export GOARCH=amd64
go build -o bin/wings-of-liberty main.go