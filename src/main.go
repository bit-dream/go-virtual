package main

import (
	"fmt"
	"github.com/bit-dream/go-virtual/pkg/candatabase"
	"go.einride.tech/can"
)

func main() {
	const inputFile = "/Users/headquarters/Documents/Code/go-virtual/src/tesla_can.dbc"
	data, err := candatabase.LoadDbc(inputFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(data.Messages))

	frame := can.Frame{
		ID:         1160,
		Length:     4,
		Data:       can.Data{0xFF, 0x10, 0x10, 0x10},
		IsRemote:   false,
		IsExtended: false,
	}
	candatabase.DecodeFrame(frame, data.Messages[0])
}
