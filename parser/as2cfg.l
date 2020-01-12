%{
package parser

import (
    "io"
    "fmt"
    "strconv"
)

type lexer struct {
    source io.Reader
    /*
     * The golex used by us can not construct yytext when we reached an accepting state, 
     * so we construct it by ourself. `buf` plays same role as yytext, the invocation of 
     * next() means that `current` was accepted, so we push it to `buf`, when we enter 
     * execution of action, yytext is saved in `buf`.
     */
    buf []byte /* used to save token */
    current byte

    insts []AddressedInst
}

func newLexer(src io.Reader) *lexer {
    l := &lexer {source: src}
    l.next()
    return l
}

func (this *lexer) Error(e string) {
    fmt.Printf("error happens in parser: %s\n", e)
    return
}

func (this *lexer) next() byte {
    if this.current != 0 {
        this.buf = append(this.buf, this.current)
    }
    this.current = 0
    var b [1]byte
    if ret, _ := this.source.Read(b[:]); ret > 0 {
        this.current = b[0]
    }
    return this.current
}

func (this *lexer) Lex(lval *yySymType) int {
%}

%yyc this.current
%yyn this.next()

%%
    this.buf = this.buf[:0]

(BYTE\ PTR)      {
    lval.u64 = BYTE_PTR
    return BYTE_PTR
    }
(DWORD\ PTR)     {
    lval.u64 = DWORD_PTR
    return DWORD_PTR
    }
(FWORD\ PTR)     {
    lval.u64 = FWORD_PTR
    return FWORD_PTR
    }
(OWORD\ PTR)     {
    lval.u64 = OWORD_PTR
    return OWORD_PTR
    }
(QWORD\ PTR)     {
    lval.u64 = QWORD_PTR
    return QWORD_PTR
    }
(TBYTE\ PTR)     {
    lval.u64 = TBYTE_PTR
    return TBYTE_PTR
    }
(WORD\ PTR)      {
    lval.u64 = WORD_PTR
    return WORD_PTR
    }
(XMMWORD\ PTR)   {
    lval.u64 = XMMWORD_PTR
    return XMMWORD_PTR
    }
(YMMWORD\ PTR)   {
    lval.u64 = YMMWORD_PTR
    return YMMWORD_PTR
    }
(ZMMWORD\ PTR)   {
    lval.u64 = ZMMWORD_PTR
    return ZMMWORD_PTR
    }


0[xX][0-9a-f]+ {
    /* HEXNUMBER */
    var err error
    if lval.u64, err = strconv.ParseUint(string(this.buf), 0, 64); err != nil {
        panic(err)
    }
    return HEXNUMBER    
    }

\n  return '\n'
:   return ':'
\+  return '+'
-   return '-'
\*   return '*'
\[  return '['
\]  return ']'
,   return ','
1   return '1'
2   return '2'
4   return '4'
8   return '8'
\0  return 0

gs    {
    lval.str = "gs"
    return TOKEN_REG_GS
    }
zmm23    {
    lval.str = "zmm23"
    return TOKEN_REG_ZMM23
    }
edi    {
    lval.str = "edi"
    return TOKEN_REG_EDI
    }
edx    {
    lval.str = "edx"
    return TOKEN_REG_EDX
    }
r13w    {
    lval.str = "r13w"
    return TOKEN_REG_R13W
    }
rsp    {
    lval.str = "rsp"
    return TOKEN_REG_RSP
    }
r8b    {
    lval.str = "r8b"
    return TOKEN_REG_R8B
    }
r8d    {
    lval.str = "r8d"
    return TOKEN_REG_R8D
    }
ds    {
    lval.str = "ds"
    return TOKEN_REG_DS
    }
r13d    {
    lval.str = "r13d"
    return TOKEN_REG_R13D
    }
r13b    {
    lval.str = "r13b"
    return TOKEN_REG_R13B
    }
r8w    {
    lval.str = "r8w"
    return TOKEN_REG_R8W
    }
ymm9    {
    lval.str = "ymm9"
    return TOKEN_REG_YMM9
    }
ymm8    {
    lval.str = "ymm8"
    return TOKEN_REG_YMM8
    }
r14    {
    lval.str = "r14"
    return TOKEN_REG_R14
    }
r15    {
    lval.str = "r15"
    return TOKEN_REG_R15
    }
r12    {
    lval.str = "r12"
    return TOKEN_REG_R12
    }
r13    {
    lval.str = "r13"
    return TOKEN_REG_R13
    }
r10    {
    lval.str = "r10"
    return TOKEN_REG_R10
    }
r11    {
    lval.str = "r11"
    return TOKEN_REG_R11
    }
ymm1    {
    lval.str = "ymm1"
    return TOKEN_REG_YMM1
    }
ymm0    {
    lval.str = "ymm0"
    return TOKEN_REG_YMM0
    }
ymm3    {
    lval.str = "ymm3"
    return TOKEN_REG_YMM3
    }
ymm2    {
    lval.str = "ymm2"
    return TOKEN_REG_YMM2
    }
ymm5    {
    lval.str = "ymm5"
    return TOKEN_REG_YMM5
    }
ymm4    {
    lval.str = "ymm4"
    return TOKEN_REG_YMM4
    }
ymm7    {
    lval.str = "ymm7"
    return TOKEN_REG_YMM7
    }
ymm6    {
    lval.str = "ymm6"
    return TOKEN_REG_YMM6
    }
zmm11    {
    lval.str = "zmm11"
    return TOKEN_REG_ZMM11
    }
dx    {
    lval.str = "dx"
    return TOKEN_REG_DX
    }
ymm32    {
    lval.str = "ymm32"
    return TOKEN_REG_YMM32
    }
ymm31    {
    lval.str = "ymm31"
    return TOKEN_REG_YMM31
    }
ymm30    {
    lval.str = "ymm30"
    return TOKEN_REG_YMM30
    }
dil    {
    lval.str = "dil"
    return TOKEN_REG_DIL
    }
r10b    {
    lval.str = "r10b"
    return TOKEN_REG_R10B
    }
rbx    {
    lval.str = "rbx"
    return TOKEN_REG_RBX
    }
bpl    {
    lval.str = "bpl"
    return TOKEN_REG_BPL
    }
r10d    {
    lval.str = "r10d"
    return TOKEN_REG_R10D
    }
xmm10    {
    lval.str = "xmm10"
    return TOKEN_REG_XMM10
    }
xmm11    {
    lval.str = "xmm11"
    return TOKEN_REG_XMM11
    }
xmm12    {
    lval.str = "xmm12"
    return TOKEN_REG_XMM12
    }
xmm13    {
    lval.str = "xmm13"
    return TOKEN_REG_XMM13
    }
xmm14    {
    lval.str = "xmm14"
    return TOKEN_REG_XMM14
    }
xmm15    {
    lval.str = "xmm15"
    return TOKEN_REG_XMM15
    }
xmm16    {
    lval.str = "xmm16"
    return TOKEN_REG_XMM16
    }
xmm17    {
    lval.str = "xmm17"
    return TOKEN_REG_XMM17
    }
xmm18    {
    lval.str = "xmm18"
    return TOKEN_REG_XMM18
    }
xmm19    {
    lval.str = "xmm19"
    return TOKEN_REG_XMM19
    }
sil    {
    lval.str = "sil"
    return TOKEN_REG_SIL
    }
r10w    {
    lval.str = "r10w"
    return TOKEN_REG_R10W
    }
zmm17    {
    lval.str = "zmm17"
    return TOKEN_REG_ZMM17
    }
mm5    {
    lval.str = "mm5"
    return TOKEN_REG_MM5
    }
mm4    {
    lval.str = "mm4"
    return TOKEN_REG_MM4
    }
mm7    {
    lval.str = "mm7"
    return TOKEN_REG_MM7
    }
mm6    {
    lval.str = "mm6"
    return TOKEN_REG_MM6
    }
mm1    {
    lval.str = "mm1"
    return TOKEN_REG_MM1
    }
mm0    {
    lval.str = "mm0"
    return TOKEN_REG_MM0
    }
mm3    {
    lval.str = "mm3"
    return TOKEN_REG_MM3
    }
mm2    {
    lval.str = "mm2"
    return TOKEN_REG_MM2
    }
ymm28    {
    lval.str = "ymm28"
    return TOKEN_REG_YMM28
    }
ymm29    {
    lval.str = "ymm29"
    return TOKEN_REG_YMM29
    }
ebp    {
    lval.str = "ebp"
    return TOKEN_REG_EBP
    }
ymm20    {
    lval.str = "ymm20"
    return TOKEN_REG_YMM20
    }
ymm21    {
    lval.str = "ymm21"
    return TOKEN_REG_YMM21
    }
ymm22    {
    lval.str = "ymm22"
    return TOKEN_REG_YMM22
    }
ymm23    {
    lval.str = "ymm23"
    return TOKEN_REG_YMM23
    }
ymm24    {
    lval.str = "ymm24"
    return TOKEN_REG_YMM24
    }
ymm25    {
    lval.str = "ymm25"
    return TOKEN_REG_YMM25
    }
ymm26    {
    lval.str = "ymm26"
    return TOKEN_REG_YMM26
    }
ymm27    {
    lval.str = "ymm27"
    return TOKEN_REG_YMM27
    }
r15d    {
    lval.str = "r15d"
    return TOKEN_REG_R15D
    }
r15b    {
    lval.str = "r15b"
    return TOKEN_REG_R15B
    }
esp    {
    lval.str = "esp"
    return TOKEN_REG_ESP
    }
r15w    {
    lval.str = "r15w"
    return TOKEN_REG_R15W
    }
esi    {
    lval.str = "esi"
    return TOKEN_REG_ESI
    }
bl    {
    lval.str = "bl"
    return TOKEN_REG_BL
    }
bh    {
    lval.str = "bh"
    return TOKEN_REG_BH
    }
xmm2    {
    lval.str = "xmm2"
    return TOKEN_REG_XMM2
    }
xmm3    {
    lval.str = "xmm3"
    return TOKEN_REG_XMM3
    }
xmm0    {
    lval.str = "xmm0"
    return TOKEN_REG_XMM0
    }
xmm1    {
    lval.str = "xmm1"
    return TOKEN_REG_XMM1
    }
xmm6    {
    lval.str = "xmm6"
    return TOKEN_REG_XMM6
    }
xmm7    {
    lval.str = "xmm7"
    return TOKEN_REG_XMM7
    }
xmm4    {
    lval.str = "xmm4"
    return TOKEN_REG_XMM4
    }
xmm5    {
    lval.str = "xmm5"
    return TOKEN_REG_XMM5
    }
xmm8    {
    lval.str = "xmm8"
    return TOKEN_REG_XMM8
    }
xmm9    {
    lval.str = "xmm9"
    return TOKEN_REG_XMM9
    }
bx    {
    lval.str = "bx"
    return TOKEN_REG_BX
    }
ecx    {
    lval.str = "ecx"
    return TOKEN_REG_ECX
    }
dl    {
    lval.str = "dl"
    return TOKEN_REG_DL
    }
r12w    {
    lval.str = "r12w"
    return TOKEN_REG_R12W
    }
r9d    {
    lval.str = "r9d"
    return TOKEN_REG_R9D
    }
r9b    {
    lval.str = "r9b"
    return TOKEN_REG_R9B
    }
r9    {
    lval.str = "r9"
    return TOKEN_REG_R9
    }
r12b    {
    lval.str = "r12b"
    return TOKEN_REG_R12B
    }
r12d    {
    lval.str = "r12d"
    return TOKEN_REG_R12D
    }
r9w    {
    lval.str = "r9w"
    return TOKEN_REG_R9W
    }
ebx    {
    lval.str = "ebx"
    return TOKEN_REG_EBX
    }
rdi    {
    lval.str = "rdi"
    return TOKEN_REG_RDI
    }
ch    {
    lval.str = "ch"
    return TOKEN_REG_CH
    }
cl    {
    lval.str = "cl"
    return TOKEN_REG_CL
    }
cx    {
    lval.str = "cx"
    return TOKEN_REG_CX
    }
cs    {
    lval.str = "cs"
    return TOKEN_REG_CS
    }
rcx    {
    lval.str = "rcx"
    return TOKEN_REG_RCX
    }
ah    {
    lval.str = "ah"
    return TOKEN_REG_AH
    }
xmm29    {
    lval.str = "xmm29"
    return TOKEN_REG_XMM29
    }
xmm28    {
    lval.str = "xmm28"
    return TOKEN_REG_XMM28
    }
rsi    {
    lval.str = "rsi"
    return TOKEN_REG_RSI
    }
xmm21    {
    lval.str = "xmm21"
    return TOKEN_REG_XMM21
    }
xmm20    {
    lval.str = "xmm20"
    return TOKEN_REG_XMM20
    }
xmm23    {
    lval.str = "xmm23"
    return TOKEN_REG_XMM23
    }
xmm22    {
    lval.str = "xmm22"
    return TOKEN_REG_XMM22
    }
xmm25    {
    lval.str = "xmm25"
    return TOKEN_REG_XMM25
    }
xmm24    {
    lval.str = "xmm24"
    return TOKEN_REG_XMM24
    }
xmm27    {
    lval.str = "xmm27"
    return TOKEN_REG_XMM27
    }
xmm26    {
    lval.str = "xmm26"
    return TOKEN_REG_XMM26
    }
zmm8    {
    lval.str = "zmm8"
    return TOKEN_REG_ZMM8
    }
zmm9    {
    lval.str = "zmm9"
    return TOKEN_REG_ZMM9
    }
zmm0    {
    lval.str = "zmm0"
    return TOKEN_REG_ZMM0
    }
zmm1    {
    lval.str = "zmm1"
    return TOKEN_REG_ZMM1
    }
zmm2    {
    lval.str = "zmm2"
    return TOKEN_REG_ZMM2
    }
zmm3    {
    lval.str = "zmm3"
    return TOKEN_REG_ZMM3
    }
zmm4    {
    lval.str = "zmm4"
    return TOKEN_REG_ZMM4
    }
zmm5    {
    lval.str = "zmm5"
    return TOKEN_REG_ZMM5
    }
zmm6    {
    lval.str = "zmm6"
    return TOKEN_REG_ZMM6
    }
zmm7    {
    lval.str = "zmm7"
    return TOKEN_REG_ZMM7
    }
zmm12    {
    lval.str = "zmm12"
    return TOKEN_REG_ZMM12
    }
zmm13    {
    lval.str = "zmm13"
    return TOKEN_REG_ZMM13
    }
zmm10    {
    lval.str = "zmm10"
    return TOKEN_REG_ZMM10
    }
r14d    {
    lval.str = "r14d"
    return TOKEN_REG_R14D
    }
zmm16    {
    lval.str = "zmm16"
    return TOKEN_REG_ZMM16
    }
r14b    {
    lval.str = "r14b"
    return TOKEN_REG_R14B
    }
zmm14    {
    lval.str = "zmm14"
    return TOKEN_REG_ZMM14
    }
zmm15    {
    lval.str = "zmm15"
    return TOKEN_REG_ZMM15
    }
zmm18    {
    lval.str = "zmm18"
    return TOKEN_REG_ZMM18
    }
zmm19    {
    lval.str = "zmm19"
    return TOKEN_REG_ZMM19
    }
rbp    {
    lval.str = "rbp"
    return TOKEN_REG_RBP
    }
r14w    {
    lval.str = "r14w"
    return TOKEN_REG_R14W
    }
ss    {
    lval.str = "ss"
    return TOKEN_REG_SS
    }
spl    {
    lval.str = "spl"
    return TOKEN_REG_SPL
    }
di    {
    lval.str = "di"
    return TOKEN_REG_DI
    }
bnd0    {
    lval.str = "bnd0"
    return TOKEN_REG_BND0
    }
bnd1    {
    lval.str = "bnd1"
    return TOKEN_REG_BND1
    }
bnd2    {
    lval.str = "bnd2"
    return TOKEN_REG_BND2
    }
bnd3    {
    lval.str = "bnd3"
    return TOKEN_REG_BND3
    }
xmm32    {
    lval.str = "xmm32"
    return TOKEN_REG_XMM32
    }
r8    {
    lval.str = "r8"
    return TOKEN_REG_R8
    }
xmm30    {
    lval.str = "xmm30"
    return TOKEN_REG_XMM30
    }
xmm31    {
    lval.str = "xmm31"
    return TOKEN_REG_XMM31
    }
al    {
    lval.str = "al"
    return TOKEN_REG_AL
    }
rdx    {
    lval.str = "rdx"
    return TOKEN_REG_RDX
    }
bp    {
    lval.str = "bp"
    return TOKEN_REG_BP
    }
ax    {
    lval.str = "ax"
    return TOKEN_REG_AX
    }
rax    {
    lval.str = "rax"
    return TOKEN_REG_RAX
    }
dh    {
    lval.str = "dh"
    return TOKEN_REG_DH
    }
zmm29    {
    lval.str = "zmm29"
    return TOKEN_REG_ZMM29
    }
zmm28    {
    lval.str = "zmm28"
    return TOKEN_REG_ZMM28
    }
zmm27    {
    lval.str = "zmm27"
    return TOKEN_REG_ZMM27
    }
zmm26    {
    lval.str = "zmm26"
    return TOKEN_REG_ZMM26
    }
r11b    {
    lval.str = "r11b"
    return TOKEN_REG_R11B
    }
zmm24    {
    lval.str = "zmm24"
    return TOKEN_REG_ZMM24
    }
r11d    {
    lval.str = "r11d"
    return TOKEN_REG_R11D
    }
zmm22    {
    lval.str = "zmm22"
    return TOKEN_REG_ZMM22
    }
zmm21    {
    lval.str = "zmm21"
    return TOKEN_REG_ZMM21
    }
zmm20    {
    lval.str = "zmm20"
    return TOKEN_REG_ZMM20
    }
es    {
    lval.str = "es"
    return TOKEN_REG_ES
    }
r11w    {
    lval.str = "r11w"
    return TOKEN_REG_R11W
    }
fs    {
    lval.str = "fs"
    return TOKEN_REG_FS
    }
ymm11    {
    lval.str = "ymm11"
    return TOKEN_REG_YMM11
    }
ymm10    {
    lval.str = "ymm10"
    return TOKEN_REG_YMM10
    }
ymm13    {
    lval.str = "ymm13"
    return TOKEN_REG_YMM13
    }
ymm12    {
    lval.str = "ymm12"
    return TOKEN_REG_YMM12
    }
ymm15    {
    lval.str = "ymm15"
    return TOKEN_REG_YMM15
    }
ymm14    {
    lval.str = "ymm14"
    return TOKEN_REG_YMM14
    }
ymm17    {
    lval.str = "ymm17"
    return TOKEN_REG_YMM17
    }
ymm16    {
    lval.str = "ymm16"
    return TOKEN_REG_YMM16
    }
ymm19    {
    lval.str = "ymm19"
    return TOKEN_REG_YMM19
    }
ymm18    {
    lval.str = "ymm18"
    return TOKEN_REG_YMM18
    }
eax    {
    lval.str = "eax"
    return TOKEN_REG_EAX
    }
zmm30    {
    lval.str = "zmm30"
    return TOKEN_REG_ZMM30
    }
zmm31    {
    lval.str = "zmm31"
    return TOKEN_REG_ZMM31
    }
zmm32    {
    lval.str = "zmm32"
    return TOKEN_REG_ZMM32
    }
sp    {
    lval.str = "sp"
    return TOKEN_REG_SP
    }
si    {
    lval.str = "si"
    return TOKEN_REG_SI
    }
zmm25    {
    lval.str = "zmm25"
    return TOKEN_REG_ZMM25
    }
eflags    {
    lval.str = "eflags"
    return TOKEN_REG_EFLAGS
    }
rep    {
    lval.str = "rep"
    return TOKEN_INST_PREFIX_REP
    }
repz    {
    lval.str = "repz"
    return TOKEN_INST_PREFIX_REPZ
    }
repe    {
    lval.str = "repe"
    return TOKEN_INST_PREFIX_REPE
    }
repne    {
    lval.str = "repne"
    return TOKEN_INST_PREFIX_REPNE
    }
repnz    {
    lval.str = "repnz"
    return TOKEN_INST_PREFIX_REPNZ
    }
lock    {
    lval.str = "lock"
    return TOKEN_INST_PREFIX_LOCK
    }

#.*$  /* ignore comment.. */
\<.*\>\:?  /*  ignore, gdb disassemble will put some comment information in `<>` */
[ \t\r]+  /* ignore */


[a-z][a-z0-9]+  {
    /* INST_MNEM */
    lval.str = string(this.buf)
    return INST_MNEM
    }

%%
	panic(fmt.Errorf("lex error; buf: %s; current: %d;", this.buf, this.current))
    return 0;
}