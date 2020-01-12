package parser

import "fmt"

/* There are 186 registers in INTEL CPU, by the way, there are at least 944 instructions.. */
const /* register */ (
	REG_GS    = "gs"
	REG_ZMM23 = "zmm23"
	REG_EDI   = "edi"
	REG_EDX   = "edx"
	REG_R13W  = "r13w"
	REG_RSP   = "rsp"
	REG_R8B   = "r8b"
	REG_R8D   = "r8d"
	REG_DS    = "ds"
	REG_R13D  = "r13d"
	REG_R13B  = "r13b"
	REG_R8W   = "r8w"
	REG_YMM9  = "ymm9"
	REG_YMM8  = "ymm8"
	REG_R14   = "r14"
	REG_R15   = "r15"
	REG_R12   = "r12"
	REG_R13   = "r13"
	REG_R10   = "r10"
	REG_R11   = "r11"
	REG_YMM1  = "ymm1"
	REG_YMM0  = "ymm0"
	REG_YMM3  = "ymm3"
	REG_YMM2  = "ymm2"
	REG_YMM5  = "ymm5"
	REG_YMM4  = "ymm4"
	REG_YMM7  = "ymm7"
	REG_YMM6  = "ymm6"
	REG_ZMM11 = "zmm11"
	REG_DX    = "dx"
	REG_YMM32 = "ymm32"
	REG_YMM31 = "ymm31"
	REG_YMM30 = "ymm30"
	REG_DIL   = "dil"
	REG_R10B  = "r10b"
	REG_RBX   = "rbx"
	REG_BPL   = "bpl"
	REG_R10D  = "r10d"
	REG_XMM10 = "xmm10"
	REG_XMM11 = "xmm11"
	REG_XMM12 = "xmm12"
	REG_XMM13 = "xmm13"
	REG_XMM14 = "xmm14"
	REG_XMM15 = "xmm15"
	REG_XMM16 = "xmm16"
	REG_XMM17 = "xmm17"
	REG_XMM18 = "xmm18"
	REG_XMM19 = "xmm19"
	REG_SIL   = "sil"
	REG_R10W  = "r10w"
	REG_ZMM17 = "zmm17"
	REG_MM5   = "mm5"
	REG_MM4   = "mm4"
	REG_MM7   = "mm7"
	REG_MM6   = "mm6"
	REG_MM1   = "mm1"
	REG_MM0   = "mm0"
	REG_MM3   = "mm3"
	REG_MM2   = "mm2"
	REG_YMM28 = "ymm28"
	REG_YMM29 = "ymm29"
	REG_EBP   = "ebp"
	REG_YMM20 = "ymm20"
	REG_YMM21 = "ymm21"
	REG_YMM22 = "ymm22"
	REG_YMM23 = "ymm23"
	REG_YMM24 = "ymm24"
	REG_YMM25 = "ymm25"
	REG_YMM26 = "ymm26"
	REG_YMM27 = "ymm27"
	REG_R15D  = "r15d"
	REG_R15B  = "r15b"
	REG_ESP   = "esp"
	REG_R15W  = "r15w"
	REG_ESI   = "esi"
	REG_BL    = "bl"
	REG_BH    = "bh"
	REG_XMM2  = "xmm2"
	REG_XMM3  = "xmm3"
	REG_XMM0  = "xmm0"
	REG_XMM1  = "xmm1"
	REG_XMM6  = "xmm6"
	REG_XMM7  = "xmm7"
	REG_XMM4  = "xmm4"
	REG_XMM5  = "xmm5"
	REG_XMM8  = "xmm8"
	REG_XMM9  = "xmm9"
	REG_BX    = "bx"
	REG_ECX   = "ecx"
	REG_DL    = "dl"
	REG_R12W  = "r12w"
	REG_R9D   = "r9d"
	REG_R9B   = "r9b"
	REG_R9    = "r9"
	REG_R12B  = "r12b"
	REG_R12D  = "r12d"
	REG_R9W   = "r9w"
	REG_EBX   = "ebx"
	REG_RDI   = "rdi"
	REG_CH    = "ch"
	REG_CL    = "cl"
	REG_CX    = "cx"
	REG_CS    = "cs"
	REG_RCX   = "rcx"
	REG_AH    = "ah"
	REG_XMM29 = "xmm29"
	REG_XMM28 = "xmm28"
	REG_RSI   = "rsi"
	REG_XMM21 = "xmm21"
	REG_XMM20 = "xmm20"
	REG_XMM23 = "xmm23"
	REG_XMM22 = "xmm22"
	REG_XMM25 = "xmm25"
	REG_XMM24 = "xmm24"
	REG_XMM27 = "xmm27"
	REG_XMM26 = "xmm26"
	REG_ZMM8  = "zmm8"
	REG_ZMM9  = "zmm9"
	REG_ZMM0  = "zmm0"
	REG_ZMM1  = "zmm1"
	REG_ZMM2  = "zmm2"
	REG_ZMM3  = "zmm3"
	REG_ZMM4  = "zmm4"
	REG_ZMM5  = "zmm5"
	REG_ZMM6  = "zmm6"
	REG_ZMM7  = "zmm7"
	REG_ZMM12 = "zmm12"
	REG_ZMM13 = "zmm13"
	REG_ZMM10 = "zmm10"
	REG_R14D  = "r14d"
	REG_ZMM16 = "zmm16"
	REG_R14B  = "r14b"
	REG_ZMM14 = "zmm14"
	REG_ZMM15 = "zmm15"
	REG_ZMM18 = "zmm18"
	REG_ZMM19 = "zmm19"
	REG_RBP   = "rbp"
	REG_R14W  = "r14w"
	REG_SS    = "ss"
	REG_SPL   = "spl"
	REG_DI    = "di"
	REG_BND0  = "bnd0"
	REG_BND1  = "bnd1"
	REG_BND2  = "bnd2"
	REG_BND3  = "bnd3"
	REG_XMM32 = "xmm32"
	REG_R8    = "r8"
	REG_XMM30 = "xmm30"
	REG_XMM31 = "xmm31"
	REG_AL    = "al"
	REG_RDX   = "rdx"
	REG_BP    = "bp"
	REG_AX    = "ax"
	REG_RAX   = "rax"
	REG_DH    = "dh"
	REG_ZMM29 = "zmm29"
	REG_ZMM28 = "zmm28"
	REG_ZMM27 = "zmm27"
	REG_ZMM26 = "zmm26"
	REG_R11B  = "r11b"
	REG_ZMM24 = "zmm24"
	REG_R11D  = "r11d"
	REG_ZMM22 = "zmm22"
	REG_ZMM21 = "zmm21"
	REG_ZMM20 = "zmm20"
	REG_ES    = "es"
	REG_R11W  = "r11w"
	REG_FS    = "fs"
	REG_YMM11 = "ymm11"
	REG_YMM10 = "ymm10"
	REG_YMM13 = "ymm13"
	REG_YMM12 = "ymm12"
	REG_YMM15 = "ymm15"
	REG_YMM14 = "ymm14"
	REG_YMM17 = "ymm17"
	REG_YMM16 = "ymm16"
	REG_YMM19 = "ymm19"
	REG_YMM18 = "ymm18"
	REG_EAX   = "eax"
	REG_ZMM30 = "zmm30"
	REG_ZMM31 = "zmm31"
	REG_ZMM32 = "zmm32"
	REG_SP    = "sp"
	REG_SI    = "si"
	REG_ZMM25 = "zmm25"

	REG_EFLAGS    = "eflags"
	REG_EFLAGS_OF = "eflags_of"
	REG_EFLAGS_SF = "eflags_sf"
	REG_EFLAGS_ZF = "eflags_zf"
	REG_EFLAGS_AF = "eflags_af"
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
	REG_GS: {
		Size:        16,
		RelatedRegs: map[string]bool{},
	},

	REG_ZMM23: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM23: true,
			REG_XMM23: true,
		},
	},

	REG_EDI: {
		Size: 32,
		RelatedRegs: map[string]bool{
			REG_DIL: true,
			REG_RDI: true,
			REG_DI:  true,
		},
	},

	REG_EDX: {
		Size: 32,
		RelatedRegs: map[string]bool{
			REG_DH:  true,
			REG_DL:  true,
			REG_DX:  true,
			REG_RDX: true,
		},
	},

	REG_R13W: {
		Size: 16,
		RelatedRegs: map[string]bool{
			REG_R13:  true,
			REG_R13B: true,
			REG_R13D: true,
		},
	},

	REG_RSP: {
		Size: 64,
		RelatedRegs: map[string]bool{
			REG_SPL: true,
			REG_SP:  true,
			REG_ESP: true,
		},
	},

	REG_R8B: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_R8:  true,
			REG_R8D: true,
			REG_R8W: true,
		},
	},

	REG_R8D: {
		Size: 32,
		RelatedRegs: map[string]bool{
			REG_R8:  true,
			REG_R8B: true,
			REG_R8W: true,
		},
	},

	REG_DS: {
		Size:        16,
		RelatedRegs: map[string]bool{},
	},

	REG_R13D: {
		Size: 32,
		RelatedRegs: map[string]bool{
			REG_R13:  true,
			REG_R13B: true,
			REG_R13W: true,
		},
	},

	REG_R13B: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_R13:  true,
			REG_R13D: true,
			REG_R13W: true,
		},
	},

	REG_R8W: {
		Size: 16,
		RelatedRegs: map[string]bool{
			REG_R8:  true,
			REG_R8B: true,
			REG_R8D: true,
		},
	},

	REG_YMM9: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM9: true,
			REG_ZMM9: true,
		},
	},

	REG_YMM8: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM8: true,
			REG_ZMM8: true,
		},
	},

	REG_R14: {
		Size: 64,
		RelatedRegs: map[string]bool{
			REG_R14B: true,
			REG_R14D: true,
			REG_R14W: true,
		},
	},

	REG_R15: {
		Size: 64,
		RelatedRegs: map[string]bool{
			REG_R15B: true,
			REG_R15D: true,
			REG_R15W: true,
		},
	},

	REG_R12: {
		Size: 64,
		RelatedRegs: map[string]bool{
			REG_R12B: true,
			REG_R12D: true,
			REG_R12W: true,
		},
	},

	REG_R13: {
		Size: 64,
		RelatedRegs: map[string]bool{
			REG_R13B: true,
			REG_R13D: true,
			REG_R13W: true,
		},
	},

	REG_R10: {
		Size: 64,
		RelatedRegs: map[string]bool{
			REG_R10B: true,
			REG_R10D: true,
			REG_R10W: true,
		},
	},

	REG_R11: {
		Size: 64,
		RelatedRegs: map[string]bool{
			REG_R11B: true,
			REG_R11D: true,
			REG_R11W: true,
		},
	},

	REG_YMM1: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM1: true,
			REG_ZMM1: true,
		},
	},

	REG_YMM0: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM0: true,
			REG_ZMM0: true,
		},
	},

	REG_YMM3: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM3: true,
			REG_ZMM3: true,
		},
	},

	REG_YMM2: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM2: true,
			REG_ZMM2: true,
		},
	},

	REG_YMM5: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM5: true,
			REG_ZMM5: true,
		},
	},

	REG_YMM4: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM4: true,
			REG_ZMM4: true,
		},
	},

	REG_YMM7: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM7: true,
			REG_ZMM7: true,
		},
	},

	REG_YMM6: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM6: true,
			REG_ZMM6: true,
		},
	},

	REG_ZMM11: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM11: true,
			REG_XMM11: true,
		},
	},

	REG_DX: {
		Size: 16,
		RelatedRegs: map[string]bool{
			REG_DH:  true,
			REG_DL:  true,
			REG_EDX: true,
			REG_RDX: true,
		},
	},

	REG_YMM32: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM32: true,
			REG_ZMM32: true,
		},
	},

	REG_YMM31: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM31: true,
			REG_ZMM31: true,
		},
	},

	REG_YMM30: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM30: true,
			REG_ZMM30: true,
		},
	},

	REG_DIL: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_EDI: true,
			REG_RDI: true,
			REG_DI:  true,
		},
	},

	REG_R10B: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_R10:  true,
			REG_R10D: true,
			REG_R10W: true,
		},
	},

	REG_RBX: {
		Size: 64,
		RelatedRegs: map[string]bool{
			REG_BH:  true,
			REG_BL:  true,
			REG_EBX: true,
			REG_BX:  true,
		},
	},

	REG_BPL: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_BP:  true,
			REG_EBP: true,
			REG_RBP: true,
		},
	},

	REG_R10D: {
		Size: 32,
		RelatedRegs: map[string]bool{
			REG_R10:  true,
			REG_R10B: true,
			REG_R10W: true,
		},
	},

	REG_XMM10: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM10: true,
			REG_ZMM10: true,
		},
	},

	REG_XMM11: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM11: true,
			REG_ZMM11: true,
		},
	},

	REG_XMM12: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM12: true,
			REG_ZMM12: true,
		},
	},

	REG_XMM13: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM13: true,
			REG_ZMM13: true,
		},
	},

	REG_XMM14: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM14: true,
			REG_ZMM14: true,
		},
	},

	REG_XMM15: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM15: true,
			REG_ZMM15: true,
		},
	},

	REG_XMM16: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM16: true,
			REG_ZMM16: true,
		},
	},

	REG_XMM17: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM17: true,
			REG_ZMM17: true,
		},
	},

	REG_XMM18: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM18: true,
			REG_ZMM18: true,
		},
	},

	REG_XMM19: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM19: true,
			REG_ZMM19: true,
		},
	},

	REG_SIL: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_SI:  true,
			REG_RSI: true,
			REG_ESI: true,
		},
	},

	REG_R10W: {
		Size: 16,
		RelatedRegs: map[string]bool{
			REG_R10:  true,
			REG_R10B: true,
			REG_R10D: true,
		},
	},

	REG_ZMM17: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM17: true,
			REG_XMM17: true,
		},
	},

	REG_MM5: {
		Size:        64,
		RelatedRegs: map[string]bool{},
	},

	REG_MM4: {
		Size:        64,
		RelatedRegs: map[string]bool{},
	},

	REG_MM7: {
		Size:        64,
		RelatedRegs: map[string]bool{},
	},

	REG_MM6: {
		Size:        64,
		RelatedRegs: map[string]bool{},
	},

	REG_MM1: {
		Size:        64,
		RelatedRegs: map[string]bool{},
	},

	REG_MM0: {
		Size:        64,
		RelatedRegs: map[string]bool{},
	},

	REG_MM3: {
		Size:        64,
		RelatedRegs: map[string]bool{},
	},

	REG_MM2: {
		Size:        64,
		RelatedRegs: map[string]bool{},
	},

	REG_YMM28: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM28: true,
			REG_ZMM28: true,
		},
	},

	REG_YMM29: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM29: true,
			REG_ZMM29: true,
		},
	},

	REG_EBP: {
		Size: 32,
		RelatedRegs: map[string]bool{
			REG_BP:  true,
			REG_BPL: true,
			REG_RBP: true,
		},
	},

	REG_YMM20: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM20: true,
			REG_ZMM20: true,
		},
	},

	REG_YMM21: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM21: true,
			REG_ZMM21: true,
		},
	},

	REG_YMM22: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM22: true,
			REG_ZMM22: true,
		},
	},

	REG_YMM23: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM23: true,
			REG_ZMM23: true,
		},
	},

	REG_YMM24: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM24: true,
			REG_ZMM24: true,
		},
	},

	REG_YMM25: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM25: true,
			REG_ZMM25: true,
		},
	},

	REG_YMM26: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM26: true,
			REG_ZMM26: true,
		},
	},

	REG_YMM27: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM27: true,
			REG_ZMM27: true,
		},
	},

	REG_R15D: {
		Size: 32,
		RelatedRegs: map[string]bool{
			REG_R15:  true,
			REG_R15B: true,
			REG_R15W: true,
		},
	},

	REG_R15B: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_R15:  true,
			REG_R15D: true,
			REG_R15W: true,
		},
	},

	REG_ESP: {
		Size: 32,
		RelatedRegs: map[string]bool{
			REG_SPL: true,
			REG_SP:  true,
			REG_RSP: true,
		},
	},

	REG_R15W: {
		Size: 16,
		RelatedRegs: map[string]bool{
			REG_R15:  true,
			REG_R15B: true,
			REG_R15D: true,
		},
	},

	REG_ESI: {
		Size: 32,
		RelatedRegs: map[string]bool{
			REG_SI:  true,
			REG_RSI: true,
			REG_SIL: true,
		},
	},

	REG_BL: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_BX:  true,
			REG_EBX: true,
			REG_RBX: true,
		},
	},

	REG_BH: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_BX:  true,
			REG_EBX: true,
			REG_RBX: true,
		},
	},

	REG_XMM2: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM2: true,
			REG_ZMM2: true,
		},
	},

	REG_XMM3: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM3: true,
			REG_ZMM3: true,
		},
	},

	REG_XMM0: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM0: true,
			REG_ZMM0: true,
		},
	},

	REG_XMM1: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM1: true,
			REG_ZMM1: true,
		},
	},

	REG_XMM6: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM6: true,
			REG_ZMM6: true,
		},
	},

	REG_XMM7: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM7: true,
			REG_ZMM7: true,
		},
	},

	REG_XMM4: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM4: true,
			REG_ZMM4: true,
		},
	},

	REG_XMM5: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM5: true,
			REG_ZMM5: true,
		},
	},

	REG_XMM8: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM8: true,
			REG_ZMM8: true,
		},
	},

	REG_XMM9: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM9: true,
			REG_ZMM9: true,
		},
	},

	REG_BX: {
		Size: 16,
		RelatedRegs: map[string]bool{
			REG_BH:  true,
			REG_BL:  true,
			REG_EBX: true,
			REG_RBX: true,
		},
	},

	REG_ECX: {
		Size: 32,
		RelatedRegs: map[string]bool{
			REG_CH:  true,
			REG_CL:  true,
			REG_CX:  true,
			REG_RCX: true,
		},
	},

	REG_DL: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_DX:  true,
			REG_EDX: true,
			REG_RDX: true,
		},
	},

	REG_R12W: {
		Size: 16,
		RelatedRegs: map[string]bool{
			REG_R12:  true,
			REG_R12B: true,
			REG_R12D: true,
		},
	},

	REG_R9D: {
		Size: 32,
		RelatedRegs: map[string]bool{
			REG_R9:  true,
			REG_R9B: true,
			REG_R9W: true,
		},
	},

	REG_R9B: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_R9:  true,
			REG_R9D: true,
			REG_R9W: true,
		},
	},

	REG_R9: {
		Size: 64,
		RelatedRegs: map[string]bool{
			REG_R9B: true,
			REG_R9D: true,
			REG_R9W: true,
		},
	},

	REG_R12B: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_R12:  true,
			REG_R12D: true,
			REG_R12W: true,
		},
	},

	REG_R12D: {
		Size: 32,
		RelatedRegs: map[string]bool{
			REG_R12:  true,
			REG_R12B: true,
			REG_R12W: true,
		},
	},

	REG_R9W: {
		Size: 16,
		RelatedRegs: map[string]bool{
			REG_R9:  true,
			REG_R9B: true,
			REG_R9D: true,
		},
	},

	REG_EBX: {
		Size: 32,
		RelatedRegs: map[string]bool{
			REG_BH:  true,
			REG_BL:  true,
			REG_BX:  true,
			REG_RBX: true,
		},
	},

	REG_RDI: {
		Size: 64,
		RelatedRegs: map[string]bool{
			REG_DIL: true,
			REG_EDI: true,
			REG_DI:  true,
		},
	},

	REG_CH: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_CX:  true,
			REG_ECX: true,
			REG_RCX: true,
		},
	},

	REG_CL: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_CX:  true,
			REG_ECX: true,
			REG_RCX: true,
		},
	},

	REG_CX: {
		Size: 16,
		RelatedRegs: map[string]bool{
			REG_CH:  true,
			REG_CL:  true,
			REG_ECX: true,
			REG_RCX: true,
		},
	},

	REG_CS: {
		Size:        16,
		RelatedRegs: map[string]bool{},
	},

	REG_RCX: {
		Size: 64,
		RelatedRegs: map[string]bool{
			REG_CH:  true,
			REG_CL:  true,
			REG_ECX: true,
			REG_CX:  true,
		},
	},

	REG_AH: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_AX:  true,
			REG_EAX: true,
			REG_RAX: true,
		},
	},

	REG_XMM29: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM29: true,
			REG_ZMM29: true,
		},
	},

	REG_XMM28: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM28: true,
			REG_ZMM28: true,
		},
	},

	REG_RSI: {
		Size: 64,
		RelatedRegs: map[string]bool{
			REG_SI:  true,
			REG_ESI: true,
			REG_SIL: true,
		},
	},

	REG_XMM21: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM21: true,
			REG_ZMM21: true,
		},
	},

	REG_XMM20: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM20: true,
			REG_ZMM20: true,
		},
	},

	REG_XMM23: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM23: true,
			REG_ZMM23: true,
		},
	},

	REG_XMM22: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM22: true,
			REG_ZMM22: true,
		},
	},

	REG_XMM25: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM25: true,
			REG_ZMM25: true,
		},
	},

	REG_XMM24: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM24: true,
			REG_ZMM24: true,
		},
	},

	REG_XMM27: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM27: true,
			REG_ZMM27: true,
		},
	},

	REG_XMM26: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM26: true,
			REG_ZMM26: true,
		},
	},

	REG_ZMM8: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM8: true,
			REG_XMM8: true,
		},
	},

	REG_ZMM9: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM9: true,
			REG_XMM9: true,
		},
	},

	REG_ZMM0: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM0: true,
			REG_XMM0: true,
		},
	},

	REG_ZMM1: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM1: true,
			REG_XMM1: true,
		},
	},

	REG_ZMM2: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM2: true,
			REG_XMM2: true,
		},
	},

	REG_ZMM3: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM3: true,
			REG_XMM3: true,
		},
	},

	REG_ZMM4: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM4: true,
			REG_XMM4: true,
		},
	},

	REG_ZMM5: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM5: true,
			REG_XMM5: true,
		},
	},

	REG_ZMM6: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM6: true,
			REG_XMM6: true,
		},
	},

	REG_ZMM7: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM7: true,
			REG_XMM7: true,
		},
	},

	REG_ZMM12: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM12: true,
			REG_XMM12: true,
		},
	},

	REG_ZMM13: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM13: true,
			REG_XMM13: true,
		},
	},

	REG_ZMM10: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM10: true,
			REG_XMM10: true,
		},
	},

	REG_R14D: {
		Size: 32,
		RelatedRegs: map[string]bool{
			REG_R14:  true,
			REG_R14B: true,
			REG_R14W: true,
		},
	},

	REG_ZMM16: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM16: true,
			REG_XMM16: true,
		},
	},

	REG_R14B: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_R14:  true,
			REG_R14D: true,
			REG_R14W: true,
		},
	},

	REG_ZMM14: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM14: true,
			REG_XMM14: true,
		},
	},

	REG_ZMM15: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM15: true,
			REG_XMM15: true,
		},
	},

	REG_ZMM18: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM18: true,
			REG_XMM18: true,
		},
	},

	REG_ZMM19: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM19: true,
			REG_XMM19: true,
		},
	},

	REG_RBP: {
		Size: 64,
		RelatedRegs: map[string]bool{
			REG_BP:  true,
			REG_EBP: true,
			REG_BPL: true,
		},
	},

	REG_R14W: {
		Size: 16,
		RelatedRegs: map[string]bool{
			REG_R14:  true,
			REG_R14B: true,
			REG_R14D: true,
		},
	},

	REG_SS: {
		Size:        16,
		RelatedRegs: map[string]bool{},
	},

	REG_SPL: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_SP:  true,
			REG_RSP: true,
			REG_ESP: true,
		},
	},

	REG_DI: {
		Size: 16,
		RelatedRegs: map[string]bool{
			REG_DIL: true,
			REG_EDI: true,
			REG_RDI: true,
		},
	},

	REG_BND0: {
		Size:        128,
		RelatedRegs: map[string]bool{},
	},

	REG_BND1: {
		Size:        128,
		RelatedRegs: map[string]bool{},
	},

	REG_BND2: {
		Size:        128,
		RelatedRegs: map[string]bool{},
	},

	REG_BND3: {
		Size:        128,
		RelatedRegs: map[string]bool{},
	},

	REG_XMM32: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM32: true,
			REG_ZMM32: true,
		},
	},

	REG_R8: {
		Size: 64,
		RelatedRegs: map[string]bool{
			REG_R8B: true,
			REG_R8D: true,
			REG_R8W: true,
		},
	},

	REG_XMM30: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM30: true,
			REG_ZMM30: true,
		},
	},

	REG_XMM31: {
		Size: 128,
		RelatedRegs: map[string]bool{
			REG_YMM31: true,
			REG_ZMM31: true,
		},
	},

	REG_AL: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_AX:  true,
			REG_EAX: true,
			REG_RAX: true,
		},
	},

	REG_RDX: {
		Size: 64,
		RelatedRegs: map[string]bool{
			REG_DH:  true,
			REG_DL:  true,
			REG_EDX: true,
			REG_DX:  true,
		},
	},

	REG_BP: {
		Size: 16,
		RelatedRegs: map[string]bool{
			REG_EBP: true,
			REG_BPL: true,
			REG_RBP: true,
		},
	},

	REG_AX: {
		Size: 16,
		RelatedRegs: map[string]bool{
			REG_AH:  true,
			REG_AL:  true,
			REG_EAX: true,
			REG_RAX: true,
		},
	},

	REG_RAX: {
		Size: 64,
		RelatedRegs: map[string]bool{
			REG_AH:  true,
			REG_AL:  true,
			REG_EAX: true,
			REG_AX:  true,
		},
	},

	REG_DH: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_DX:  true,
			REG_EDX: true,
			REG_RDX: true,
		},
	},

	REG_ZMM29: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM29: true,
			REG_XMM29: true,
		},
	},

	REG_ZMM28: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM28: true,
			REG_XMM28: true,
		},
	},

	REG_ZMM27: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM27: true,
			REG_XMM27: true,
		},
	},

	REG_ZMM26: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM26: true,
			REG_XMM26: true,
		},
	},

	REG_R11B: {
		Size: 8,
		RelatedRegs: map[string]bool{
			REG_R11:  true,
			REG_R11D: true,
			REG_R11W: true,
		},
	},

	REG_ZMM24: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM24: true,
			REG_XMM24: true,
		},
	},

	REG_R11D: {
		Size: 32,
		RelatedRegs: map[string]bool{
			REG_R11:  true,
			REG_R11B: true,
			REG_R11W: true,
		},
	},

	REG_ZMM22: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM22: true,
			REG_XMM22: true,
		},
	},

	REG_ZMM21: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM21: true,
			REG_XMM21: true,
		},
	},

	REG_ZMM20: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM20: true,
			REG_XMM20: true,
		},
	},

	REG_ES: {
		Size:        16,
		RelatedRegs: map[string]bool{},
	},

	REG_R11W: {
		Size: 16,
		RelatedRegs: map[string]bool{
			REG_R11:  true,
			REG_R11B: true,
			REG_R11D: true,
		},
	},

	REG_FS: {
		Size:        16,
		RelatedRegs: map[string]bool{},
	},

	REG_YMM11: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM11: true,
			REG_ZMM11: true,
		},
	},

	REG_YMM10: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM10: true,
			REG_ZMM10: true,
		},
	},

	REG_YMM13: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM13: true,
			REG_ZMM13: true,
		},
	},

	REG_YMM12: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM12: true,
			REG_ZMM12: true,
		},
	},

	REG_YMM15: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM15: true,
			REG_ZMM15: true,
		},
	},

	REG_YMM14: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM14: true,
			REG_ZMM14: true,
		},
	},

	REG_YMM17: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM17: true,
			REG_ZMM17: true,
		},
	},

	REG_YMM16: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM16: true,
			REG_ZMM16: true,
		},
	},

	REG_YMM19: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM19: true,
			REG_ZMM19: true,
		},
	},

	REG_YMM18: {
		Size: 256,
		RelatedRegs: map[string]bool{
			REG_XMM18: true,
			REG_ZMM18: true,
		},
	},

	REG_EAX: {
		Size: 32,
		RelatedRegs: map[string]bool{
			REG_AH:  true,
			REG_AL:  true,
			REG_AX:  true,
			REG_RAX: true,
		},
	},

	REG_ZMM30: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM30: true,
			REG_XMM30: true,
		},
	},

	REG_ZMM31: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM31: true,
			REG_XMM31: true,
		},
	},

	REG_ZMM32: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM32: true,
			REG_XMM32: true,
		},
	},

	REG_SP: {
		Size: 16,
		RelatedRegs: map[string]bool{
			REG_SPL: true,
			REG_RSP: true,
			REG_ESP: true,
		},
	},

	REG_SI: {
		Size: 16,
		RelatedRegs: map[string]bool{
			REG_RSI: true,
			REG_ESI: true,
			REG_SIL: true,
		},
	},

	REG_ZMM25: {
		Size: 512,
		RelatedRegs: map[string]bool{
			REG_YMM25: true,
			REG_XMM25: true,
		},
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
