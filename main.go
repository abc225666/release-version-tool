package main

import (
	"fmt"
	"os"
)

func main() {
	v, err := GetNewVersion()
	if err != nil {
		fmt.Println("fail to up to new version", err)
		os.Exit(-1)
	}
	fmt.Printf("%s", v)
}
