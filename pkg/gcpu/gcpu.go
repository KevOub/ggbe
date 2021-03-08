package gcpu

import (
	"fmt"
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
	f uint16
}

func (reg SplitRegister) l() uint16 {
	return (reg.f & 255)
}

//Change takes the value and updates f()
func (reg SplitRegister) Change(num uint16) {
	reg.f = num
	// TODO broken...
}

//Get gets the value from the register
func (reg SplitRegister) Get() uint16 {
	return reg.f
}

func (reg SplitRegister) h() uint16 {
	return (reg.f & 65280)
}

// Registers are the cpu registers
type Registers struct {
	AF SplitRegister
	BC SplitRegister
	DE SplitRegister
	HL SplitRegister
	SP int
	PC int
}

//Show displays the registers nicely
func (reg Registers) Show() {
	// HEADER
	fmt.Println("ADDR\tHIGH\tLOW\tTOTAL\tTOTAL")
	var value = ""
	// AF
	fmt.Print("AF\t")
	value = fmt.Sprintf("%04x", reg.AF.h())
	fmt.Printf("%s\t", value)

	value = fmt.Sprintf("%04x", reg.AF.l())
	fmt.Printf("%s\t", value)

	value = fmt.Sprintf("%04x", reg.AF.f)
	fmt.Printf("%s\t", value)

	fmt.Printf("%d", reg.AF.f)
	fmt.Println()

	// BC
	fmt.Print("BC\t")
	value = fmt.Sprintf("%04x", reg.BC.h())
	fmt.Printf("%s\t", value)

	value = fmt.Sprintf("%04x", reg.BC.l())
	fmt.Printf("%s\t", value)

	value = fmt.Sprintf("%04x", reg.BC.f)
	fmt.Printf("%s\t", value)

	fmt.Printf("%d", reg.BC.f)
	fmt.Println()

	// DE
	fmt.Print("DE\t")
	value = fmt.Sprintf("%04x", reg.DE.h())
	fmt.Printf("%s\t", value)

	value = fmt.Sprintf("%04x", reg.DE.l())
	fmt.Printf("%s\t", value)

	value = fmt.Sprintf("%04x", reg.DE.f)
	fmt.Printf("%s\t", value)

	fmt.Printf("%d", reg.DE.f)
	fmt.Println()

	// HL
	fmt.Print("HL\t")
	value = fmt.Sprintf("%04x", reg.HL.h())
	fmt.Printf("%s\t", value)

	value = fmt.Sprintf("%04x", reg.HL.l())
	fmt.Printf("%s\t", value)

	value = fmt.Sprintf("%04x", reg.HL.f)
	fmt.Printf("%s\t", value)

	fmt.Printf("%d", reg.HL.f)

	fmt.Println()

	// SP
	fmt.Print("SP\t")
	fmt.Printf("-\t")

	fmt.Printf("-\t")

	value = fmt.Sprintf("%04x", reg.SP)
	fmt.Printf("%s\t", value)

	fmt.Printf("%d", reg.SP)

	fmt.Println()

	// PC
	fmt.Print("PC\t")
	fmt.Printf("-\t")

	fmt.Printf("-\t")

	value = fmt.Sprintf("%04x", reg.SP)
	fmt.Printf("%s\t", value)

	fmt.Printf("%d", reg.HL.f)

	fmt.Println()

}

// CPUReg are the live registers
var CPUReg = Registers{SplitRegister{0}, SplitRegister{0}, SplitRegister{0}, SplitRegister{0}, 0, 0}

// InitCPU creates cpu
func InitCPU(debug bool) {
	DEBUG = debug
	fmt.Println("CPU Started")
	if DEBUG {
		fmt.Println("\t debugging active")
	}
}

// DoInstruction does instruction
func DoInstruction(opcode uint8) {
	switch opcode {
	case 0x00:
		fmt.Println(":)")
		// fmt.Printf("%02x\n", opcode)
	default:
		fmt.Printf("%02x\n is not implemented", opcode)

	}
}
