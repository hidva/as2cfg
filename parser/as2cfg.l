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

\<.*\>\:?  /*  ignore, gdb disassemble will put some comment information in `<>` */
[ \t\r]+  /* ignore */

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

[a-z]+  {
    /* LOWERCASESTRING */
    lval.str = string(this.buf)
    return LOWERCASESTRING
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

%%
	panic(fmt.Errorf("lex error; buf: %s; current: %d;", this.buf, this.current))
    return 0;
}