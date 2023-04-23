package main

import (
	"fmt"
	"github.com/bit-dream/go-virtual/src/candatabase"
)

func main() {
	const inputFile = "/Users/headquarters/Documents/Code/go-virtual/src/tesla_can.dbc"
	data, err := candatabase.LoadDbc(inputFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
}
