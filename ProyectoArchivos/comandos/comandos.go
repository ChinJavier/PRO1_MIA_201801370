package comandos

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var signatureDiscos []int64

var TamanioSint string = ""
var RutaSint string = ""
var NameSint string = ""
var DimenSin string = "v"

var TypePartSin string = ""

var ReporteMbrInit string = "digraph G{ tbl [shape=plaintext label=< <table border='0' cellborder='1' color='blue' cellspacing='0'> \n"
var ReporteMbrFin string = "      </table> >]; } \n"
var ReporteMbrAux string = ""

/*func ImprimirSint() {
	fmt.Println("¿¿¿¿", TamanioSint)
	fmt.Println("¿¿¿¿", RutaSint)
	fmt.Println("¿¿¿¿", NameSint)
	fmt.Println("¿¿¿¿", DimenSin)
} */

type Particion struct {
	StatusParticion byte
	TypeParticion   byte
	FitParticion    byte
	StartParticion  int64
	SizeParticion   int64
	NameParticion   [16]byte
}

type Mbr struct {
	TamanioMbr     int64
	FechaMbr       [18]byte
	SignatureMbr   int64
	ParticionesMbr [4]Particion
}

func Esperar() {
	fmt.Print("... ")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Println("")
}

func CrearArchivo(tam string, ruta string, name string, dimension string) {
	soloCarps := ruta
	dimension = strings.ToLower(dimension)      //la dimension está en minuscula
	tamanio, _ := strconv.ParseInt(tam, 10, 64) //mi tamanio esta en el tipo correcto
	ruta = strings.TrimSpace(ruta)
	name = strings.TrimSpace(name)
	ruta = ruta + name
	if tamanio > 0 {

	} else {
		fmt.Println("SIZE INCORRECTO DEL DISCO")
		tamanio = 1024
	}

	if dimension == "k" {
		tamanio += tamanio * 1024
	} else {
		tamanio += tamanio * 1024 * 1024
	}

	err := os.MkdirAll(soloCarps, 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(soloCarps)
	fmt.Println(ruta)
	manejador, err := os.Create(ruta)
	defer manejador.Close()

	if err != nil {
		log.Fatal(err)
	}

	var aux int8 = 0

	var binarioInit bytes.Buffer
	binary.Write(&binarioInit, binary.BigEndian, &aux)
	escribirBytes(manejador, binarioInit.Bytes())

	manejador.Seek(tamanio-1, 0)
	var binarioFin bytes.Buffer
	binary.Write(&binarioFin, binary.BigEndian, &aux)
	escribirBytes(manejador, binarioFin.Bytes())

	manejador.Seek(0, 0)
	indiceMaster := Mbr{}
	indiceMaster.TamanioMbr = tamanio
	indiceMaster.SignatureMbr = getSignature()
	copy(indiceMaster.FechaMbr[:], time.Now().String())
	indiceMaster.ParticionesMbr[0].StatusParticion = 'f'
	indiceMaster.ParticionesMbr[0].SizeParticion = 0
	indiceMaster.ParticionesMbr[0].FitParticion = 'w'
	indiceMaster.ParticionesMbr[0].TypeParticion = '-'
	indiceMaster.ParticionesMbr[1].StatusParticion = 'f'
	indiceMaster.ParticionesMbr[1].SizeParticion = 0
	indiceMaster.ParticionesMbr[1].FitParticion = 'w'
	indiceMaster.ParticionesMbr[1].TypeParticion = '-'
	indiceMaster.ParticionesMbr[2].StatusParticion = 'f'
	indiceMaster.ParticionesMbr[2].SizeParticion = 0
	indiceMaster.ParticionesMbr[2].FitParticion = 'w'
	indiceMaster.ParticionesMbr[2].TypeParticion = '-'
	indiceMaster.ParticionesMbr[3].StatusParticion = 'f'
	indiceMaster.ParticionesMbr[3].SizeParticion = 0
	indiceMaster.ParticionesMbr[3].FitParticion = 'w'
	indiceMaster.ParticionesMbr[3].TypeParticion = '-'
	var binarioMbr bytes.Buffer
	binary.Write(&binarioMbr, binary.BigEndian, &indiceMaster)
	escribirBytes(manejador, binarioMbr.Bytes())
	reinit()
}

func escribirBytes(manejador *os.File, bytes []byte) {

	_, err := manejador.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}

}

func getSignature() int64 {
	signatur := int64(rand.Intn(100))
	for nuevoSig(signatur) {
		signatur = int64(rand.Intn(100))
	}
	signatureDiscos = append(signatureDiscos, signatur)
	return signatur
}

func nuevoSig(elRandom int64) bool {
	for i := 0; i < len(signatureDiscos); i++ {
		if signatureDiscos[i] == elRandom {
			return true
		}
	}
	return false
}

func reinit() {
	TamanioSint = ""
	RutaSint = ""
	NameSint = ""
	DimenSin = "v"
}

func reinit2() {
	TamanioSint = ""
	RutaSint = ""
	NameSint = ""
	DimenSin = "v"
	TypePartSin = ""
}

func LeerArchivo(ruta string) {

	manejadorLec, err := os.OpenFile(ruta, os.O_RDWR, 0644)
	defer manejadorLec.Close()
	if err != nil {
		log.Fatal(err)
	}

	m := Mbr{}

	var size int = binary.Size(m)

	leidoBytes := leerBytes(manejadorLec, size)
	buffer := bytes.NewBuffer(leidoBytes)

	fmt.Println(leidoBytes)

	err = binary.Read(buffer, binary.BigEndian, &m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("******************", m.TamanioMbr)
	fmt.Println(string(m.FechaMbr[:]))
	fmt.Println(m.SignatureMbr)

	for i := 0; i < 4; i++ {
		fmt.Println("Fit: ", string(m.ParticionesMbr[i].FitParticion))
		fmt.Println("Name: ", string(m.ParticionesMbr[i].NameParticion[:]))
		fmt.Println("Tamanio: ", m.ParticionesMbr[i].SizeParticion)
		fmt.Println("Start: ", m.ParticionesMbr[i].StartParticion)
		fmt.Println("Status: ", string(m.ParticionesMbr[i].StatusParticion))
		fmt.Println("Type: ", string(m.ParticionesMbr[i].TypeParticion))
	}

	//ESTO ES PARA GRAPHVIZ

	ReporteMbrAux += "<tr> <td colspan='2'> Reporte MBR </td> </tr> \n"
	ReporteMbrAux += "<tr> <td>Tamanio</td> <td>" + strconv.FormatInt(m.TamanioMbr, 10) + "</td></tr> \n"
	ReporteMbrAux += "<tr> <td>Fecha</td> <td>" + string(m.FechaMbr[:]) + "</td></tr> \n"
	ReporteMbrAux += "<tr> <td>Signature</td> <td>" + strconv.FormatInt(m.SignatureMbr, 10) + "</td></tr> \n"

	for i := 0; i < 4; i++ {
		ReporteMbrAux += "<tr> <td colspan='2'> Particion" + strconv.Itoa(i+1) + "</td> </tr> \n"
		ReporteMbrAux += "<tr> <td>Fit</td> <td>" + string(m.ParticionesMbr[i].FitParticion) + "</td></tr> \n"
		ReporteMbrAux += "<tr> <td>Name</td> <td>" + arregloComoCadena(m.ParticionesMbr[i].NameParticion) + "</td></tr> \n"
		ReporteMbrAux += "<tr> <td>Tamanio</td> <td>" + strconv.FormatInt(m.ParticionesMbr[i].SizeParticion, 10) + "</td></tr> \n"
		ReporteMbrAux += "<tr> <td>Start</td> <td>" + strconv.FormatInt(m.ParticionesMbr[i].StartParticion, 10) + "</td></tr> \n"
		ReporteMbrAux += "<tr> <td>Status</td> <td>" + string(m.ParticionesMbr[i].StatusParticion) + "</td></tr> \n"
		ReporteMbrAux += "<tr> <td>Type</td> <td>" + string(m.ParticionesMbr[i].TypeParticion) + "</td></tr> \n"
	}

	ReporteMbrInit = ReporteMbrInit + ReporteMbrAux + ReporteMbrFin
	fmt.Println(ReporteMbrInit)
	crearDot()
	crearImg()
	reiniciarGraphMbr()
	reinit2()
}

func arregloComoCadena(arre [16]byte) string {
	salida := ""
	for i := 0; i < 16; i++ {
		if arre[i] != 0 {
			salida += string(arre[i])
		}
	}
	return salida
}

func crearDot() {
	man, err := os.Create("/home/repo1.txt")
	defer man.Close()
	if err != nil {
		return
	}
	man.WriteString(ReporteMbrInit)
}

func crearImg() {
	cmCon := exec.Command("dot", "-Tjpg", "/home/repo1.txt", "-o", "/home/repo1.jpg")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmCon.Stdout = &out
	cmCon.Stderr = &stderr
	err := cmCon.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
}

func reiniciarGraphMbr() {
	ReporteMbrInit = "digraph G{ tbl [shape=plaintext label=< <table border='0' cellborder='1' color='blue' cellspacing='0'> \n"
	ReporteMbrFin = "      </table> >]; } \n"
	ReporteMbrAux = ""
}

func leerBytes(manejador *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := manejador.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func EliminarArchivo(ruta string) {
	var seleccion int // Para leer elementos de teclado
	_, err := os.Stat(ruta)
	if err != nil {
		fmt.Println("DISCO NO ENCONTRADO")
		fmt.Println(err)
	} else {
		fmt.Println("¿Desea eliminar el disco? Presione 1 para continuar")
		fmt.Scanf("%d\n", &seleccion)
		if seleccion == 1 {

			err := os.Remove(ruta)
			if err != nil {
				fmt.Printf("Error eliminando archivo: %v\n", err)
			} else {
				fmt.Println("Eliminado correctamente")
			}

		} else {
			fmt.Println("No lo eliminaste")
		}
	}
}

func CrearParticiones() {
	//El tamanio de la particion está en TamanioSint      --------
	//Si esta en 'k'|'m'|'b' esta en DimenSin             --------
	//Sobre que disco voy a trabajar es RutaSint          --------
	//El tipo de particion 'E'|'P' esta en TypePartSint   --------
	//El fit será w                                       --------
	//EL nombre de la particion esta en NameSint          --------
	dimension := strings.ToLower(DimenSin)

	tamanio, _ := strconv.ParseInt(TamanioSint, 10, 64) //mi tamanio esta en el tipo correcto
	//fmt.Println("ojo al tejo", tamanio)
	if tamanio > 0 {

	} else {
		fmt.Println("SIZE INCORRECTO DEL DISCO")
		tamanio = 1024
	}

	if dimension == "k" {
		tamanio += tamanio * 1024
	} else if dimension == "m" {
		tamanio += tamanio * 1024 * 1024
	}

	//vamos a leer el mbr 7u7 y actualizarla
	manejadorLec, err := os.OpenFile(RutaSint, os.O_RDWR, 0644)
	defer manejadorLec.Close()
	if err != nil {
		log.Fatal(err)
	}

	m := Mbr{}

	var size int = binary.Size(m)

	leidoBytes := leerBytes(manejadorLec, size)
	buffer := bytes.NewBuffer(leidoBytes)

	//fmt.Println(leidoBytes)

	err = binary.Read(buffer, binary.BigEndian, &m)
	if err != nil {
		log.Fatal(err)
	}

	var bandera bool = false
	for i := 0; i < 4; i++ {
		if m.ParticionesMbr[i].StatusParticion == 'f' {
			copy(m.ParticionesMbr[i].NameParticion[:], NameSint)
			m.ParticionesMbr[i].StatusParticion = 'v'
			m.ParticionesMbr[i].TypeParticion = TypePartSin[0]
			m.ParticionesMbr[i].FitParticion = 'w'
			m.ParticionesMbr[i].SizeParticion = tamanio
			if i == 0 {
				m.ParticionesMbr[i].StartParticion = int64(binary.Size(m)) + 1
			} else {
				m.ParticionesMbr[i].StartParticion = m.ParticionesMbr[i-1].StartParticion + m.ParticionesMbr[i-1].SizeParticion
			}
			bandera = true
			break
		}

	}
	if !bandera {
		fmt.Println("YA NO HAY ESPACIO")
	}

	//Me voy a echar el disco xd

	erro := os.Remove(RutaSint)
	if erro != nil {
		fmt.Printf("Error eliminando archivo: %v\n", erro)
	} else {
		//fmt.Println("Eliminado correctamente")
	}

	//Escribimos el archivo con el mbr modificado UwU
	soloCarpsFa := RutaSint
	indice := strings.LastIndex(soloCarpsFa, "/")
	soloCarpsFa = soloCarpsFa[0 : indice+1]

	tamanioFa := m.TamanioMbr
	rutaFa := RutaSint

	errFa := os.MkdirAll(soloCarpsFa, 0777)
	if errFa != nil {
		fmt.Println(errFa)
	}
	fmt.Println(soloCarpsFa)
	fmt.Println(rutaFa)
	manejadorFa, err := os.Create(rutaFa)
	defer manejadorFa.Close()

	if err != nil {
		log.Fatal(err)
	}

	var aux int8 = 0

	var binarioInit bytes.Buffer
	binary.Write(&binarioInit, binary.BigEndian, &aux)
	escribirBytes(manejadorFa, binarioInit.Bytes())

	manejadorFa.Seek(tamanioFa-1, 0)
	var binarioFin bytes.Buffer
	binary.Write(&binarioFin, binary.BigEndian, &aux)
	escribirBytes(manejadorFa, binarioFin.Bytes())

	manejadorFa.Seek(0, 0)
	var binarioMbr bytes.Buffer
	binary.Write(&binarioMbr, binary.BigEndian, &m)
	escribirBytes(manejadorFa, binarioMbr.Bytes())
	reinit()
}
