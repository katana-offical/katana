build:
	go build -o target/katana main.go
 
run:
	go run main.go
 
compile: compileCore compilePlugin

compilePlugin:
	#fuck！

compileCore:
	GOOS=darwin GOARCH=amd64 go build -o build/katana main.go

runTest:


clean:
	echo "clean"

help:
	echo "help"

all:
	echo "all"