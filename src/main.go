package main

import (
	"fmt"
	"github.com/bit-dream/go-virtual/pkg/can_database"
)

func main() {
	fmt.Println("Hello, world!")

	ecu1, err := can_database.LoadEcu("/Users/headquarters/Documents/Code/go-virtual/pkg/ecu_model/ecu_model_test/test1.ecu")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ecu1)

}
