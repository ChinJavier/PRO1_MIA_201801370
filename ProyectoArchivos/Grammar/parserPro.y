%{

package Grammar

import(
  "PRO1_MIA_201801370/ProyectoArchivos/comandos"
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

%token <cad> T_ID T_NUMERO T_LETRA T_EXEC T_PATH T_ARROW T_ROUTE T_PAUSE T_MKDISK T_SIZE T_NAME T_UNIT T_RMDISK T_FDISK T_TYPE T_FIT T_DELETE T_ADD T_ARC_DKS T_REP

%type <cad> main listaMkDisk atributosMkDisk listaFDisk atributosFDisk


%start main

%%


main : T_EXEC '-' T_PATH T_ARROW T_ROUTE   {fmt.Println("Comando EXEC"); execLeer($5)}
     | T_PAUSE                             {fmt.Println("Comando PAUSE"); comandos.Esperar()}
     | T_MKDISK	listaMkDisk		  	       {fmt.Println("Comando MKDISK"); comandos.CrearArchivo(comandos.TamanioSint,comandos.RutaSint,comandos.NameSint, comandos.DimenSin)}
	 | T_RMDISK '-' T_PATH T_ARROW T_ROUTE {fmt.Println("Comando RMDISK"); 	comandos.EliminarArchivo($5)}
	 | T_FDISK listaFDisk                  {fmt.Println("Comando FDISK"); comandos.CrearParticiones()}
	 | T_REP '-' T_PATH T_ARROW T_ROUTE    {comandos.LeerArchivo($5)}
     ;

listaMkDisk : listaMkDisk atributosMkDisk {}
            | atributosMkDisk {}
			;

atributosMkDisk : '-' T_SIZE T_ARROW T_NUMERO {	comandos.TamanioSint = $4 }
	 			| '-' T_PATH T_ARROW T_ROUTE { comandos.RutaSint = $4 }
				| '-' T_NAME T_ARROW T_ARC_DKS { comandos.NameSint = $4 }
				| '-' T_UNIT T_ARROW T_LETRA { comandos.DimenSin = $4 }
	 			;

listaFDisk : listaFDisk atributosFDisk  {}
		   | atributosFDisk {}
	       ;

atributosFDisk : '-' T_SIZE T_ARROW T_NUMERO { comandos.TamanioSint = $4 }
			   | '-' T_UNIT T_ARROW T_LETRA { comandos.DimenSin = $4 }
			   | '-' T_PATH T_ARROW T_ROUTE { comandos.RutaSint = $4 }
			   | '-' T_TYPE T_ARROW T_LETRA { comandos.TypePartSin = $4 }
			   | '-' T_FIT T_ARROW  T_ID {}
			   | '-' T_DELETE T_ARROW  T_ID {}
			   | '-' T_ADD T_ARROW T_NUMERO {}
			   | '-' T_NAME T_ARROW T_ID { comandos.NameSint = $4 }
			   ;

%%

func execLeer(ruta string){

    var bandera bool= true
	archivo, err :=  os.OpenFile(ruta, os.O_RDWR, 0644)
	defer archivo.Close()                      //<---------------------------------------------se ejecuta al retornar
    yyDebug = 0
	var textMult string = ""
	if err != nil {
		fmt.Println("Hubo un error: ",err)
		return
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan(){  
		bandera = true                               //Itera en cada una de las líneas del archivo
		linea := scanner.Text()
		linea = strings.TrimSpace(linea)
		textMult += linea
		if verificarMultilinea(textMult) {				   //Recolectamos la multilinea
			bandera = false
			textMult = textMult[0:len(textMult)-2]	
		}
		if bandera{
			fmt.Println("Estas ejecutando -->  ",textMult)
			l := parseoCompleto(bytes.NewBufferString(textMult),"file.name") 
			yyParse(l)
			textMult = ""
		}
		
	}

}

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



