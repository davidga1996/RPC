package main

import (
	"fmt"
	"net/rpc"
)

type Block struct {
	Alumno       string
	Materia      string
	Calificacion float64
}

func main() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:8000")

	if err != nil {
		fmt.Println(err)
		return
	}

	var op int64

	for {
		fmt.Println("1.- Agregar calificación de una materia")
		fmt.Println("2.- Mostrar el promedio de un Alumno")
		fmt.Println("3.- Mostrar el promedio general")
		fmt.Println("4.- Mostrar el promedio de una materia ")
		fmt.Println("5.- Ver info")
		fmt.Println("6.- Exit")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var name, class string
			var qualification float64
			fmt.Println("Nombre del alumno: ")
			fmt.Scanln(&name)
			fmt.Println("Materia: ")
			fmt.Scanln(&class)
			fmt.Println("Calificación: ")
			fmt.Scanln(&qualification)

			block := Block{name, class, qualification}

			var result string
			err = conn.Call("Server.AgregarCalificiacion", block, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("AgregarCalificiacion", result)
			}
		case 2:
			var alumn string
			fmt.Print("Alumno: ")
			fmt.Scanln(&alumn)

			var result float64
			err = conn.Call("Server.PromedioAlumno", alumn, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("PromedioAlumno", result)
			}
		case 3:
			var result float64
			err = conn.Call("Server.PromedioGeneral", "", &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("PromedioGeneral: ", result)
			}
		case 4:
			var class string
			fmt.Println("Materia: ")
			fmt.Scanln(&class)
			var result float64
			err = conn.Call("Server.PromedioMateria", class, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("PromedioMateria: ", result)
			}
		case 5:
			var result string
			err = conn.Call("Server.Info", "", &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Info", result)
			}
		case 6:
			return
		}
	}
}
