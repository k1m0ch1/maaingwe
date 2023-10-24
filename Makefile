build:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/maaingwe-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/maaingwe-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/maaingwe-windows-386.exe main.go
	GOOS=freebsd GOARCH=amd64 go build -o bin/maaingwe-freebsd-amd64 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/maaingwe-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/maaingwe-windows-amd64.exe main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/maaingwe-darwin-arm64 main.go