package main 

import (
	"fmt"
	m "./recursos_morse"
	"os"
)

func main() {
	var opcion int

	// Menu con las opciones del programa
	fmt.Println("Elige una de las siguientes opciones:")
	fmt.Println(" - 1. Leer fichero morse")
	fmt.Println(" - 2. Escribir fichero morse")
	_, err1 := fmt.Scanf("%d", &opcion)
	if err1 != nil {
		fmt.Println("Error al leer la opcion")
		os.Exit(1)
	}

	switch opcion {
	case 1: 
		// fmt.Println("Introduce el nombre del fichero que quieres leer:")
		// Se llama a la funcion leer fichero que traduce un ficheri .ms y lo muestra por pantalla
		m.Leer_fichero()
	case 2:
		// Se llama a la funcion escribir fichero para leer del teclado el mensaje
		// y pasarlo a morse copiado en un fichero
		m.Escribir_fichero()
	default:
		fmt.Println("No has introducido una opcion correcta")
	}


}