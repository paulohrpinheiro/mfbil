package main

import (
	"fmt"
	"os"

	bf "github.com/paulohrpinheiro/mfbil/bf"
)

func main() {
	argsCount := len(os.Args)
	if argsCount < 2 {
		panic("Please, provide a BF program as argument, and optionally, a string of input data")
	}

	inputData := ""
	if argsCount == 3 {
		inputData = os.Args[2]
	}

	vm := bf.New(os.Args[1], 3000, inputData)
	err := vm.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(vm.Output))
}
