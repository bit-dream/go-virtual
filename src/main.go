package main

import (
	"fmt"
	"github.com/bit-dream/go-virtual/pkg/virtualizer"
	"time"
)

func main() {
	fmt.Println("Hello, world!")

	stop := make(chan bool)

	go virtualizer.PeriodicTimer(2000, stop, func() {
		fmt.Println("This is a timer!")
	})

	time.Sleep(10 * time.Second)
	stop <- true
	/*
		const networkFile = "/Users/headquarters/Documents/Code/go-virtual/src/tesla.network"

		virtualModel, err := model.LoadModel(networkFile)
		if err != nil {
			fmt.Println(err)
		}

		err = model.UpdateVirtualModelByDefinitions(virtualModel)
		if err != nil {
			fmt.Println(err)
		}

		var a = 2
		fmt.Println(a)
	*/
}
