package main

import (
	"fmt"

	"github.com/KevOub/ggbe/pkg/gmmu"

	gcpu "github.com/KevOub/ggbe/pkg/gcpu"
)

func main() {
	fmt.Println("Test")
	gcpu.InitCPU()
	gmmu.InitMMU()

	gcpu.DoInstruction(0x11)
	gcpu.DoInstruction(0x21)

	/*
		for {

			gcpu.DoInstruction()
		} */
}
