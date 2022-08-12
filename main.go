package main

import "fmt"

func main() {
	code := []byte{1, 5, 1, 7, 10, 0}
	ret, _ := Run(code)
	fmt.Println(ret)
}
