package gcpu

import (
	"fmt"
	"log"

	"github.com/KevOub/ggbe/pkg/gdebug"
	"github.com/KevOub/ggbe/pkg/gmmu"
)

// DEBUG decides whether to enable debugging
var DEBUG bool

/*
REGISTERS
	The first four regs are general purpose 16 bits
	Access the high and low for the individual aspects
	Accomplished that by inline h() and f() for the SplitRegister

*/

// SplitRegister split from 16 bits to 8 bits
type SplitRegister struct {
	h uint8
	l uint8
}

func (reg SplitRegister) f() uint16 {
	return uint16(((reg.h << 8) | reg.l))
}

//Set does exactly what it sounds like
func (reg *SplitRegister) Set(num uint16) {
	reg.h = uint8(num >> 8)
	reg.l = uint8(num & 255)
}

//Increment increses lower bit. If overflow, set high bit
func (reg *SplitRegister) Increment() {
	tmp := reg.l + 1
	if tmp < CPUReg.PC.l {
		reg.h++
	} else {
		reg.l++
	}

}

//PrettyPrint outputs the register, pretty
func (reg SplitRegister) PrettyPrint(name string) {
	fmt.Printf("%s\t", name)
	value := fmt.Sprintf("%02x", reg.h)
	fmt.Printf("%s\t", value)

	value = fmt.Sprintf("%02x", reg.l)
	fmt.Printf("%s\t", value)

	value = fmt.Sprintf("%04x", reg.f())
	fmt.Printf("%s\t", value)

	fmt.Printf("%d\t", reg.f())
	fmt.Printf("%08b", reg.f())
	fmt.Println()
}

/*
//SetHigh takes the value and updates f()
func (reg *SplitRegister) SetHigh(num uint16) {
	reg.f = (reg.f & 255) // Clear higher bits
	reg.f = (num << 8) | reg.f
	// TODO broken...
}

//SetLow takes the value and updates f()
func (reg *SplitRegister) SetLow(num uint16) {
	reg.f = (reg.f & 65280) // Clear lower bits
	reg.f = (num) | reg.f
	// TODO broken...
}

func (reg SplitRegister) h() uint16 {
	return (reg.f & 65280)
}

//Get gets the value from the register
func (reg *SplitRegister) Get() uint16 {
	return reg.f
}
*/

// Registers are the cpu registers
type Registers struct {
	AF SplitRegister
	BC SplitRegister
	DE SplitRegister
	HL SplitRegister
	SP SplitRegister
	PC SplitRegister
}

//Show displays the registers nicely
func (reg Registers) Show() {
	// HEADER
	fmt.Println("ADDR\tHIGH\tLOW\tTOTAL\tTOTAL\tBIN")
	// AF
	CPUReg.AF.PrettyPrint("AF")

	// BC
	CPUReg.BC.PrettyPrint("BC")

	// DE
	CPUReg.DE.PrettyPrint("DE")

	// HL
	CPUReg.HL.PrettyPrint("HL")

	// SP
	CPUReg.SP.PrettyPrint("SP")

	// PC
	CPUReg.PC.PrettyPrint("PC")

}

// CPUReg are the live registers
var CPUReg = Registers{SplitRegister{0, 0}, SplitRegister{0, 0}, SplitRegister{0, 0}, SplitRegister{0, 0}, SplitRegister{0, 0}, SplitRegister{0, 0}}

// InitCPU creates cpu
func InitCPU(debug bool) {
	DEBUG = debug
	fmt.Println("CPU Started")
	if DEBUG {
		fmt.Println("\t debugging active")
		fmt.Println("")
	}

	CPUReg.PC.Set(0)
	CPUReg.SP.Set(0xFFFE)
	gmmu.InitMMU()
}

// DoInstruction does instruction
func DoInstruction() {
	opcode := gmmu.Memory[CPUReg.PC.f()]
	value := fmt.Sprintf("%02x", opcode)
	fmt.Printf("OPCODE:\t0x%s\n", value)
	CPUReg.Show()
	fmt.Println("\n+++")
	switch opcode {
	// NOOP
	case 0x00:
		fmt.Print()

	// LD SP,PC
	case 0x31:
		// TODO check setting low and high bits
		// Load SP into the program counter

		CPUReg.SP.h = uint8(gmmu.Memory[CPUReg.PC.f()])
		CPUReg.PC.Increment()
		CPUReg.SP.l = uint8(gmmu.Memory[CPUReg.PC.f()])

	// HALT
	case 0x76:
		CPUReg.PC.Increment()
		CPUReg.AF.h += uint8(gmmu.Memory[int(CPUReg.PC.f()+1)])

	case 0xC6:

	case 0xff:
		// RST n = CALL n*8

		CPUReg.PC.Set(0x0038)

	default:
		fmt.Printf("%08b \n", opcode)
		log.Fatal(gdebug.WhatIsThisCode(int(opcode), false))

	}

	CPUReg.PC.Increment()

}
