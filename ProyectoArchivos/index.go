package main

import (
	"ProyectoArchivos/comandos"
	"fmt"
)

func main() {
	fmt.Println("Universidad de San Carlos de Guatemala")
	fmt.Println("Facultad de Ingenieria")
	fmt.Println("Javier Chin, 201801370")
	fmt.Println("MIA A+")
	fmt.Println("--->Arrancando...")
	comandos.LeerArchivo("/home/mia/ufalala.dsk")
	//comandos.CrearArchivo("5", "/home/Mia/probando/chinsi/crea/todo/", "paas.bin", "k")
	//Grammar.ScannearEntrada()
}
