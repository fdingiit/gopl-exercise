package main

import (
	"os"
	"fmt"
)

func main() {
	for _, a := range os.Args[1:] {
		fmt.Println(a)
	}
}
