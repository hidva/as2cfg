package parser

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/* Mnemonics and Register is always lower case. */
type MemOffset struct {
	Disp  int64
	Base  string
	Index string
	Scale uint8
}

const (
	REGISTER_OPERAND = iota
	IMMEDIATE_OPERAND
	MEMORY_OPERAND
)

/*
 * The kind field determines what kind of current Operand object.
 * REGISTER_OPERAND; the register name saved in reg is the operand.
 * IMMEDIATE_OPERAND; the immediate operand is saved in imme.
 * MEMORY_OPERAND; Segment Selector is saved in reg and offset is represented by offset,
 * and empty reg means that we don't specify a segment selector explicitly.
 * The value saved in Imme represent operand size, 0 means that the operand size isn't specified, such as in `lea`.
 *
 * Implicit is true means that current operand doesn't directly appear on one instruction.
 * such as 'edx', 'eax' is an implicit operand of div.
 *
 * How to handle AL, AH, AX, EAX, etc in generating SSAname?
 *
 * First, we construct a set of registers for each register, the register and its set share the same underlying storage,
 * it means that when we change the value saved in this register, the value of all registers of its set will change too.
 * Now when we save the register in Operand object, we just save the register used in input instructions.
 * And when we find that a register has changed, we increase the SSA version number of this register and of all registers of its set.
 */
type Operand struct {
	Kind int
	/* the immediate operand output by gdb disassemble always is an unsigned value. */
	Imme   uint64
	Reg    string
	Offset MemOffset

	Implicit bool
	SSAname  string
}

func (this *Operand) IsReg(reg string) bool {
	return this.Kind == REGISTER_OPERAND && this.Reg == reg
}

func RegIsEflagsStatusFlags(reg string) bool {
	for _, ereg := range EflagsStatusFlags {
		if reg == ereg {
			return true
		}
	}
	return false
}

func (this *Operand) IsEflagsStatusFlags() bool {
	if this.Kind != REGISTER_OPERAND {
		return false
	}
	return RegIsEflagsStatusFlags(this.Reg)
}

func (this *Operand) String() string {
	if this.Kind == IMMEDIATE_OPERAND {
		return fmt.Sprintf("0x%x", this.Imme)
	}
	return this.SSAname
}

/* just as a hint, may be wrong... */
func GetOpSizeRepr(size int) uint64 {
	switch size {
	case 8:
		return BYTE_PTR
	case 16:
		return WORD_PTR
	case 32:
		return DWORD_PTR
	case 48:
		return FWORD_PTR
	case 64:
		return QWORD_PTR
	case 80:
		return TBYTE_PTR
	case 128:
		return OWORD_PTR
	case 256:
		return YMMWORD_PTR
	case 512:
		return ZMMWORD_PTR
	}
	panic("unreachable")
	return 0
}

/* bit size of operand */
func (this *Operand) Size() int {
	switch this.Kind {
	case IMMEDIATE_OPERAND:
		/* INTEL MANUL: but can never be greater than the maximum value of an unsigned doubleword integer (2 ** 32). */
		return 32
	case REGISTER_OPERAND:
		regattr, exists := g_reg_attr_map[this.Reg]
		if !exists {
			panic(fmt.Errorf("Can't know the attribute of register '%s', "+
				"you can add it in g_reg_attr_map and submit a PR if you want to", this.Reg))
		}
		return regattr.Size
	case MEMORY_OPERAND:
		switch this.Imme {
		default:
			panic(fmt.Errorf("can't get the operand size. Imme: %d", this.Imme))
		case BYTE_PTR:
			return 1 * 8
		case DWORD_PTR:
			return 4 * 8
		case FWORD_PTR:
			return 6 * 8
		case OWORD_PTR:
			return 16 * 8
		case QWORD_PTR:
			return 8 * 8
		case TBYTE_PTR:
			return 10 * 8
		case WORD_PTR:
			return 2 * 8
		case XMMWORD_PTR:
			return 128
		case YMMWORD_PTR:
			return 256
		case ZMMWORD_PTR:
			return 512
		}
	default:
		panic("unreachable")
	}
	return 0
}

type Instruction struct {
	Instprefix string
	Instmnem   string
	/* the order of operands represents the modification order. */
	Output []Operand
	/* the order of operands is in left-to-right */
	Input []Operand
}

func ops2strs(ops []Operand) []string {
	var ret []string
	for idx := range ops {
		if ops[idx].IsEflagsStatusFlags() {
			continue
		}
		ret = append(ret, ops[idx].String())
	}
	return ret
}

func (this *Instruction) String() string {
	outops := ops2strs(this.Output)
	inops := ops2strs(this.Input)
	return fmt.Sprintf("%s = %s %s(%s)", strings.Join(outops, ","), this.Instprefix, this.Instmnem, strings.Join(inops, ","))
}

func hasReg(ops []Operand, reg string) bool {
	for idx := range ops {
		if ops[idx].IsReg(reg) {
			return true
		}
	}
	return false
}

func (this *Instruction) GetInputReg(reg string) *Operand {
	for idx := range this.Input {
		if this.Input[idx].IsReg(reg) {
			return &this.Input[idx]
		}
	}
	return nil
}

func (this *Instruction) WriteReg(reg string) bool {
	return hasReg(this.Output, reg)
}

func (this *Instruction) Use(reg string) bool {
	return hasReg(this.Input, reg)
}

func (this *Instruction) GetInstAttr() *InstAttr {
	return GetMnemAttr(this.Instmnem)
}

type AddressedInst struct {
	Addr uint64
	Inst Instruction
}

func Parse(input io.Reader) []AddressedInst {
	lexer := newLexer(input)
	ret := yyParse(lexer)
	if ret != 0 {
		panic(fmt.Errorf("parse error: %d", ret))
	}
	return lexer.insts
}

func newRegOperand(reg string) *Operand {
	newop := newOperand(REGISTER_OPERAND)
	newop.Reg = reg
	return newop
}

func NewImmeOperand(imme uint64) *Operand {
	newop := newOperand(IMMEDIATE_OPERAND)
	newop.Imme = imme
	return newop
}

func newOperand(kind int) *Operand {
	return &Operand{
		Kind: kind,
		/* the invocation of this function means that we are constructing a
		   register operand that doesn't exist in instruction. */
		Implicit: true,
	}
}

func newInstruction(mnem string, ops ...Operand) *Instruction {
	/* Look 'Intel® 64 and IA-32 Architectures Software Developer’s Manual Volume 2 (2A, 2B, 2C & 2D): Instruction Set Reference, A-Z'
	   for the introduction of one instruction. */
	inst := &Instruction{
		Instmnem: mnem,
	}
	regattr := g_inst_attr_map[mnem]
	if regattr == nil {
		regattr = g_def_instattr
		fmt.Fprintf(os.Stderr, "WARNING: don't know how to parse instruction '%s', use the default method.\n"+
			"And you can add the logic for this instruction in newInstruction(), and submit a PR if you want to.\n", mnem)
	}
	regattr.fillcb(inst, ops)
	return inst
}

var g_ignored_insts = map[string]bool{
	"nop": true,
}
