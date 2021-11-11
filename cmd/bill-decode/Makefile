.PHONY: clean
all: clean
	GOOS=windows GOARCH=amd64 go build -o bill-decode.exe main.go
	GOOS=linux GOARCH=amd64 go build -o bill-decode-linux main.go
	GOOS=darwin GOARCH=amd64 go build -o bill-decode-darwin main.go

win: clean
	GOOS=windows GOARCH=amd64 go build -o bill-decode.exe main.go

clean:
	rm -f bill-decode*