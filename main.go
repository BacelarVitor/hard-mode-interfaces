package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	fp, err := getFilePath(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	file, err := os.Open(fp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	
}

func getFilePath(args []string) (string, error) {
	if args != nil && len(args) > 0 {
		return args[1], nil
	}

	return "", errors.New("A file path most be provided")
}

func printContent(file *os.File) {
	io.Copy(os.Stdout, file)
	fmt.Println(file)
}