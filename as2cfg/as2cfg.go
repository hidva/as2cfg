package main

import (
	"flag"
	"fmt"
	"github.com/awalterschulze/gographviz"
	"github.com/hidva/as2cfg/parser"
	"io"
	"os"
	"strconv"
	"strings"
)

type CFGNode struct {
	id       int /* begin from 0 */
	insts    []*parser.AddressedInst
	from, to []*CFGEdge
}

func (this *CFGNode) getDOTLabel() string {
	var inststrs []string
	for _, inst := range this.insts {
		inststrs = append(inststrs, inst.Inst.String())
	}
	return strings.Join(inststrs, `\n`)
}

func (this *CFGNode) getDOTName() string {
	return strconv.Itoa(this.id)
}

func (this *CFGNode) addInst(inst *parser.AddressedInst) {
	this.insts = append(this.insts, inst)
}

func (this *CFGNode) contain(addr uint64) bool {
	if len(this.insts) <= 0 {
		return false
	}
	return addr >= this.insts[0].Addr && addr <= this.insts[len(this.insts)-1].Addr
}

func (this *CFGNode) find(addr uint64) int {
	for idx := range this.insts {
		if this.insts[idx].Addr == addr {
			return idx
		}
	}
	return -1
}

/* a CFGEdge means that the control flow will go to `to` from `from` when the `expr` is true. */
type CFGEdge struct {
	from, to *CFGNode
	expr     parser.Expression
}

/* The same CFGEdge will have the same address, the address of a CFGEdge will be used as the ID of the CFGEdge. */
type CFGGraph struct {
	nodes []*CFGNode
}

func newCFGGraph() *CFGGraph {
	return &CFGGraph{}
}

func (this *CFGGraph) find(addr uint64) *CFGNode {
	for _, node := range this.nodes {
		if node.contain(addr) {
			return node
		}
	}
	return nil
}

/* Split container at the address specified by addr, and return new CFGNode which contains addr */
func (this *CFGGraph) split(container *CFGNode, addr uint64) *CFGNode {
	idx := container.find(addr)
	if idx < 0 {
		panic("bad input")
	}
	if idx == 0 {
		return container
	}
	newnode := this.newCFGNode()

	newnode.insts = append(newnode.insts, container.insts[idx:]...)
	container.insts = container.insts[:idx]

	newnode.from = container.from
	container.from = nil
	for _, edge := range newnode.from {
		edge.from = newnode
	}

	newedge := &CFGEdge{
		from: container,
		to:   newnode,
		expr: &parser.Always{},
	}
	this.addEdge(newedge)
	return newnode
}

func (this *CFGGraph) newCFGNode() *CFGNode {
	node := &CFGNode{
		id: len(this.nodes),
	}
	this.nodes = append(this.nodes, node)
	return node
}

func (this *CFGGraph) addEdge(edge *CFGEdge) {
	edge.from.from = append(edge.from.from, edge)
	edge.to.to = append(edge.to.to, edge)
}

func isJmp(inst *parser.Instruction) bool {
	return inst.Instmnem == parser.INST_JMP
}

func getJMPDst(jcc *parser.Instruction) uint64 {
	op := &jcc.Input[0]
	if op.Kind != parser.IMMEDIATE_OPERAND {
		panic("unsupported jmp")
	}
	return op.Imme
}

func getJMPExpr(inst *parser.Instruction) parser.Expression {
	switch inst.Instmnem {
	case parser.INST_JE: /* ZF = 1 */
		return &parser.CompExpression{
			Op:    parser.COMP_OP_EQUAL,
			Left:  inst.GetInputReg(parser.REG_EFLAGS_ZF),
			Right: parser.NewImmeOperand(1),
		}
	case parser.INST_JS: /* SF = 1 */
		return &parser.CompExpression{
			Op:    parser.COMP_OP_EQUAL,
			Left:  inst.GetInputReg(parser.REG_EFLAGS_SF),
			Right: parser.NewImmeOperand(1),
		}
	case parser.INST_JG: /* ZF=0 and SF=OF */
		return &parser.LogicalExpr{
			Op: parser.LOGICAL_OP_AND,
			Left: &parser.CompExpression{
				Op:    parser.COMP_OP_EQUAL,
				Left:  inst.GetInputReg(parser.REG_EFLAGS_ZF),
				Right: parser.NewImmeOperand(0),
			},
			Right: &parser.CompExpression{
				Op:    parser.COMP_OP_EQUAL,
				Left:  inst.GetInputReg(parser.REG_EFLAGS_SF),
				Right: inst.GetInputReg(parser.REG_EFLAGS_OF),
			},
		}
	case parser.INST_JL: /* SF != OF */
		return &parser.CompExpression{
			Op:    parser.COMP_OP_NOTEQUAL,
			Left:  inst.GetInputReg(parser.REG_EFLAGS_SF),
			Right: inst.GetInputReg(parser.REG_EFLAGS_OF),
		}
	case parser.INST_JLE: /* ZF=1 or SF != OF */
		return &parser.LogicalExpr{
			Op: parser.LOGICAL_OP_OR,
			Left: &parser.CompExpression{
				Op:    parser.COMP_OP_EQUAL,
				Left:  inst.GetInputReg(parser.REG_EFLAGS_ZF),
				Right: parser.NewImmeOperand(1),
			},
			Right: &parser.CompExpression{
				Op:    parser.COMP_OP_NOTEQUAL,
				Left:  inst.GetInputReg(parser.REG_EFLAGS_SF),
				Right: inst.GetInputReg(parser.REG_EFLAGS_OF),
			},
		}
	case parser.INST_JNE: /* ZF = 0 */
		return &parser.CompExpression{
			Op:    parser.COMP_OP_EQUAL,
			Left:  inst.GetInputReg(parser.REG_EFLAGS_ZF),
			Right: parser.NewImmeOperand(0),
		}
	case parser.INST_JNS: /* SF=0 */
		return &parser.CompExpression{
			Op:    parser.COMP_OP_EQUAL,
			Left:  inst.GetInputReg(parser.REG_EFLAGS_SF),
			Right: parser.NewImmeOperand(0),
		}
	case parser.INST_JA: /* CF=0 and ZF=0 */
		return &parser.LogicalExpr{
			Op: parser.LOGICAL_OP_AND,
			Left: &parser.CompExpression{
				Op:    parser.COMP_OP_EQUAL,
				Left:  inst.GetInputReg(parser.REG_EFLAGS_CF),
				Right: parser.NewImmeOperand(0),
			},
			Right: &parser.CompExpression{
				Op:    parser.COMP_OP_EQUAL,
				Left:  inst.GetInputReg(parser.REG_EFLAGS_ZF),
				Right: parser.NewImmeOperand(0),
			},
		}
	case parser.INST_JBE: /* CF=1 or ZF=1*/
		return &parser.LogicalExpr{
			Op: parser.LOGICAL_OP_OR,
			Left: &parser.CompExpression{
				Op:    parser.COMP_OP_EQUAL,
				Left:  inst.GetInputReg(parser.REG_EFLAGS_CF),
				Right: parser.NewImmeOperand(1),
			},
			Right: &parser.CompExpression{
				Op:    parser.COMP_OP_EQUAL,
				Left:  inst.GetInputReg(parser.REG_EFLAGS_ZF),
				Right: parser.NewImmeOperand(1),
			},
		}
	}
	panic("unreacheable")
	return nil
}

func isEnd(inst *parser.Instruction) bool {
	return inst.Instmnem == parser.INST_RET
}

/*
 * How to construct a CFGEdge? we use status flags in EFLAGS such as CF, ZF, PF, etc. because of the existence of
 * SETC-like instruction which can modify status flags individually, we can't use the eflags as a whole..
 * So now we handle the status flags individually, we add each status flag which will be modified by the instruction
 * to the Output field of the Instruction, and add any status flag to Input if the instruction
 * will use this status flag.
 *
 * We will attempt to beautify the expression used in CFGEdge:
 *
 * ```
 * cfgnode1:
 *  cmp rax, rbx
 *  jg cfgnode2
 * ```
 *
 * the expression of CFGEdge introduced by JG is `rax > rbx`, not `ZF = 0 && (SF = OF)`,
 * `rax > rbx` is more meanful.
 *
 * There is a lot of states when we construct the CFG, so carefully here!
 */
type CFGContext struct {
	/* current CFGNode, instruction will be pushed in this node.
	   nil means that we should construct a new CFGNode. */
	curnode *CFGNode
	/* if pending_edge isn't nil, it means that the next new CFGNode will be `to` field of pending_edge,
	   and the pending_edge should be added to CFGGraph.  */
	pending_edge *CFGEdge
	/* if newnodes[Addr] exists, it means that instruction after Addr should be pushed to newnodes[Addr]. */
	newnodes map[uint64]*CFGNode
	/* last_inst[REG] records that the last instruction that modifies REG. */
	last_inst map[string]*parser.AddressedInst

	/*
	 * The key of ssamap is a variable, and the value is the modified version of this variable.
	 * because the SSAname makes sense only when they are in the same CFGnode,
	 * so we will clear ssamap when a new cfgnode created.
	 *
	 * if variable is a register, the key is `REG:$RegisterName`.
	 * if variable is in memory, the key is `MEM:$MemoryRepInIntel`.
	 */
	ssamap map[string]uint64

	graph   *CFGGraph
	curinst *parser.AddressedInst
}

func (context *CFGContext) beginNode(node *CFGNode) {
	context.curnode = node
	context.ssamap = make(map[string]uint64)
	// clear last_inst?
	if context.pending_edge != nil {
		context.pending_edge.to = context.curnode
		context.graph.addEdge(context.pending_edge)
		context.pending_edge = nil
	}
}

func (context *CFGContext) endNode() {
	context.curnode = nil
}

func (context *CFGContext) getSSAName(op *parser.Operand) string {
	switch op.Kind {
	case parser.REGISTER_OPERAND:
		return context.getRegSSAName(op.Reg)
	case parser.MEMORY_OPERAND:
		memkey := context.getMEMKey(op)
		version := context.ssamap[memkey]
		return fmt.Sprintf("%s_%d", context.getMEMSSAName(op.Reg, op.Offset), version)
	case parser.IMMEDIATE_OPERAND:
		return fmt.Sprintf("0x%x", op.Imme)
	}
	panic("unreachable")
	return "hidva.com"
}

func (context *CFGContext) getRegSSAName(reg string) string {
	ssakey := getREGKey(reg)
	version := context.ssamap[ssakey]
	return fmt.Sprintf("%s_%d", reg, version)
}

func (context *CFGContext) getMEMSSAName(segreg string, offset parser.MemOffset) string {
	if segreg != "" {
		segreg = context.getRegSSAName(segreg)
	}
	if offset.Base != "" {
		offset.Base = context.getRegSSAName(offset.Base)
	}
	if offset.Index != "" {
		offset.Index = context.getRegSSAName(offset.Index)
	}
	return getMEMRepr(segreg, &offset)
}

func getREGKey(reg string) string {
	return fmt.Sprintf("REG:%s", reg)
}

func getMEMRepr(segreg string, offset *parser.MemOffset) string {
	ret := ""
	if segreg != "" {
		ret = (segreg + ":")
	}
	if offset.Base == "" && offset.Index == "" {
		ret += fmt.Sprintf("0x%x", uint64(offset.Disp))
		return ret
	}
	ret += "["
	if offset.Base != "" {
		ret += offset.Base
	}
	if offset.Index != "" {
		if offset.Base != "" {
			ret += "+"
		}
		ret += fmt.Sprintf("%s*%d", offset.Index, offset.Scale)
	}
	if offset.Disp > 0 {
		ret += fmt.Sprintf("+0x%x", uint64(offset.Disp))
	} else if offset.Disp < 0 {
		ret += fmt.Sprintf("-0x%x", uint64(-offset.Disp))
	}
	ret += "]"
	return ret
}

func (context *CFGContext) getMEMKey(op *parser.Operand) string {
	memrepr := context.getMEMSSAName(op.Reg, op.Offset)
	return fmt.Sprintf("MEM:%s", memrepr)
}

func (context *CFGContext) getSSAKey(op *parser.Operand) string {
	switch op.Kind {
	case parser.REGISTER_OPERAND:
		return getREGKey(op.Reg)
	case parser.MEMORY_OPERAND:
		return context.getMEMKey(op)
	}
	panic("unreachable")
	return "hidva.com"
}

func (context *CFGContext) feedInst(inst *parser.AddressedInst) {
	context.curinst = inst

	newnode, exist := context.newnodes[context.curinst.Addr]
	if exist {
		if context.curnode != nil {
			context.graph.addEdge(&CFGEdge{from: context.curnode, to: newnode, expr: &parser.Always{}})
		}
		context.beginNode(newnode)
		delete(context.newnodes, context.curinst.Addr)
	} else if context.curnode == nil {
		context.beginNode(context.graph.newCFGNode())
	}
	context.curnode.addInst(context.curinst)

	for reg := range context.last_inst {
		if context.curinst.Inst.WriteReg(reg) {
			context.last_inst[reg] = context.curinst
		}
	}

	for inopidx := range context.curinst.Inst.Input {
		inop := &context.curinst.Inst.Input[inopidx]
		inop.SSAname = context.getSSAName(inop)
	}
	for outopidx := range context.curinst.Inst.Output {
		outop := &context.curinst.Inst.Output[outopidx]
		ssakey := context.getSSAKey(outop)
		context.ssamap[ssakey]++
		if outop.Kind == parser.REGISTER_OPERAND {
			regattr := parser.GetRegAttr(outop.Reg)
			if regattr != nil {
				for relreg := range regattr.RelatedRegs {
					context.ssamap[getREGKey(relreg)]++
				}
			}
		}
		outop.SSAname = context.getSSAName(outop)
	}
}

func newCFGContext() *CFGContext {
	ctx := &CFGContext{
		graph:     newCFGGraph(),
		last_inst: make(map[string]*parser.AddressedInst),
		newnodes:  make(map[uint64]*CFGNode),
		ssamap:    make(map[string]uint64),
	}
	for _, reg := range parser.EflagsStatusFlags {
		ctx.last_inst[reg] = nil
	}
	return ctx
}

func (context *CFGContext) addJmp(jmpexp parser.Expression) {
	jmpdst := getJMPDst(&context.curinst.Inst)
	newedge := &CFGEdge{from: context.curnode, to: nil, expr: jmpexp}
	if jmpdst > context.curinst.Addr {
		newnode, exist := context.newnodes[jmpdst]
		if !exist {
			newnode = context.graph.newCFGNode()
			context.newnodes[jmpdst] = newnode
		}
		newedge.to = newnode
	} else {
		oldnode := context.graph.find(jmpdst)
		if oldnode == nil {
			panic(fmt.Errorf("unknown jmpdst: 0x%x", jmpdst))
		}
		newnode := context.graph.split(oldnode, jmpdst)
		newedge.to = newnode
	}
	context.graph.addEdge(newedge)
}

func generateCFG(insts []parser.AddressedInst) *CFGContext {
	context := newCFGContext()
	for instidx := range insts {
		context.feedInst(&insts[instidx])
		if isEnd(&context.curinst.Inst) {
			context.endNode()
			continue
		}
		if isJmp(&context.curinst.Inst) {
			context.addJmp(&parser.Always{})
			context.endNode()
			continue
		}
		if context.curinst.Inst.GetInstAttr().IsJcc() {
			var jmpexpr parser.Expression
			var theinst *parser.AddressedInst
			for sflags := range context.last_inst {
				if !context.curinst.Inst.Use(sflags) {
					continue
				}
				if theinst == nil {
					theinst = context.last_inst[sflags]
				} else if theinst.Addr != context.last_inst[sflags].Addr {
					theinst = nil
					break
				}
			}
			if theinst == nil || theinst.Inst.GetInstAttr().Beautifier() == nil {
				jmpexpr = getJMPExpr(&context.curinst.Inst)
			} else {
				jmpexpr = theinst.Inst.GetInstAttr().Beautifier()(&theinst.Inst, &context.curinst.Inst)
				if jmpexpr == nil {
					jmpexpr = getJMPExpr(&context.curinst.Inst)
				}
			}
			context.addJmp(jmpexpr)

			var notjmpexpr parser.Expression
			jmpcmpexpr, iscmpexpr := jmpexpr.(*parser.CompExpression)
			if !iscmpexpr {
				notjmpexpr = &parser.NotExpression{
					Operand: jmpexpr,
				}
			} else {
				notjmpexpr = &parser.CompExpression{
					Op:    parser.Negate(jmpcmpexpr.Op),
					Left:  jmpcmpexpr.Left,
					Right: jmpcmpexpr.Right,
				}
			}
			context.pending_edge = &CFGEdge{from: context.curnode, to: nil, expr: notjmpexpr}

			context.endNode()
			continue
		}
	}
	return context
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func CFGGraph2Dot(graph *CFGGraph) string {
	vizgraph := gographviz.NewGraph()
	graphAst, _ := gographviz.ParseString(`digraph G {}`)
	if err := gographviz.Analyse(graphAst, vizgraph); err != nil {
		panic(err)
	}
	for _, node := range graph.nodes {
		for _, edgeout := range node.from {
			from := edgeout.from.getDOTName()
			to := edgeout.to.getDOTName()
			must(vizgraph.AddNode("G", from, map[string]string{
				"label": fmt.Sprintf(`"%s"`, edgeout.from.getDOTLabel()),
				"shape": "box",
			}))
			must(vizgraph.AddNode("G", to, map[string]string{
				"label": fmt.Sprintf(`"%s"`, edgeout.to.getDOTLabel()),
				"shape": "box",
			}))
			must(vizgraph.AddEdge(from, to, true, map[string]string{
				"label": fmt.Sprintf(`"%s"`, edgeout.expr.String()),
			}))
		}
	}
	return vizgraph.String()
}

func main() {
	flag.Parse()
	var input io.Reader
	input_file := flag.Arg(0)
	if len(input_file) <= 0 {
		input = os.Stdin
	} else {
		var err error
		input, err = os.Open(input_file)
		if err != nil {
			panic(err)
		}
	}
	insts := parser.Parse(input)
	cfgctx := generateCFG(insts)
	if cfgctx.pending_edge != nil {
		fmt.Fprintf(os.Stderr, "WARNING: pending_edge != nil\n")
	}
	if len(cfgctx.newnodes) > 0 {
		fmt.Fprintf(os.Stderr, "WARNING: newnodes = %#v\n", cfgctx.newnodes)
	}
	fmt.Println(CFGGraph2Dot(cfgctx.graph))
}
