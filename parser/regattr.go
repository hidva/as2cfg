package parser

import "fmt"

const /* register */ (
	REG_AL  = "al"
	REG_AH  = "ah"
	REG_CL  = "cl"
	REG_DL  = "dl"
	REG_AX  = "ax"
	REG_DX  = "dx"
	REG_EAX = "eax"
	REG_ECX = "ecx"
	REG_EDI = "edi"
	REG_EDX = "edx"
	REG_RAX = "rax"
	REG_RBP = "rbp"
	REG_RCX = "rcx"
	REG_RDX = "rdx"
	REG_RSI = "rsi"
	REG_RSP = "rsp"

	REG_EFLAGS    = "eflags"
	REG_EFLAGS_OF = "eflags_of"
	REG_EFLAGS_SF = "eflags_sf"
	REG_EFLAGS_ZF = "eflags_zf"
	REG_ELFAGS_AF = "eflags_af"
	REG_EFLAGS_CF = "eflags_cf"
	REG_EFLAGS_PF = "eflags_pf"
)

type RegAttr struct {
	Size int /* size in bit */
	/* a set of regs that any reg in this set has the same storage as current register. */
	RelatedRegs map[string]bool
}

func GetRegAttr(reg string) *RegAttr {
	return g_reg_attr_map[reg]
}

var g_reg_attr_map = map[string]*RegAttr{
	REG_AL: {
		Size: 1 * 8,
		RelatedRegs: map[string]bool{
			REG_AX:  true,
			REG_EAX: true,
			REG_RAX: true,
		},
	},
	REG_AH: {
		Size: 1 * 8,
		RelatedRegs: map[string]bool{
			REG_AX:  true,
			REG_EAX: true,
			REG_RAX: true,
		},
	},
	REG_CL: {
		Size: 1 * 8,
		RelatedRegs: map[string]bool{
			REG_ECX: true,
			REG_RCX: true,
		},
	},
	REG_DL: {
		Size: 1 * 8,
		RelatedRegs: map[string]bool{
			REG_DX:  true,
			REG_EDX: true,
			REG_RDX: true,
		},
	},
	REG_AX: {
		Size: 2 * 8,
		RelatedRegs: map[string]bool{
			REG_AL:  true,
			REG_AH:  true,
			REG_EAX: true,
			REG_RAX: true,
		},
	},
	REG_DX: {
		Size: 2 * 8,
		RelatedRegs: map[string]bool{
			REG_DL:  true,
			REG_EDX: true,
			REG_RDX: true,
		},
	},
	REG_EAX: {
		Size: 4 * 8,
		RelatedRegs: map[string]bool{
			REG_AL:  true,
			REG_AH:  true,
			REG_AX:  true,
			REG_RAX: true,
		},
	},
	REG_ECX: {
		Size: 4 * 8,
		RelatedRegs: map[string]bool{
			REG_CL:  true,
			REG_RCX: true,
		},
	},
	REG_EDI: {
		Size:        4 * 8,
		RelatedRegs: map[string]bool{},
	},
	REG_EDX: {
		Size: 4 * 8,
		RelatedRegs: map[string]bool{
			REG_DL:  true,
			REG_DX:  true,
			REG_RDX: true,
		},
	},
	REG_RAX: {
		Size: 8 * 8,
		RelatedRegs: map[string]bool{
			REG_AL:  true,
			REG_AH:  true,
			REG_AX:  true,
			REG_EAX: true,
		},
	},
	REG_RBP: {
		Size:        8 * 8,
		RelatedRegs: map[string]bool{},
	},
	REG_RCX: {
		Size: 8 * 8,
		RelatedRegs: map[string]bool{
			REG_ECX: true,
			REG_CL:  true,
		},
	},
	REG_RDX: {
		Size: 8 * 8,
		RelatedRegs: map[string]bool{
			REG_DL:  true,
			REG_DX:  true,
			REG_EDX: true,
		},
	},
	REG_RSI: {
		Size:        8 * 8,
		RelatedRegs: map[string]bool{},
	},
	REG_RSP: {
		Size:        8 * 8,
		RelatedRegs: map[string]bool{},
	},
}

func init() {
	/* check the possible configuration error in g_reg_attr_map */
	for reg, regattr := range g_reg_attr_map {
		for relatedreg := range regattr.RelatedRegs {
			relatedregattr, exists := g_reg_attr_map[relatedreg]
			if !exists {
				panic(fmt.Errorf("lack '%s' in g_reg_attr_map", relatedreg))
			}
			_, relatedtoo := relatedregattr.RelatedRegs[reg]
			if !relatedtoo {
				panic(fmt.Errorf("should put '%s' in RelatedRegs of '%s'", reg, relatedreg))
			}
		}
	}
}
