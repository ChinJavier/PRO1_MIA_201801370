package main

import (
	"PRO1_MIA_201801370/ProyectoArchivos/comandos"
	"fmt"
)

func main() {
	fmt.Println("Universidad de San Carlos de Guatemala")
	fmt.Println("Facultad de Ingenieria")
	fmt.Println("Javier Chin, 201801370")
	fmt.Println("MIA A+")
	fmt.Println("--->Arrancando...")
	//comandos.EliminarArchivo("/home/mia/ufalala.dsk")
	comandos.LeerArchivo("/home/mare/discku.dsk")
	//comandos.CrearArchivo("5", "/home/Mia4/probando", "paas.bin", "m")
	//Grammar.ScannearEntrada()
}
