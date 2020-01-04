package parser

type InstAttr struct {
	/* Jcc—Jump if Condition Is Met */
	is_jcc bool
	/* fill the output/input/outin operand of the instruction */
	fillcb func(inst *Instruction, ops []Operand)
	/*
	 * if beautify is true, it means that current instruction will modify EFLAGS STATUS FLAGS,
	 * and the instruction have an ability to make the expression more meaningful.
	 *
	 * return nil means can't beautify..
	 */
	beautify func(cmpinst, jccinst *Instruction) Expression
}

func (this *InstAttr) IsJcc() bool {
	return this.is_jcc
}

func (this *InstAttr) Beautifier() func(cmpinst, jccinst *Instruction) Expression {
	return this.beautify
}

const /* instruction */ (
	INST_ADD    = "add"
	INST_CDQ    = "cdq"
	INST_CMP    = "cmp"
	INST_IDIV   = "idiv"
	INST_JE     = "je"
	INST_JL     = "jl"
	INST_JLE    = "jle"
	INST_JMP    = "jmp"
	INST_JNE    = "jne"
	INST_JNS    = "jns"
	INST_LEA    = "lea"
	INST_MOV    = "mov"
	INST_MOVSXD = "movsxd"
	INST_MOVZX  = "movzx"
	INST_MOVSX  = "movsx"
	INST_NEG    = "neg"
	INST_POP    = "pop"
	INST_PUSH   = "push"
	INST_RET    = "ret"
	INST_SUB    = "sub"
	INST_TEST   = "test"
	INST_JS     = "js"
	INST_JG     = "jg"
	INST_SHL    = "shl"
	INST_IMUL   = "imul"
	INST_SETE   = "sete"
	INST_XOR    = "xor"
	INST_JA     = "ja"
	INST_JBE    = "jbe"
)

var g_inst_attr_map = map[string]*InstAttr{
	INST_ADD:    {fillcb: fillADDInstruction},
	INST_XOR:    {fillcb: fillADDInstruction},
	INST_SHL:    {fillcb: fillADDInstruction},
	INST_CDQ:    {fillcb: fillCDQInstruction},
	INST_CMP:    {beautify: cmpBeautifierfunc, fillcb: fillCMPInstruction},
	INST_IDIV:   {fillcb: fillIDIVInstruction},
	INST_IMUL:   {fillcb: fillIMULInstruction},
	INST_JE:     {is_jcc: true, fillcb: fillJccInstruction},
	INST_JA:     {is_jcc: true, fillcb: fillJccInstruction},
	INST_JS:     {is_jcc: true, fillcb: fillJccInstruction},
	INST_JG:     {is_jcc: true, fillcb: fillJccInstruction},
	INST_JL:     {is_jcc: true, fillcb: fillJccInstruction},
	INST_JLE:    {is_jcc: true, fillcb: fillJccInstruction},
	INST_JNE:    {is_jcc: true, fillcb: fillJccInstruction},
	INST_JNS:    {is_jcc: true, fillcb: fillJccInstruction},
	INST_JBE:    {is_jcc: true, fillcb: fillJccInstruction},
	INST_JMP:    {fillcb: fillJccInstruction},
	INST_LEA:    {fillcb: fillMOVInstruction},
	INST_MOV:    {fillcb: fillMOVInstruction},
	INST_MOVSXD: {fillcb: fillMOVInstruction},
	INST_MOVZX:  {fillcb: fillMOVInstruction},
	INST_MOVSX:  {fillcb: fillMOVInstruction},
	INST_NEG:    {fillcb: fillNEGInstruction},
	INST_POP:    {fillcb: fillPOPInstruction},
	INST_PUSH:   {fillcb: fillPUSHInstruction},
	INST_RET:    {fillcb: fillRETInstruction},
	INST_SUB:    {beautify: cmpBeautifierfunc, fillcb: fillADDInstruction},
	INST_TEST:   {beautify: testBeautifierfunc, fillcb: fillCMPInstruction},
	INST_SETE:   {fillcb: fillSETEInstruction},
}

func testBeautifierfunc(cmpinst, jccinst *Instruction) Expression {
	if jccinst.Instmnem != INST_JNE && jccinst.Instmnem != INST_JE {
		return nil
	}
	if cmpinst.Input[0] != cmpinst.Input[1] {
		return nil
	}
	expr := &CompExpression{
		Op:    COMP_OP_EQUAL,
		Left:  &cmpinst.Input[0],
		Right: NewImmeOperand(0),
	}
	if jccinst.Instmnem == INST_JNE {
		expr.Not()
	}
	return expr
}

func cmpBeautifierfunc(cmpinst, jccinst *Instruction) Expression {
	cond_map := map[string]string{
		INST_JE:  COMP_OP_EQUAL,
		INST_JG:  COMP_OP_GREATER,
		INST_JL:  COMP_OP_LESS,
		INST_JLE: COMP_OP_LESSEQUAL,
		INST_JNE: COMP_OP_NOTEQUAL,
		INST_JA:  COMP_OP_GREATER,
		INST_JBE: COMP_OP_LESSEQUAL,
	}
	compop, exists := cond_map[jccinst.Instmnem]
	if !exists {
		return nil
	}
	return &CompExpression{
		Op:    compop,
		Left:  &cmpinst.Input[0],
		Right: &cmpinst.Input[1],
	}
}

var EflagsStatusFlags = []string{
	REG_EFLAGS_OF, REG_EFLAGS_SF, REG_EFLAGS_ZF,
	REG_ELFAGS_AF, REG_EFLAGS_CF, REG_EFLAGS_PF,
}

func addEflagsOp(ops *[]Operand, virtregs ...string) {
	for _, reg := range virtregs {
		*ops = append(*ops, *newRegOperand(reg))
	}
	// *ops = append(*ops, *newRegOperand(REG_EFLAGS, false))
	return
}

/*
 * We don't check the validity of operands, such as do ops have enough operands? etc.
 * because we are parsing the output of GDB disassembler, it should be right.
 */
func fillADDInstruction(inst *Instruction, ops []Operand) {
	inst.Input = append(inst.Input, ops...)
	inst.Output = append(inst.Output, ops[0])
	addEflagsOp(&inst.Output, EflagsStatusFlags...)
	return
}

func fillSETEInstruction(inst *Instruction, ops []Operand) {
	/* Set byte if equal (ZF=1). */
	inst.Input = append(inst.Input, *newRegOperand(REG_EFLAGS_ZF))
	inst.Output = append(inst.Output, ops[0])
}

func fillCDQInstruction(inst *Instruction, ops []Operand) {
	inst.Input = append(inst.Input, *newRegOperand(REG_EAX))
	inst.Output = append(inst.Output, *newRegOperand(REG_EDX))
}

func fillCMPInstruction(inst *Instruction, ops []Operand) {
	inst.Input = append(inst.Input, ops...)
	addEflagsOp(&inst.Output, EflagsStatusFlags...)
}

func fillIMULInstruction(inst *Instruction, ops []Operand) {
	switch len(ops) {
	case 1:
		inst.Input = append(inst.Input, ops[0])
		switch ops[0].Size() {
		case 8:
			inst.Input = append(inst.Input, *newRegOperand(REG_AL))
			inst.Output = append(inst.Output, *newRegOperand(REG_AX))
		case 16:
			inst.Input = append(inst.Input, *newRegOperand(REG_AX))
			inst.Output = append(inst.Output, *newRegOperand(REG_DX))
			inst.Output = append(inst.Output, *newRegOperand(REG_AX))
		case 32:
			inst.Input = append(inst.Input, *newRegOperand(REG_EAX))
			inst.Output = append(inst.Output, *newRegOperand(REG_EDX))
			inst.Output = append(inst.Output, *newRegOperand(REG_EAX))
		case 64:
			inst.Input = append(inst.Input, *newRegOperand(REG_RAX))
			inst.Output = append(inst.Output, *newRegOperand(REG_RDX))
			inst.Output = append(inst.Output, *newRegOperand(REG_RAX))
		default:
			panic("unreachable")
		}
		addEflagsOp(&inst.Output, EflagsStatusFlags...)
	case 2:
		fillADDInstruction(inst, ops)
	case 3:
		inst.Input = append(inst.Input, ops[1], ops[2])
		inst.Output = append(inst.Output, ops[0])
		addEflagsOp(&inst.Output, EflagsStatusFlags...)
	default:
		panic("unreachable")
	}
}

func fillIDIVInstruction(inst *Instruction, ops []Operand) {
	var inputregs []string
	var outputregs []string
	switch ops[0].Size() {
	case 8:
		inputregs = append(inputregs, REG_AX)
		outputregs = append(outputregs, REG_AL, REG_AH)
	case 16:
		inputregs = append(inputregs, REG_DX, REG_AX)
		outputregs = append(outputregs, REG_DX, REG_AX)
	case 32:
		inputregs = append(inputregs, REG_EDX, REG_EAX)
		outputregs = append(outputregs, REG_EDX, REG_EAX)
	case 64:
		inputregs = append(inputregs, REG_RDX, REG_RAX)
		outputregs = append(outputregs, REG_RDX, REG_RAX)
	default:
		panic("unreachable")
	}
	for _, reg := range inputregs {
		inst.Input = append(inst.Input, *newRegOperand(reg))
	}
	inst.Input = append(inst.Input, ops[0])
	for _, reg := range outputregs {
		inst.Output = append(inst.Output, *newRegOperand(reg))
	}
	addEflagsOp(&inst.Output, EflagsStatusFlags...)
}

func fillJccInstruction(inst *Instruction, ops []Operand) {
	var sflags_used []string
	switch inst.Instmnem {
	case INST_JLE:
		fallthrough
	case INST_JG:
		sflags_used = append(sflags_used, REG_EFLAGS_ZF)
		fallthrough
	case INST_JL:
		sflags_used = append(sflags_used, REG_EFLAGS_SF, REG_EFLAGS_OF)
	case INST_JNE:
		fallthrough
	case INST_JE:
		sflags_used = append(sflags_used, REG_EFLAGS_ZF)
	case INST_JNS:
		fallthrough
	case INST_JS:
		sflags_used = append(sflags_used, REG_EFLAGS_SF)
	case INST_JA:
		/* Jump short if above (CF=0 and ZF=0). */
		sflags_used = append(sflags_used, REG_EFLAGS_CF, REG_EFLAGS_ZF)
	case INST_JBE:
		/* CF=1 or ZF=1 */
		sflags_used = append(sflags_used, REG_EFLAGS_CF, REG_EFLAGS_ZF)
	case INST_JMP:
		/* don't need eflags status flags */
	default:
		panic("unreacheable: " + inst.Instmnem)
	}
	inst.Input = append(inst.Input, ops...)
	for _, reg := range sflags_used {
		inst.Input = append(inst.Input, *newRegOperand(reg))
	}
}

func fillMOVInstruction(inst *Instruction, ops []Operand) {
	inst.Input = append(inst.Input, ops[1])
	inst.Output = append(inst.Output, ops[0])
}

func fillNEGInstruction(inst *Instruction, ops []Operand) {
	inst.Output = append(inst.Output, ops[0])
	inst.Input = append(inst.Input, ops[0])
	addEflagsOp(&inst.Output, EflagsStatusFlags...)
}

func fillPOPInstruction(inst *Instruction, ops []Operand) {
	inst.Output = append(inst.Output, ops[0], *newRegOperand(REG_RSP))
}

func fillPUSHInstruction(inst *Instruction, ops []Operand) {
	inst.Input = append(inst.Input, ops[0])
	memop := newOperand(MEMORY_OPERAND)
	memop.Imme = GetOpSizeRepr(ops[0].Size())
	memop.Offset.Base = REG_RSP
	inst.Output = append(inst.Output, *memop, *newRegOperand(REG_RSP))
}

func fillRETInstruction(inst *Instruction, ops []Operand) {
	inst.Input = append(inst.Input, ops...)
	/* we should add RSP, RIP to inst.Output, but we don't care.. */
}
