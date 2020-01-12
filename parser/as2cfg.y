
%{
package parser
%}

%union {
    u64 uint64
    str string
    offset_type MemOffset
    operand_type Operand
    inst_type Instruction
    addr_inst_type AddressedInst
    inst_block_type []AddressedInst
}

/* we use 0 to represent that the operand size isn't specified. */
%token <u64> BYTE_PTR  30000
%token <u64> DWORD_PTR 30001
%token <u64> FWORD_PTR 30002
%token <u64> OWORD_PTR 30003
%token <u64> QWORD_PTR 30004
%token <u64> TBYTE_PTR 30005
%token <u64> WORD_PTR 30006
%token <u64> XMMWORD_PTR 30007
%token <u64> YMMWORD_PTR 30008
%token <u64> ZMMWORD_PTR 30009
%token <str> TOKEN_REG_GS
%token <str> TOKEN_REG_ZMM23
%token <str> TOKEN_REG_EDI
%token <str> TOKEN_REG_EDX
%token <str> TOKEN_REG_R13W
%token <str> TOKEN_REG_RSP
%token <str> TOKEN_REG_R8B
%token <str> TOKEN_REG_R8D
%token <str> TOKEN_REG_DS
%token <str> TOKEN_REG_R13D
%token <str> TOKEN_REG_R13B
%token <str> TOKEN_REG_R8W
%token <str> TOKEN_REG_YMM9
%token <str> TOKEN_REG_YMM8
%token <str> TOKEN_REG_R14
%token <str> TOKEN_REG_R15
%token <str> TOKEN_REG_R12
%token <str> TOKEN_REG_R13
%token <str> TOKEN_REG_R10
%token <str> TOKEN_REG_R11
%token <str> TOKEN_REG_YMM1
%token <str> TOKEN_REG_YMM0
%token <str> TOKEN_REG_YMM3
%token <str> TOKEN_REG_YMM2
%token <str> TOKEN_REG_YMM5
%token <str> TOKEN_REG_YMM4
%token <str> TOKEN_REG_YMM7
%token <str> TOKEN_REG_YMM6
%token <str> TOKEN_REG_ZMM11
%token <str> TOKEN_REG_DX
%token <str> TOKEN_REG_YMM32
%token <str> TOKEN_REG_YMM31
%token <str> TOKEN_REG_YMM30
%token <str> TOKEN_REG_DIL
%token <str> TOKEN_REG_R10B
%token <str> TOKEN_REG_RBX
%token <str> TOKEN_REG_BPL
%token <str> TOKEN_REG_R10D
%token <str> TOKEN_REG_XMM10
%token <str> TOKEN_REG_XMM11
%token <str> TOKEN_REG_XMM12
%token <str> TOKEN_REG_XMM13
%token <str> TOKEN_REG_XMM14
%token <str> TOKEN_REG_XMM15
%token <str> TOKEN_REG_XMM16
%token <str> TOKEN_REG_XMM17
%token <str> TOKEN_REG_XMM18
%token <str> TOKEN_REG_XMM19
%token <str> TOKEN_REG_SIL
%token <str> TOKEN_REG_R10W
%token <str> TOKEN_REG_ZMM17
%token <str> TOKEN_REG_MM5
%token <str> TOKEN_REG_MM4
%token <str> TOKEN_REG_MM7
%token <str> TOKEN_REG_MM6
%token <str> TOKEN_REG_MM1
%token <str> TOKEN_REG_MM0
%token <str> TOKEN_REG_MM3
%token <str> TOKEN_REG_MM2
%token <str> TOKEN_REG_YMM28
%token <str> TOKEN_REG_YMM29
%token <str> TOKEN_REG_EBP
%token <str> TOKEN_REG_YMM20
%token <str> TOKEN_REG_YMM21
%token <str> TOKEN_REG_YMM22
%token <str> TOKEN_REG_YMM23
%token <str> TOKEN_REG_YMM24
%token <str> TOKEN_REG_YMM25
%token <str> TOKEN_REG_YMM26
%token <str> TOKEN_REG_YMM27
%token <str> TOKEN_REG_R15D
%token <str> TOKEN_REG_R15B
%token <str> TOKEN_REG_ESP
%token <str> TOKEN_REG_R15W
%token <str> TOKEN_REG_ESI
%token <str> TOKEN_REG_BL
%token <str> TOKEN_REG_BH
%token <str> TOKEN_REG_XMM2
%token <str> TOKEN_REG_XMM3
%token <str> TOKEN_REG_XMM0
%token <str> TOKEN_REG_XMM1
%token <str> TOKEN_REG_XMM6
%token <str> TOKEN_REG_XMM7
%token <str> TOKEN_REG_XMM4
%token <str> TOKEN_REG_XMM5
%token <str> TOKEN_REG_XMM8
%token <str> TOKEN_REG_XMM9
%token <str> TOKEN_REG_BX
%token <str> TOKEN_REG_ECX
%token <str> TOKEN_REG_DL
%token <str> TOKEN_REG_R12W
%token <str> TOKEN_REG_R9D
%token <str> TOKEN_REG_R9B
%token <str> TOKEN_REG_R9
%token <str> TOKEN_REG_R12B
%token <str> TOKEN_REG_R12D
%token <str> TOKEN_REG_R9W
%token <str> TOKEN_REG_EBX
%token <str> TOKEN_REG_RDI
%token <str> TOKEN_REG_CH
%token <str> TOKEN_REG_CL
%token <str> TOKEN_REG_CX
%token <str> TOKEN_REG_CS
%token <str> TOKEN_REG_RCX
%token <str> TOKEN_REG_AH
%token <str> TOKEN_REG_XMM29
%token <str> TOKEN_REG_XMM28
%token <str> TOKEN_REG_RSI
%token <str> TOKEN_REG_XMM21
%token <str> TOKEN_REG_XMM20
%token <str> TOKEN_REG_XMM23
%token <str> TOKEN_REG_XMM22
%token <str> TOKEN_REG_XMM25
%token <str> TOKEN_REG_XMM24
%token <str> TOKEN_REG_XMM27
%token <str> TOKEN_REG_XMM26
%token <str> TOKEN_REG_ZMM8
%token <str> TOKEN_REG_ZMM9
%token <str> TOKEN_REG_ZMM0
%token <str> TOKEN_REG_ZMM1
%token <str> TOKEN_REG_ZMM2
%token <str> TOKEN_REG_ZMM3
%token <str> TOKEN_REG_ZMM4
%token <str> TOKEN_REG_ZMM5
%token <str> TOKEN_REG_ZMM6
%token <str> TOKEN_REG_ZMM7
%token <str> TOKEN_REG_ZMM12
%token <str> TOKEN_REG_ZMM13
%token <str> TOKEN_REG_ZMM10
%token <str> TOKEN_REG_R14D
%token <str> TOKEN_REG_ZMM16
%token <str> TOKEN_REG_R14B
%token <str> TOKEN_REG_ZMM14
%token <str> TOKEN_REG_ZMM15
%token <str> TOKEN_REG_ZMM18
%token <str> TOKEN_REG_ZMM19
%token <str> TOKEN_REG_RBP
%token <str> TOKEN_REG_R14W
%token <str> TOKEN_REG_SS
%token <str> TOKEN_REG_SPL
%token <str> TOKEN_REG_DI
%token <str> TOKEN_REG_BND0
%token <str> TOKEN_REG_BND1
%token <str> TOKEN_REG_BND2
%token <str> TOKEN_REG_BND3
%token <str> TOKEN_REG_XMM32
%token <str> TOKEN_REG_R8
%token <str> TOKEN_REG_XMM30
%token <str> TOKEN_REG_XMM31
%token <str> TOKEN_REG_AL
%token <str> TOKEN_REG_RDX
%token <str> TOKEN_REG_BP
%token <str> TOKEN_REG_AX
%token <str> TOKEN_REG_RAX
%token <str> TOKEN_REG_DH
%token <str> TOKEN_REG_ZMM29
%token <str> TOKEN_REG_ZMM28
%token <str> TOKEN_REG_ZMM27
%token <str> TOKEN_REG_ZMM26
%token <str> TOKEN_REG_R11B
%token <str> TOKEN_REG_ZMM24
%token <str> TOKEN_REG_R11D
%token <str> TOKEN_REG_ZMM22
%token <str> TOKEN_REG_ZMM21
%token <str> TOKEN_REG_ZMM20
%token <str> TOKEN_REG_ES
%token <str> TOKEN_REG_R11W
%token <str> TOKEN_REG_FS
%token <str> TOKEN_REG_YMM11
%token <str> TOKEN_REG_YMM10
%token <str> TOKEN_REG_YMM13
%token <str> TOKEN_REG_YMM12
%token <str> TOKEN_REG_YMM15
%token <str> TOKEN_REG_YMM14
%token <str> TOKEN_REG_YMM17
%token <str> TOKEN_REG_YMM16
%token <str> TOKEN_REG_YMM19
%token <str> TOKEN_REG_YMM18
%token <str> TOKEN_REG_EAX
%token <str> TOKEN_REG_ZMM30
%token <str> TOKEN_REG_ZMM31
%token <str> TOKEN_REG_ZMM32
%token <str> TOKEN_REG_SP
%token <str> TOKEN_REG_SI
%token <str> TOKEN_REG_ZMM25
%token <str> TOKEN_REG_EFLAGS

%token <str> TOKEN_INST_PREFIX_LOCK
%token <str> TOKEN_INST_PREFIX_REP
%token <str> TOKEN_INST_PREFIX_REPZ
%token <str> TOKEN_INST_PREFIX_REPE
%token <str> TOKEN_INST_PREFIX_REPNE
%token <str> TOKEN_INST_PREFIX_REPNZ  

%token <u64> HEXNUMBER 
%token <str> INST_MNEM

%type <str> inst_prefix
%type <str> register_operand
%type <u64> operand_size_repr
%type <offset_type> memoffset
%type <offset_type> onlydispoffset
%type <offset_type> otheroffset
%type <operand_type> operand
%type <inst_type> inst
%type <addr_inst_type> AddressedInst
%type <inst_block_type> instblock

%%

instblock:
    AddressedInst 
    {
        if (!g_ignored_insts[$1.Inst.Instmnem]) {
            $$ = append($$, $1)
            alexer := yylex.(*lexer)
            alexer.insts = $$
        }
    }
|   instblock AddressedInst
    {
        if (!g_ignored_insts[$2.Inst.Instmnem]) {
            $$ = append($1, $2)
            alexer := yylex.(*lexer)
            alexer.insts = $$
        }
    }

AddressedInst:
    HEXNUMBER inst
    {
        $$.Addr = $1
        $$.Inst = $2
    }

/* The INTEL MANUAL says: There may be from zero to three operands, depending on the opcode. */
inst:
    INST_MNEM '\n'
    {
        $$ = *newInstruction($1)
    }    
|   INST_MNEM operand '\n'
    {
        $$ = *newInstruction($1, $2)
    }    
|   INST_MNEM operand ',' operand '\n'
    {
        $$ = *newInstruction($1, $2, $4)
    }    
|   INST_MNEM operand ','  operand ',' operand '\n'
    {
        $$ = *newInstruction($1, $2, $4, $6)        
    }    
|   inst_prefix INST_MNEM '\n'
    {
        $$ = *newInstruction($2)
        $$.Instprefix = $1
    }    
|   inst_prefix INST_MNEM operand '\n'
    {
        $$ = *newInstruction($2, $3)
        $$.Instprefix = $1
    }    
|   inst_prefix INST_MNEM operand ',' operand '\n'
    {
        $$ = *newInstruction($2, $3, $5)
        $$.Instprefix = $1
    }    
|   inst_prefix INST_MNEM operand ','  operand ',' operand '\n'
    {
        $$ = *newInstruction($2, $3, $5, $7)        
        $$.Instprefix = $1
    }    

operand:
    register_operand
    {
        $$.Kind = REGISTER_OPERAND
        $$.Reg = $1
    }
|   HEXNUMBER
    {
        $$.Kind = IMMEDIATE_OPERAND
        $$.Imme = $1
    }
|   operand_size_repr memoffset /* a memory location without a segment selector */
    {
        $$.Kind = MEMORY_OPERAND   
        $$.Imme = $1   
        $$.Offset = $2  
    }
|   operand_size_repr register_operand ':' memoffset  /* with a segment selector */
    {
        $$.Kind = MEMORY_OPERAND
        $$.Imme = $1   
        $$.Reg = $2
        $$.Offset = $4
    }
    /* some instructions like lea etc hasn't operand_size_repr. */
|   register_operand ':' onlydispoffset 
    {
        $$.Kind = MEMORY_OPERAND
        $$.Reg = $1
        $$.Offset = $3
    }
|   otheroffset
    {
        $$.Kind = MEMORY_OPERAND
        $$.Offset = $1
    }
|   register_operand ':' otheroffset
    {
        $$.Kind = MEMORY_OPERAND
        $$.Reg = $1
        $$.Offset = $3
    }

onlydispoffset:
    HEXNUMBER
    {
        $$.Disp = int64($1)
    }


otheroffset:
    '[' register_operand '+' HEXNUMBER ']'
    {
        $$.Disp = int64($4)
        $$.Base = $2
    }
|   '[' register_operand '-' HEXNUMBER ']'
    {
        $$.Disp = -int64($4)
        $$.Base = $2
    }
|   '[' register_operand ']'
    {
        $$.Base = $2
    }
|   '[' register_operand '+' register_operand '*' '1' ']'
    {
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 1
    }
|   '[' register_operand '+' register_operand '*' '2' ']'
    {
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 2
    }
|   '[' register_operand '+' register_operand '*' '4' ']'
    {
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 4
    }
|   '[' register_operand '+' register_operand '*' '8' ']'
    {
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 8
    }
|   '[' register_operand '+' register_operand '*' '1' '+' HEXNUMBER ']'
    {
        $$.Disp = int64($8)
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 1
    }
|   '[' register_operand '+' register_operand '*' '2' '+' HEXNUMBER ']'
    {
        $$.Disp = int64($8)
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 2
    }
|   '[' register_operand '+' register_operand '*' '4' '+' HEXNUMBER ']'
    {
        $$.Disp = int64($8)
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 4
    }
|   '[' register_operand '+' register_operand '*' '8' '+' HEXNUMBER ']'
    {
        $$.Disp = int64($8)
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 8
    }
|   '[' register_operand '+' register_operand '*' '1' '-' HEXNUMBER ']'
    {
        $$.Disp = -int64($8)
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 1
    }
|   '[' register_operand '+' register_operand '*' '2' '-' HEXNUMBER ']'
    {
        $$.Disp = -int64($8)
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 2
    }
|   '[' register_operand '+' register_operand '*' '4' '-' HEXNUMBER ']'
    {
        $$.Disp = -int64($8)
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 4
    }
|   '[' register_operand '+' register_operand '*' '8' '-' HEXNUMBER ']'
    {
        $$.Disp = -int64($8)
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 8
    }
|   '[' register_operand '*' '1' '+' HEXNUMBER ']'
    {
        $$.Disp = int64($6)
        $$.Index = $2
        $$.Scale = 1
    }
|   '[' register_operand '*' '2' '+' HEXNUMBER ']'
    {
        $$.Disp = int64($6)
        $$.Index = $2
        $$.Scale = 2
    }
|   '[' register_operand '*' '4' '+' HEXNUMBER ']'
    {
        $$.Disp = int64($6)
        $$.Index = $2
        $$.Scale = 4
    }
|   '[' register_operand '*' '8' '+' HEXNUMBER ']'
    {
        $$.Disp = int64($6)
        $$.Index = $2
        $$.Scale = 8
    }
|   '[' register_operand '*' '1' '-' HEXNUMBER ']'
    {
        $$.Disp = -int64($6)
        $$.Index = $2
        $$.Scale = 1
    }
|   '[' register_operand '*' '2' '-' HEXNUMBER ']'
    {
        $$.Disp = -int64($6)
        $$.Index = $2
        $$.Scale = 2
    }
|   '[' register_operand '*' '4' '-' HEXNUMBER ']'
    {
        $$.Disp = -int64($6)
        $$.Index = $2
        $$.Scale = 4
    }
|   '[' register_operand '*' '8' '-' HEXNUMBER ']'
    {
        $$.Disp = -int64($6)
        $$.Index = $2
        $$.Scale = 8
    }
|   '[' register_operand '*' '1' ']'
    {
        $$.Index = $2
        $$.Scale = 1
    }
|   '[' register_operand '*' '2' ']'
    {
        $$.Index = $2
        $$.Scale = 2
    }
|   '[' register_operand '*' '4' ']'
    {
        $$.Index = $2
        $$.Scale = 4
    }
|   '[' register_operand '*' '8' ']'
    {
        $$.Index = $2
        $$.Scale = 8
    }


memoffset:
    onlydispoffset
|   otheroffset


operand_size_repr: 
    BYTE_PTR 
|   DWORD_PTR
|   FWORD_PTR
|   OWORD_PTR
|   QWORD_PTR
|   TBYTE_PTR
|   WORD_PTR
|   XMMWORD_PTR
|   YMMWORD_PTR
|   ZMMWORD_PTR

register_operand:
    TOKEN_REG_GS
|   TOKEN_REG_ZMM23
|   TOKEN_REG_EDI
|   TOKEN_REG_EDX
|   TOKEN_REG_R13W
|   TOKEN_REG_RSP
|   TOKEN_REG_R8B
|   TOKEN_REG_R8D
|   TOKEN_REG_DS
|   TOKEN_REG_R13D
|   TOKEN_REG_R13B
|   TOKEN_REG_R8W
|   TOKEN_REG_YMM9
|   TOKEN_REG_YMM8
|   TOKEN_REG_R14
|   TOKEN_REG_R15
|   TOKEN_REG_R12
|   TOKEN_REG_R13
|   TOKEN_REG_R10
|   TOKEN_REG_R11
|   TOKEN_REG_YMM1
|   TOKEN_REG_YMM0
|   TOKEN_REG_YMM3
|   TOKEN_REG_YMM2
|   TOKEN_REG_YMM5
|   TOKEN_REG_YMM4
|   TOKEN_REG_YMM7
|   TOKEN_REG_YMM6
|   TOKEN_REG_ZMM11
|   TOKEN_REG_DX
|   TOKEN_REG_YMM32
|   TOKEN_REG_YMM31
|   TOKEN_REG_YMM30
|   TOKEN_REG_DIL
|   TOKEN_REG_R10B
|   TOKEN_REG_RBX
|   TOKEN_REG_BPL
|   TOKEN_REG_R10D
|   TOKEN_REG_XMM10
|   TOKEN_REG_XMM11
|   TOKEN_REG_XMM12
|   TOKEN_REG_XMM13
|   TOKEN_REG_XMM14
|   TOKEN_REG_XMM15
|   TOKEN_REG_XMM16
|   TOKEN_REG_XMM17
|   TOKEN_REG_XMM18
|   TOKEN_REG_XMM19
|   TOKEN_REG_SIL
|   TOKEN_REG_R10W
|   TOKEN_REG_ZMM17
|   TOKEN_REG_MM5
|   TOKEN_REG_MM4
|   TOKEN_REG_MM7
|   TOKEN_REG_MM6
|   TOKEN_REG_MM1
|   TOKEN_REG_MM0
|   TOKEN_REG_MM3
|   TOKEN_REG_MM2
|   TOKEN_REG_YMM28
|   TOKEN_REG_YMM29
|   TOKEN_REG_EBP
|   TOKEN_REG_YMM20
|   TOKEN_REG_YMM21
|   TOKEN_REG_YMM22
|   TOKEN_REG_YMM23
|   TOKEN_REG_YMM24
|   TOKEN_REG_YMM25
|   TOKEN_REG_YMM26
|   TOKEN_REG_YMM27
|   TOKEN_REG_R15D
|   TOKEN_REG_R15B
|   TOKEN_REG_ESP
|   TOKEN_REG_R15W
|   TOKEN_REG_ESI
|   TOKEN_REG_BL
|   TOKEN_REG_BH
|   TOKEN_REG_XMM2
|   TOKEN_REG_XMM3
|   TOKEN_REG_XMM0
|   TOKEN_REG_XMM1
|   TOKEN_REG_XMM6
|   TOKEN_REG_XMM7
|   TOKEN_REG_XMM4
|   TOKEN_REG_XMM5
|   TOKEN_REG_XMM8
|   TOKEN_REG_XMM9
|   TOKEN_REG_BX
|   TOKEN_REG_ECX
|   TOKEN_REG_DL
|   TOKEN_REG_R12W
|   TOKEN_REG_R9D
|   TOKEN_REG_R9B
|   TOKEN_REG_R9
|   TOKEN_REG_R12B
|   TOKEN_REG_R12D
|   TOKEN_REG_R9W
|   TOKEN_REG_EBX
|   TOKEN_REG_RDI
|   TOKEN_REG_CH
|   TOKEN_REG_CL
|   TOKEN_REG_CX
|   TOKEN_REG_CS
|   TOKEN_REG_RCX
|   TOKEN_REG_AH
|   TOKEN_REG_XMM29
|   TOKEN_REG_XMM28
|   TOKEN_REG_RSI
|   TOKEN_REG_XMM21
|   TOKEN_REG_XMM20
|   TOKEN_REG_XMM23
|   TOKEN_REG_XMM22
|   TOKEN_REG_XMM25
|   TOKEN_REG_XMM24
|   TOKEN_REG_XMM27
|   TOKEN_REG_XMM26
|   TOKEN_REG_ZMM8
|   TOKEN_REG_ZMM9
|   TOKEN_REG_ZMM0
|   TOKEN_REG_ZMM1
|   TOKEN_REG_ZMM2
|   TOKEN_REG_ZMM3
|   TOKEN_REG_ZMM4
|   TOKEN_REG_ZMM5
|   TOKEN_REG_ZMM6
|   TOKEN_REG_ZMM7
|   TOKEN_REG_ZMM12
|   TOKEN_REG_ZMM13
|   TOKEN_REG_ZMM10
|   TOKEN_REG_R14D
|   TOKEN_REG_ZMM16
|   TOKEN_REG_R14B
|   TOKEN_REG_ZMM14
|   TOKEN_REG_ZMM15
|   TOKEN_REG_ZMM18
|   TOKEN_REG_ZMM19
|   TOKEN_REG_RBP
|   TOKEN_REG_R14W
|   TOKEN_REG_SS
|   TOKEN_REG_SPL
|   TOKEN_REG_DI
|   TOKEN_REG_BND0
|   TOKEN_REG_BND1
|   TOKEN_REG_BND2
|   TOKEN_REG_BND3
|   TOKEN_REG_XMM32
|   TOKEN_REG_R8
|   TOKEN_REG_XMM30
|   TOKEN_REG_XMM31
|   TOKEN_REG_AL
|   TOKEN_REG_RDX
|   TOKEN_REG_BP
|   TOKEN_REG_AX
|   TOKEN_REG_RAX
|   TOKEN_REG_DH
|   TOKEN_REG_ZMM29
|   TOKEN_REG_ZMM28
|   TOKEN_REG_ZMM27
|   TOKEN_REG_ZMM26
|   TOKEN_REG_R11B
|   TOKEN_REG_ZMM24
|   TOKEN_REG_R11D
|   TOKEN_REG_ZMM22
|   TOKEN_REG_ZMM21
|   TOKEN_REG_ZMM20
|   TOKEN_REG_ES
|   TOKEN_REG_R11W
|   TOKEN_REG_FS
|   TOKEN_REG_YMM11
|   TOKEN_REG_YMM10
|   TOKEN_REG_YMM13
|   TOKEN_REG_YMM12
|   TOKEN_REG_YMM15
|   TOKEN_REG_YMM14
|   TOKEN_REG_YMM17
|   TOKEN_REG_YMM16
|   TOKEN_REG_YMM19
|   TOKEN_REG_YMM18
|   TOKEN_REG_EAX
|   TOKEN_REG_ZMM30
|   TOKEN_REG_ZMM31
|   TOKEN_REG_ZMM32
|   TOKEN_REG_SP
|   TOKEN_REG_SI
|   TOKEN_REG_ZMM25
|   TOKEN_REG_EFLAGS

inst_prefix:
    TOKEN_INST_PREFIX_LOCK
|   TOKEN_INST_PREFIX_REP
|   TOKEN_INST_PREFIX_REPZ
|   TOKEN_INST_PREFIX_REPNZ
|   TOKEN_INST_PREFIX_REPE
|   TOKEN_INST_PREFIX_REPNE

%%

