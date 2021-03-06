package gcpu

import "fmt"

/* // SplitRegister split from 16 bits to 8 bits
type SplitRegister struct {
	h uint16
	l uint16
}

// Function that combines the high and low bits to create
func (r SplitRegister) f() uint16 {
	return ((r.h << 8) | (r.l))
} */

// SplitRegister split from 16 bits to 8 bits
type SplitRegister struct {
	f uint16
}

func (reg SplitRegister) l() uint16 {
	return (reg.f & 255)
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

func (reg Registers) show() {
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
var CPUReg = Registers{SplitRegister{65535}, SplitRegister{0}, SplitRegister{0}, SplitRegister{0}, 0, 0}

// InitCPU creates cpu
func InitCPU() {

}

// DoInstruction does instruction
func DoInstruction() {

}
