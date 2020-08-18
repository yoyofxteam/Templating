package main

import (
	"fmt"
	"os"
)

func main() {
	tempDir := os.Args[0]
	workDir, _ := os.Getwd()
	fmt.Println(tempDir)
	fmt.Println(workDir)
}
