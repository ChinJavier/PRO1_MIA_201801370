// Code generated by golex. DO NOT EDIT.

package Grammar

import (
	"bufio"
	"go/token"
	"io"
	"modernc.org/golex/lex"
	"unicode"
)

type lexer struct {
	*lex.Lexer
}

const (
	classUnicodeLeter = iota + 0x80
	classUnicodeDigit
	classOther
)

func identificadorTipo(actual rune) int {
	if actual >= 0 && actual < 0x80 {
		return int(actual)
	}
	if unicode.IsLetter(actual) {
		return classUnicodeLeter
	}
	if unicode.IsDigit(actual) {
		return classUnicodeDigit
	}
	return classOther
}

//Es como el main
func parseoCompleto(lectura io.Reader, nombreAr string) *lexer {
	fInit := token.NewFileSet()
	archi := fInit.AddFile(nombreAr, -1, 1<<31-1)
	anLex, err := lex.New(archi, bufio.NewReader(lectura), lex.RuneClass(identificadorTipo))
	if err != nil {
		panic(err)
	}
	return &lexer{anLex}
}

func (l *lexer) Lex(lval *yySymType) int {
	c := l.Enter()
	yyErrorVerbose = true

yystate0:
	yyrule := -1
	_ = yyrule
	c = l.Rule0()

	goto yystart1

yyAction:
	switch yyrule {
	case 1:
		goto yyrule1
	case 2:
		goto yyrule2
	case 3:
		goto yyrule3
	case 4:
		goto yyrule4
	case 5:
		goto yyrule5
	case 6:
		goto yyrule6
	case 7:
		goto yyrule7
	case 8:
		goto yyrule8
	case 9:
		goto yyrule9
	case 10:
		goto yyrule10
	case 11:
		goto yyrule11
	case 12:
		goto yyrule12
	case 13:
		goto yyrule13
	case 14:
		goto yyrule14
	case 15:
		goto yyrule15
	case 16:
		goto yyrule16
	case 17:
		goto yyrule17
	case 18:
		goto yyrule18
	case 19:
		goto yyrule19
	case 20:
		goto yyrule20
	}
yystate1:
	c = l.Next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '-':
		goto yystate3
	case c == '/':
		goto yystate6
	case c == 'A' || c == 'a':
		goto yystate11
	case c == 'B' || c == 'C' || c >= 'G' && c <= 'L' || c == 'O' || c == 'Q' || c >= 'V' && c <= 'Z' || c == '_' || c == 'b' || c == 'c' || c >= 'g' && c <= 'l' || c == 'o' || c == 'q' || c >= 'v' && c <= 'z' || c == '\u0080':
		goto yystate19
	case c == 'D' || c == 'd':
		goto yystate20
	case c == 'E' || c == 'e':
		goto yystate26
	case c == 'F' || c == 'f':
		goto yystate30
	case c == 'M' || c == 'm':
		goto yystate37
	case c == 'N' || c == 'n':
		goto yystate43
	case c == 'P' || c == 'p':
		goto yystate47
	case c == 'R' || c == 'r':
		goto yystate54
	case c == 'S' || c == 's':
		goto yystate60
	case c == 'T' || c == 't':
		goto yystate64
	case c == 'U' || c == 'u':
		goto yystate68
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	case c == '\u0081':
		goto yystate72
	case c >= '0' && c <= '9':
		goto yystate4
	}

yystate2:
	c = l.Next()
	yyrule = 1
	l.Mark()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	}

yystate3:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '>':
		goto yystate5
	case c >= '0' && c <= '9':
		goto yystate4
	}

yystate4:
	c = l.Next()
	yyrule = 19
	l.Mark()
	switch {
	default:
		goto yyrule19
	case c >= '0' && c <= '9':
		goto yystate4
	}

yystate5:
	c = l.Next()
	yyrule = 4
	l.Mark()
	goto yyrule4

yystate6:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= 'A' && c <= 'z':
		goto yystate7
	}

yystate7:
	c = l.Next()
	yyrule = 5
	l.Mark()
	switch {
	default:
		goto yyrule5
	case c == '.':
		goto yystate8
	case c == '/':
		goto yystate10
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'z':
		goto yystate7
	}

yystate8:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= 'A' && c <= 'z':
		goto yystate9
	}

yystate9:
	c = l.Next()
	yyrule = 5
	l.Mark()
	switch {
	default:
		goto yyrule5
	case c >= 'A' && c <= 'z':
		goto yystate9
	}

yystate10:
	c = l.Next()
	yyrule = 5
	l.Mark()
	switch {
	default:
		goto yyrule5
	case c == '.':
		goto yystate8
	case c >= 'A' && c <= 'z':
		goto yystate7
	}

yystate11:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == '.':
		goto yystate12
	case c == 'D' || c == 'd':
		goto yystate17
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate12:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'D' || c == 'd':
		goto yystate13
	}

yystate13:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'S' || c == 's':
		goto yystate14
	}

yystate14:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'K' || c == 'k':
		goto yystate15
	}

yystate15:
	c = l.Next()
	yyrule = 20
	l.Mark()
	goto yyrule20

yystate16:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate17:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'D' || c == 'd':
		goto yystate18
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate18:
	c = l.Next()
	yyrule = 15
	l.Mark()
	switch {
	default:
		goto yyrule15
	case c == '.':
		goto yystate12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate19:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == '.':
		goto yystate12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate20:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == '.':
		goto yystate12
	case c == 'E' || c == 'e':
		goto yystate21
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate21:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'L' || c == 'l':
		goto yystate22
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate22:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'E' || c == 'e':
		goto yystate23
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate23:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'T' || c == 't':
		goto yystate24
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate24:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'E' || c == 'e':
		goto yystate25
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate25:
	c = l.Next()
	yyrule = 14
	l.Mark()
	switch {
	default:
		goto yyrule14
	case c == '.':
		goto yystate12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate26:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == '.':
		goto yystate12
	case c == 'X' || c == 'x':
		goto yystate27
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z' || c == '\u0080':
		goto yystate16
	}

yystate27:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'E' || c == 'e':
		goto yystate28
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate28:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'C' || c == 'c':
		goto yystate29
	case c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate29:
	c = l.Next()
	yyrule = 2
	l.Mark()
	switch {
	default:
		goto yyrule2
	case c == '.':
		goto yystate12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate30:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == '.':
		goto yystate12
	case c == 'D' || c == 'd':
		goto yystate31
	case c == 'I' || c == 'i':
		goto yystate35
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate31:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'I' || c == 'i':
		goto yystate32
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate32:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'S' || c == 's':
		goto yystate33
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate33:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'K' || c == 'k':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate34:
	c = l.Next()
	yyrule = 11
	l.Mark()
	switch {
	default:
		goto yyrule11
	case c == '.':
		goto yystate12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate35:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'T' || c == 't':
		goto yystate36
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate36:
	c = l.Next()
	yyrule = 13
	l.Mark()
	switch {
	default:
		goto yyrule13
	case c == '.':
		goto yystate12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate37:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == '.':
		goto yystate12
	case c == 'K' || c == 'k':
		goto yystate38
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate38:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'D' || c == 'd':
		goto yystate39
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate39:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'I' || c == 'i':
		goto yystate40
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate40:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'S' || c == 's':
		goto yystate41
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate41:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'K' || c == 'k':
		goto yystate42
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate42:
	c = l.Next()
	yyrule = 6
	l.Mark()
	switch {
	default:
		goto yyrule6
	case c == '.':
		goto yystate12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate43:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == '.':
		goto yystate12
	case c == 'A' || c == 'a':
		goto yystate44
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate44:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'M' || c == 'm':
		goto yystate45
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate45:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'E' || c == 'e':
		goto yystate46
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate46:
	c = l.Next()
	yyrule = 8
	l.Mark()
	switch {
	default:
		goto yyrule8
	case c == '.':
		goto yystate12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate47:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == '.':
		goto yystate12
	case c == 'A' || c == 'a':
		goto yystate48
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate48:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'T' || c == 't':
		goto yystate49
	case c == 'U' || c == 'u':
		goto yystate51
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'v' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate49:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'H' || c == 'h':
		goto yystate50
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate50:
	c = l.Next()
	yyrule = 3
	l.Mark()
	switch {
	default:
		goto yyrule3
	case c == '.':
		goto yystate12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate51:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'S' || c == 's':
		goto yystate52
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate52:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'E' || c == 'e':
		goto yystate53
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate53:
	c = l.Next()
	yyrule = 16
	l.Mark()
	switch {
	default:
		goto yyrule16
	case c == '.':
		goto yystate12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate54:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == '.':
		goto yystate12
	case c == 'M' || c == 'm':
		goto yystate55
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate55:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'D' || c == 'd':
		goto yystate56
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate56:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'I' || c == 'i':
		goto yystate57
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate57:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'S' || c == 's':
		goto yystate58
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate58:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'K' || c == 'k':
		goto yystate59
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate59:
	c = l.Next()
	yyrule = 10
	l.Mark()
	switch {
	default:
		goto yyrule10
	case c == '.':
		goto yystate12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate60:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == '.':
		goto yystate12
	case c == 'I' || c == 'i':
		goto yystate61
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate61:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'Z' || c == 'z':
		goto yystate62
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Y' || c == '_' || c >= 'a' && c <= 'y' || c == '\u0080':
		goto yystate16
	}

yystate62:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'E' || c == 'e':
		goto yystate63
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate63:
	c = l.Next()
	yyrule = 7
	l.Mark()
	switch {
	default:
		goto yyrule7
	case c == '.':
		goto yystate12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate64:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == '.':
		goto yystate12
	case c == 'Y' || c == 'y':
		goto yystate65
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'x' || c == 'z' || c == '\u0080':
		goto yystate16
	}

yystate65:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'P' || c == 'p':
		goto yystate66
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate66:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'E' || c == 'e':
		goto yystate67
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate67:
	c = l.Next()
	yyrule = 12
	l.Mark()
	switch {
	default:
		goto yyrule12
	case c == '.':
		goto yystate12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate68:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == '.':
		goto yystate12
	case c == 'N' || c == 'n':
		goto yystate69
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate69:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'I' || c == 'i':
		goto yystate70
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate70:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate12
	case c == 'T' || c == 't':
		goto yystate71
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate71:
	c = l.Next()
	yyrule = 9
	l.Mark()
	switch {
	default:
		goto yyrule9
	case c == '.':
		goto yystate12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080':
		goto yystate16
	}

yystate72:
	c = l.Next()
	yyrule = 19
	l.Mark()
	switch {
	default:
		goto yyrule19
	case c == '\u0081':
		goto yystate72
	}

yyrule1: // [ \t\r\n]+

	goto yystate0
yyrule2: // exec
	{
		lval.cad = string(l.TokenBytes(nil))
		return T_EXEC
		goto yystate0
	}
yyrule3: // path
	{
		lval.cad = string(l.TokenBytes(nil))
		return T_PATH
		goto yystate0
	}
yyrule4: // {arrow_}
	{
		lval.cad = string(l.TokenBytes(nil))
		return T_ARROW
		goto yystate0
	}
yyrule5: // {route_}
	{
		lval.cad = string(l.TokenBytes(nil))
		return T_ROUTE
		goto yystate0
	}
yyrule6: // mkdisk
	{
		lval.cad = string(l.TokenBytes(nil))
		return T_MKDISK
		goto yystate0
	}
yyrule7: // size
	{
		lval.cad = string(l.TokenBytes(nil))
		return T_SIZE
		goto yystate0
	}
yyrule8: // name
	{
		lval.cad = string(l.TokenBytes(nil))
		return T_NAME
		goto yystate0
	}
yyrule9: // unit
	{
		lval.cad = string(l.TokenBytes(nil))
		return T_UNIT
		goto yystate0
	}
yyrule10: // rmdisk
	{
		lval.cad = string(l.TokenBytes(nil))
		return T_RMDISK
		goto yystate0
	}
yyrule11: // fdisk
	{
		lval.cad = string(l.TokenBytes(nil))
		return T_FDISK
		goto yystate0
	}
yyrule12: // type
	{
		lval.cad = string(l.TokenBytes(nil))
		return T_TYPE
		goto yystate0
	}
yyrule13: // fit
	{
		lval.cad = string(l.TokenBytes(nil))
		return T_FIT
		goto yystate0
	}
yyrule14: // delete
	{
		lval.cad = string(l.TokenBytes(nil))
		return T_DELETE
		goto yystate0
	}
yyrule15: // add
	{
		lval.cad = string(l.TokenBytes(nil))
		return T_ADD
		goto yystate0
	}
yyrule16: // pause
	{
		lval.cad = string(l.TokenBytes(nil))
		return T_PAUSE
		goto yystate0
	}
yyrule17: // {letra_}
	{
		lval.cad = string(l.TokenBytes(nil))
		return T_LETRA
		goto yystate0
	}
yyrule18: // {id_}
	{
		lval.cad = string(l.TokenBytes(nil))
		return T_ID
		goto yystate0
	}
yyrule19: // {numero_}
	{
		lval.cad = string(l.TokenBytes(nil))
		return T_NUMERO
		goto yystate0
	}
yyrule20: // {arch_dsk_}
	if true { // avoid go vet determining the below panic will not be reached
		lval.cad = string(l.TokenBytes(nil))
		return T_ARC_DKS
		goto yystate0
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

	if c, ok := l.Abort(); ok {
		return int(c)
	}

	goto yyAction

}
