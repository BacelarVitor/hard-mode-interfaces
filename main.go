package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fp, err := getFilePath(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(fp)
}

func getFilePath(args []string) (string, error) {
	if args != nil && len(args) > 0 {
		return args[1], nil
	}

	return "", errors.New("A file path most be provided")
}