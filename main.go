package main

import (
	"fmt"

	"github.com/KevOub/ggbe/pkg/gdebug"
	"github.com/KevOub/ggbe/pkg/gmmu"

	gcpu "github.com/KevOub/ggbe/pkg/gcpu"
)

func main() {
	fmt.Println("Test")
	gcpu.InitCPU(true)
	gmmu.InitMMU()

	gcpu.DoInstruction(0x11)
	gcpu.DoInstruction(0x21)
	fmt.Print(gdebug.WhatIsThisCode(0x21, false))

	/*
		for {

			gcpu.DoInstruction()
		} */
}
