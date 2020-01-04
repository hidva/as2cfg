
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

%token <u64> HEXNUMBER 
%token <str> LOWERCASESTRING

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
        $$ = append($$, $1)
        alexer := yylex.(*lexer)
        alexer.insts = $$
    }
|   instblock AddressedInst
    {
        $$ = append($1, $2)
        alexer := yylex.(*lexer)
        alexer.insts = $$
    }

AddressedInst:
    HEXNUMBER inst
    {
        $$.Addr = $1
        $$.Inst = $2
    }

/* The INTEL MANUAL says: There may be from zero to three operands, depending on the opcode. */
inst:
    LOWERCASESTRING '\n'
    {
        $$ = *newInstruction($1)
    }    
|   LOWERCASESTRING operand '\n'
    {
        $$ = *newInstruction($1, $2)
    }    
|   LOWERCASESTRING operand ',' operand '\n'
    {
        $$ = *newInstruction($1, $2, $4)
    }    
|   LOWERCASESTRING operand ','  operand ',' operand '\n'
    {
        $$ = *newInstruction($1, $2, $4, $6)        
    }    

operand:
    LOWERCASESTRING
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
|   operand_size_repr LOWERCASESTRING ':' memoffset  /* with a segment selector */
    {
        $$.Kind = MEMORY_OPERAND
        $$.Imme = $1   
        $$.Reg = $2
        $$.Offset = $4
    }
    /* some instructions like lea etc hasn't operand_size_repr. */
|   LOWERCASESTRING ':' onlydispoffset 
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
|    LOWERCASESTRING ':' otheroffset
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
    '[' LOWERCASESTRING '+' HEXNUMBER ']'
    {
        $$.Disp = int64($4)
        $$.Base = $2
    }
|   '[' LOWERCASESTRING '-' HEXNUMBER ']'
    {
        $$.Disp = -int64($4)
        $$.Base = $2
    }
|   '[' LOWERCASESTRING ']'
    {
        $$.Base = $2
    }
|   '[' LOWERCASESTRING '+' LOWERCASESTRING '*' '1' ']'
    {
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 1
    }
|   '[' LOWERCASESTRING '+' LOWERCASESTRING '*' '2' ']'
    {
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 2
    }
|   '[' LOWERCASESTRING '+' LOWERCASESTRING '*' '4' ']'
    {
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 4
    }
|   '[' LOWERCASESTRING '+' LOWERCASESTRING '*' '8' ']'
    {
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 8
    }
|   '[' LOWERCASESTRING '+' LOWERCASESTRING '*' '1' '+' HEXNUMBER ']'
    {
        $$.Disp = int64($8)
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 1
    }
|   '[' LOWERCASESTRING '+' LOWERCASESTRING '*' '2' '+' HEXNUMBER ']'
    {
        $$.Disp = int64($8)
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 2
    }
|   '[' LOWERCASESTRING '+' LOWERCASESTRING '*' '4' '+' HEXNUMBER ']'
    {
        $$.Disp = int64($8)
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 4
    }
|   '[' LOWERCASESTRING '+' LOWERCASESTRING '*' '8' '+' HEXNUMBER ']'
    {
        $$.Disp = int64($8)
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 8
    }
|   '[' LOWERCASESTRING '+' LOWERCASESTRING '*' '1' '-' HEXNUMBER ']'
    {
        $$.Disp = -int64($8)
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 1
    }
|   '[' LOWERCASESTRING '+' LOWERCASESTRING '*' '2' '-' HEXNUMBER ']'
    {
        $$.Disp = -int64($8)
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 2
    }
|   '[' LOWERCASESTRING '+' LOWERCASESTRING '*' '4' '-' HEXNUMBER ']'
    {
        $$.Disp = -int64($8)
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 4
    }
|   '[' LOWERCASESTRING '+' LOWERCASESTRING '*' '8' '-' HEXNUMBER ']'
    {
        $$.Disp = -int64($8)
        $$.Base = $2
        $$.Index = $4
        $$.Scale = 8
    }
|   '[' LOWERCASESTRING '*' '1' '+' HEXNUMBER ']'
    {
        $$.Disp = int64($6)
        $$.Index = $2
        $$.Scale = 1
    }
|   '[' LOWERCASESTRING '*' '2' '+' HEXNUMBER ']'
    {
        $$.Disp = int64($6)
        $$.Index = $2
        $$.Scale = 2
    }
|   '[' LOWERCASESTRING '*' '4' '+' HEXNUMBER ']'
    {
        $$.Disp = int64($6)
        $$.Index = $2
        $$.Scale = 4
    }
|   '[' LOWERCASESTRING '*' '8' '+' HEXNUMBER ']'
    {
        $$.Disp = int64($6)
        $$.Index = $2
        $$.Scale = 8
    }
|   '[' LOWERCASESTRING '*' '1' '-' HEXNUMBER ']'
    {
        $$.Disp = -int64($6)
        $$.Index = $2
        $$.Scale = 1
    }
|   '[' LOWERCASESTRING '*' '2' '-' HEXNUMBER ']'
    {
        $$.Disp = -int64($6)
        $$.Index = $2
        $$.Scale = 2
    }
|   '[' LOWERCASESTRING '*' '4' '-' HEXNUMBER ']'
    {
        $$.Disp = -int64($6)
        $$.Index = $2
        $$.Scale = 4
    }
|   '[' LOWERCASESTRING '*' '8' '-' HEXNUMBER ']'
    {
        $$.Disp = -int64($6)
        $$.Index = $2
        $$.Scale = 8
    }
|   '[' LOWERCASESTRING '*' '1' ']'
    {
        $$.Index = $2
        $$.Scale = 1
    }
|   '[' LOWERCASESTRING '*' '2' ']'
    {
        $$.Index = $2
        $$.Scale = 2
    }
|   '[' LOWERCASESTRING '*' '4' ']'
    {
        $$.Index = $2
        $$.Scale = 4
    }
|   '[' LOWERCASESTRING '*' '8' ']'
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

%%

