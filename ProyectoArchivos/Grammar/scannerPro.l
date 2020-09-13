%yyc c
%yyn c = l.Next()
%yym l.Mark()
%option case-insensitive

%{

package Grammar

import (
  "unicode"
  "io"
  "bufio"
  "modernc.org/golex/lex"
  "go/token"
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
  if (err != nil) { 
      panic(err) 
  }
  return &lexer{anLex}
}

func (l *lexer) Lex(lval *yySymType) int {
  c := l.Enter()
  yyErrorVerbose = true
%}


unicodeLetra      \x80
unicodeDigito     \x81 
numero_           [0-9]+|{unicodeDigito}+|[-][0-9]+ 
id_               {letra_}({letra_}|[0-9])*
letra_            [_a-zA-Z]|{unicodeLetra}
arrow_			  \-\>
route_            (\/[a-zA-z]([_a-zA-z]|[0-9])*)+(\/)?(\.[a-zA-z]+)?
arch_dsk_		  {id_}\.dsk


%%
	c = l.Rule0()

[ \t\r\n]+


exec        lval.cad = string(l.TokenBytes(nil)); return T_EXEC
path		lval.cad = string(l.TokenBytes(nil)); return T_PATH
{arrow_}    lval.cad = string(l.TokenBytes(nil)); return T_ARROW
{route_}	lval.cad = string(l.TokenBytes(nil)); return T_ROUTE
mkdisk	 	lval.cad = string(l.TokenBytes(nil)); return T_MKDISK 

size		lval.cad = string(l.TokenBytes(nil)); return T_SIZE
name		lval.cad = string(l.TokenBytes(nil)); return T_NAME
unit		lval.cad = string(l.TokenBytes(nil)); return T_UNIT
rmdisk		lval.cad = string(l.TokenBytes(nil)); return T_RMDISK

fdisk		lval.cad = string(l.TokenBytes(nil)); return T_FDISK
type		lval.cad = string(l.TokenBytes(nil)); return T_TYPE
fit			lval.cad = string(l.TokenBytes(nil)); return T_FIT 
delete		lval.cad = string(l.TokenBytes(nil)); return T_DELETE
add			lval.cad = string(l.TokenBytes(nil)); return T_ADD
rep         lval.cad = string(l.TokenBytes(nil)); return T_REP

pause       lval.cad = string(l.TokenBytes(nil)); return T_PAUSE
{letra_}    lval.cad = string(l.TokenBytes(nil)); return T_LETRA
{id_}       lval.cad = string(l.TokenBytes(nil)); return T_ID 
{numero_}   lval.cad = string(l.TokenBytes(nil)); return T_NUMERO
{arch_dsk_} lval.cad = string(l.TokenBytes(nil)); return T_ARC_DKS


%%

	if c, ok := l.Abort(); ok {
		return int(c)
	}

	goto yyAction

}