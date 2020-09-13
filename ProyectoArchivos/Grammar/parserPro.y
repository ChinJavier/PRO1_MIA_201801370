%{

package Grammar

import(
  "ProyectoArchivos/comandos"
  "strings"
  "fmt"
  "bytes"
  "bufio"
  "os"
)

%}


%union{
    cad string
}

%token <cad> T_ID T_NUMERO T_LETRA T_EXEC T_PATH T_ARROW T_ROUTE T_PAUSE T_MKDISK T_SIZE T_NAME T_UNIT T_RMDISK T_FDISK T_TYPE T_FIT T_DELETE T_ADD T_ARC_DKS

%type <cad> main listaMkDisk atributosMkDisk atributosFDisk


%start main

%%


main : T_EXEC '-' T_PATH T_ARROW T_ROUTE   {fmt.Println("Comando EXEC")}
     | T_PAUSE                             {fmt.Println("Comando PAUSE")}
     | T_MKDISK	listaMkDisk		  	       {fmt.Println("Comando MKDISK"); comandos.CrearArchivo(comandos.TamanioSint,comandos.RutaSint,comandos.NameSint, comandos.DimenSin)}
	 | T_RMDISK '-' T_PATH				   {fmt.Println("Comando RMDISK")}
	 | T_FDISK atributosFDisk              {fmt.Println("Comando FDISK")}
     ;

listaMkDisk : listaMkDisk atributosMkDisk {}
            | atributosMkDisk {}
			;

atributosMkDisk : '-' T_SIZE T_ARROW T_NUMERO {	comandos.TamanioSint = $4 }
	 			| '-' T_PATH T_ARROW T_ROUTE { comandos.RutaSint = $4 }
				| '-' T_NAME T_ARROW T_ARC_DKS { comandos.NameSint = $4 }
				| '-' T_UNIT T_ARROW T_LETRA { comandos.DimenSin = $4 }
	 			;

atributosFDisk : T_SIZE T_LETRA {}
			   | T_UNIT {}
			   | T_PATH {}
			   | T_TYPE {}
			   | T_FIT {}
			   | T_DELETE {}
			   | T_ADD {}
			   ;

%%

func ScannearEntrada(){
	reader := bufio.NewReader(os.Stdin)
	yyDebug = 0
	var text string
	var textMult string
	var err error
	for {
		text, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ha ocurrido un error")
		}else{															//Leemos el input
			text = strings.TrimSpace(text)
			if verificarMultilinea(text) {								//Recolectamos la multilinea
				text = text[0:len(text)-2]	
				for {
					fmt.Println("...")
					textMult, err = reader.ReadString('\n')
					if err != nil {
						fmt.Println("Ha ocurrido un error")
					}else{
						text += textMult
						text = strings.TrimSpace(text)
					}
					if !verificarMultilinea(text){
						break;
					}else{
						text = text[0:len(text)-2]
					}
				}
			}
			fmt.Println("-->",text)
			l := parseoCompleto(bytes.NewBufferString(text),"file.name") 
			yyParse(l)
		}
	}
}

func verificarMultilinea(cadena string) bool {
	longitud := len(cadena)
	if cadena[longitud-2] == '\\' && cadena[longitud-1] == '*' {   
		return true
	}
	return false
}



