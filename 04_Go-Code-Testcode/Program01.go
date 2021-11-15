package main

//import libraries test02
import (
	"fmt"
)

//main function code
func main() {
	//test the code with go run Program01.go
	fmt.Println("Hello, Black Hat Gophers")
}

//go build Program01.go
//-> create an executable ./Program01

// reduce the binary size: go build -ldflags"-w -s"
//without debugging and symbol table

//GOOS Operating System
//GOARCH Architecture
//GOOS="linux" GOARCH="amd64" go build Program01.go

//go doc fmt.Println -> get information
