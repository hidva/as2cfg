// Code generated by golex. DO NOT EDIT.
// em... there are some bugs in golex, so we must apply as2cfg.l.go.patch, and git add this file.

package parser

import (
	"fmt"
	"io"
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
	buf     []byte /* used to save token */
	current byte

	insts []AddressedInst
}

func newLexer(src io.Reader) *lexer {
	l := &lexer{source: src}
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

yystate0:

	this.buf = this.buf[:0]

	goto yystart1

yystate1:
	this.next()
yystart1:
	switch {
	default:
		goto yyabort
	case this.current == '*':
		goto yystate5
	case this.current == '+':
		goto yystate6
	case this.current == ',':
		goto yystate7
	case this.current == '-':
		goto yystate8
	case this.current == '0':
		goto yystate9
	case this.current == '1':
		goto yystate12
	case this.current == '2':
		goto yystate13
	case this.current == '4':
		goto yystate14
	case this.current == '8':
		goto yystate15
	case this.current == ':':
		goto yystate16
	case this.current == '<':
		goto yystate17
	case this.current == 'B':
		goto yystate20
	case this.current == 'D':
		goto yystate28
	case this.current == 'F':
		goto yystate37
	case this.current == 'O':
		goto yystate46
	case this.current == 'Q':
		goto yystate55
	case this.current == 'T':
		goto yystate64
	case this.current == 'W':
		goto yystate73
	case this.current == 'X':
		goto yystate81
	case this.current == 'Y':
		goto yystate92
	case this.current == 'Z':
		goto yystate103
	case this.current == '[':
		goto yystate114
	case this.current == '\n':
		goto yystate4
	case this.current == '\t' || this.current == '\r' || this.current == ' ':
		goto yystate3
	case this.current == '\x00':
		goto yystate2
	case this.current == ']':
		goto yystate115
	case this.current >= 'a' && this.current <= 'z':
		goto yystate116
	}

yystate2:
	this.next()
	goto yyrule27

yystate3:
	this.next()
	switch {
	default:
		goto yyrule2
	case this.current == '\t' || this.current == '\r' || this.current == ' ':
		goto yystate3
	}

yystate4:
	this.next()
	goto yyrule15

yystate5:
	this.next()
	goto yyrule19

yystate6:
	this.next()
	goto yyrule17

yystate7:
	this.next()
	goto yyrule22

yystate8:
	this.next()
	goto yyrule18

yystate9:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'X' || this.current == 'x':
		goto yystate10
	}

yystate10:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current >= '0' && this.current <= '9' || this.current >= 'a' && this.current <= 'f':
		goto yystate11
	}

yystate11:
	this.next()
	switch {
	default:
		goto yyrule14
	case this.current >= '0' && this.current <= '9' || this.current >= 'a' && this.current <= 'f':
		goto yystate11
	}

yystate12:
	this.next()
	goto yyrule23

yystate13:
	this.next()
	goto yyrule24

yystate14:
	this.next()
	goto yyrule25

yystate15:
	this.next()
	goto yyrule26

yystate16:
	this.next()
	goto yyrule16

yystate17:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == '>':
		goto yystate18
	case this.current >= '\x01' && this.current <= '\t' || this.current >= '\v' && this.current <= '=' || this.current >= '?' && this.current <= 'ÿ':
		goto yystate17
	}

yystate18:
	this.next()
	switch {
	default:
		goto yyrule1
	case this.current == ':':
		goto yystate19
	}

yystate19:
	this.next()
	switch {
	default:
		goto yyrule1
	}

yystate20:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'Y':
		goto yystate21
	}

yystate21:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'T':
		goto yystate22
	}

yystate22:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'E':
		goto yystate23
	}

yystate23:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == ' ':
		goto yystate24
	}

yystate24:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'P':
		goto yystate25
	}

yystate25:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'T':
		goto yystate26
	}

yystate26:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'R':
		goto yystate27
	}

yystate27:
	this.next()
	goto yyrule3

yystate28:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'W':
		goto yystate29
	}

yystate29:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'O':
		goto yystate30
	}

yystate30:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'R':
		goto yystate31
	}

yystate31:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'D':
		goto yystate32
	}

yystate32:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == ' ':
		goto yystate33
	}

yystate33:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'P':
		goto yystate34
	}

yystate34:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'T':
		goto yystate35
	}

yystate35:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'R':
		goto yystate36
	}

yystate36:
	this.next()
	goto yyrule4

yystate37:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'W':
		goto yystate38
	}

yystate38:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'O':
		goto yystate39
	}

yystate39:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'R':
		goto yystate40
	}

yystate40:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'D':
		goto yystate41
	}

yystate41:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == ' ':
		goto yystate42
	}

yystate42:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'P':
		goto yystate43
	}

yystate43:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'T':
		goto yystate44
	}

yystate44:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'R':
		goto yystate45
	}

yystate45:
	this.next()
	goto yyrule5

yystate46:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'W':
		goto yystate47
	}

yystate47:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'O':
		goto yystate48
	}

yystate48:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'R':
		goto yystate49
	}

yystate49:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'D':
		goto yystate50
	}

yystate50:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == ' ':
		goto yystate51
	}

yystate51:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'P':
		goto yystate52
	}

yystate52:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'T':
		goto yystate53
	}

yystate53:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'R':
		goto yystate54
	}

yystate54:
	this.next()
	goto yyrule6

yystate55:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'W':
		goto yystate56
	}

yystate56:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'O':
		goto yystate57
	}

yystate57:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'R':
		goto yystate58
	}

yystate58:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'D':
		goto yystate59
	}

yystate59:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == ' ':
		goto yystate60
	}

yystate60:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'P':
		goto yystate61
	}

yystate61:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'T':
		goto yystate62
	}

yystate62:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'R':
		goto yystate63
	}

yystate63:
	this.next()
	goto yyrule7

yystate64:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'B':
		goto yystate65
	}

yystate65:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'Y':
		goto yystate66
	}

yystate66:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'T':
		goto yystate67
	}

yystate67:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'E':
		goto yystate68
	}

yystate68:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == ' ':
		goto yystate69
	}

yystate69:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'P':
		goto yystate70
	}

yystate70:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'T':
		goto yystate71
	}

yystate71:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'R':
		goto yystate72
	}

yystate72:
	this.next()
	goto yyrule8

yystate73:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'O':
		goto yystate74
	}

yystate74:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'R':
		goto yystate75
	}

yystate75:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'D':
		goto yystate76
	}

yystate76:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == ' ':
		goto yystate77
	}

yystate77:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'P':
		goto yystate78
	}

yystate78:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'T':
		goto yystate79
	}

yystate79:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'R':
		goto yystate80
	}

yystate80:
	this.next()
	goto yyrule9

yystate81:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'M':
		goto yystate82
	}

yystate82:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'M':
		goto yystate83
	}

yystate83:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'W':
		goto yystate84
	}

yystate84:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'O':
		goto yystate85
	}

yystate85:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'R':
		goto yystate86
	}

yystate86:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'D':
		goto yystate87
	}

yystate87:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == ' ':
		goto yystate88
	}

yystate88:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'P':
		goto yystate89
	}

yystate89:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'T':
		goto yystate90
	}

yystate90:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'R':
		goto yystate91
	}

yystate91:
	this.next()
	goto yyrule10

yystate92:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'M':
		goto yystate93
	}

yystate93:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'M':
		goto yystate94
	}

yystate94:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'W':
		goto yystate95
	}

yystate95:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'O':
		goto yystate96
	}

yystate96:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'R':
		goto yystate97
	}

yystate97:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'D':
		goto yystate98
	}

yystate98:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == ' ':
		goto yystate99
	}

yystate99:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'P':
		goto yystate100
	}

yystate100:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'T':
		goto yystate101
	}

yystate101:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'R':
		goto yystate102
	}

yystate102:
	this.next()
	goto yyrule11

yystate103:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'M':
		goto yystate104
	}

yystate104:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'M':
		goto yystate105
	}

yystate105:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'W':
		goto yystate106
	}

yystate106:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'O':
		goto yystate107
	}

yystate107:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'R':
		goto yystate108
	}

yystate108:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'D':
		goto yystate109
	}

yystate109:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == ' ':
		goto yystate110
	}

yystate110:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'P':
		goto yystate111
	}

yystate111:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'T':
		goto yystate112
	}

yystate112:
	this.next()
	switch {
	default:
		goto yyabort
	case this.current == 'R':
		goto yystate113
	}

yystate113:
	this.next()
	goto yyrule12

yystate114:
	this.next()
	goto yyrule20

yystate115:
	this.next()
	goto yyrule21

yystate116:
	this.next()
	switch {
	default:
		goto yyrule13
	case this.current >= 'a' && this.current <= 'z':
		goto yystate116
	}

yyrule1: // \<.*\>\:?
	{
		/*  ignore, gdb disassemble will put some comment information in `<>` */
		goto yystate0
	}
yyrule2: // [ \t\r]+
	{
		/* ignore */
		goto yystate0
	}
yyrule3: // (BYTE\ PTR)
	{
		{
			lval.u64 = BYTE_PTR
			return BYTE_PTR
		}
		goto yystate0
	}
yyrule4: // (DWORD\ PTR)
	{
		{
			lval.u64 = DWORD_PTR
			return DWORD_PTR
		}
		goto yystate0
	}
yyrule5: // (FWORD\ PTR)
	{
		{
			lval.u64 = FWORD_PTR
			return FWORD_PTR
		}
		goto yystate0
	}
yyrule6: // (OWORD\ PTR)
	{
		{
			lval.u64 = OWORD_PTR
			return OWORD_PTR
		}
		goto yystate0
	}
yyrule7: // (QWORD\ PTR)
	{
		{
			lval.u64 = QWORD_PTR
			return QWORD_PTR
		}
		goto yystate0
	}
yyrule8: // (TBYTE\ PTR)
	{
		{
			lval.u64 = TBYTE_PTR
			return TBYTE_PTR
		}
		goto yystate0
	}
yyrule9: // (WORD\ PTR)
	{
		{
			lval.u64 = WORD_PTR
			return WORD_PTR
		}
		goto yystate0
	}
yyrule10: // (XMMWORD\ PTR)
	{
		{
			lval.u64 = XMMWORD_PTR
			return XMMWORD_PTR
		}
		goto yystate0
	}
yyrule11: // (YMMWORD\ PTR)
	{
		{
			lval.u64 = YMMWORD_PTR
			return YMMWORD_PTR
		}
		goto yystate0
	}
yyrule12: // (ZMMWORD\ PTR)
	{
		{
			lval.u64 = ZMMWORD_PTR
			return ZMMWORD_PTR
		}
		goto yystate0
	}
yyrule13: // [a-z]+
	{
		{
			/* LOWERCASESTRING */
			lval.str = string(this.buf)
			return LOWERCASESTRING
		}
		goto yystate0
	}
yyrule14: // 0[xX][0-9a-f]+
	{
		{
			/* HEXNUMBER */
			var err error
			if lval.u64, err = strconv.ParseUint(string(this.buf), 0, 64); err != nil {
				panic(err)
			}
			return HEXNUMBER
		}
		goto yystate0
	}
yyrule15: // \n
	{
		return '\n'
	}
yyrule16: // :
	{
		return ':'
	}
yyrule17: // \+
	{
		return '+'
	}
yyrule18: // -
	{
		return '-'
	}
yyrule19: // \*
	{
		return '*'
	}
yyrule20: // \[
	{
		return '['
	}
yyrule21: // \]
	{
		return ']'
	}
yyrule22: // ,
	{
		return ','
	}
yyrule23: // 1
	{
		return '1'
	}
yyrule24: // 2
	{
		return '2'
	}
yyrule25: // 4
	{
		return '4'
	}
yyrule26: // 8
	{
		return '8'
	}
yyrule27: // \0
	if true { // avoid go vet determining the below panic will not be reached
		return 0
	}
	panic("unreachable")

yyabort: // no lexem recognized
	//
	// silence unused label errors for build and satisfy go vet reachability analysis
	//
	{
		if false {
			goto yyabort
		}
		if false {
			goto yystate0
		}
		if false {
			goto yystate1
		}
	}

	panic(fmt.Errorf("lex error; buf: %s; current: %d;", this.buf, this.current))
	return 0
}
