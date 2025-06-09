package main

import (
	"fmt"
	"http_cli/flags"
	"os"
)

func main() {
	fl, err := flags.GetFlags()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(fl.List)
	fmt.Println(fl.Method)
	fmt.Println(fl.Filepath)
	fmt.Println(fl.BinName)
	fmt.Println(fl.ID)
}
