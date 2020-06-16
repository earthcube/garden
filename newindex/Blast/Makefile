BINARY := blast
VERSION :=`cat VERSION`
.DEFAULT_GOAL := blast
ORG := geocodes

blast:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 env go build -o $(BINARY)
