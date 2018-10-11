package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	os.Args = append(os.Args[:0], os.Args[len(os.Args)-1])
	filename := os.Args[0]

	fmt.Println(filename)
	if filename == "" {
		fmt.Println("An error occured with the input.")
		os.Exit(1)
	}

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("Content of", filename, ":", string(content))
}
