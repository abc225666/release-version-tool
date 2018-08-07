package main

import (
	"fmt"
)

func main() {
	res, err := GetLatestTag()
	fmt.Println(res)
	fmt.Println(err)
}
