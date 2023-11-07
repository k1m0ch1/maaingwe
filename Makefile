# Get version from git hash
git_hash := $(shell git describe --tags)

# Get current date
current_time = $(shell date +"%Y-%m-%d:T%H:%M:%S")

# Add linker flags
linker_flags = '-s -w -X main.buildTime=${current_time} -X main.version=${git_hash}'

build:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=386 go build -ldflags=${linker_flags} -gcflags=all="-l -B" -o bin/maaingwe-linux-386 main.go
	upx --best --ultra-brute bin/maaingwe-linux-386
	GOOS=windows GOARCH=386 go build -ldflags=${linker_flags} -gcflags=all="-l -B" -o bin/maaingwe-windows-386.exe main.go
	upx --best --ultra-brute bin/maaingwe-windows-386.exe
	GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -gcflags=all="-l -B" -o bin/maaingwe-linux-amd64 main.go
	upx --best --ultra-brute bin/maaingwe-linux-amd64
	GOOS=windows GOARCH=amd64 go build -ldflags=${linker_flags} -gcflags=all="-l -B" -o bin/maaingwe-windows-amd64.exe main.go
	upx --best --ultra-brute bin/maaingwe-windows-amd64.exe
	GOOS=darwin GOARCH=arm64 go build -ldflags=${linker_flags} -gcflags=all="-l -B" -o bin/maaingwe-darwin-arm64 main.go
	

