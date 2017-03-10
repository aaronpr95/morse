package recursos_morse

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"os/exec"
	"strings"
)

// Funcion que devuelve un Slice de strings con lo que ha leido del fichero
// Recive nombre del fichero
func Guardar_fichero(nombre string) string {
	// Abrimos el fichero y guardamos en el descriptor_de_fichero
	fd, err := os.Open(nombre)
	if err != nil {
		// Comprobamos si ha habido errores
		fmt.Println("Error al abir el fichero", nombre)
		log.Fatal(err)
	}

	// Stat nos devuelve la informacion acerca del fichero
	fi, err := fd.Stat()
	if err != nil {
		// Could not obtain stat, handle error
		log.Fatal(err)
	}
	// Con size obtenemos el tamaño del fichero
	tamano := fi.Size()

	// Creamos un slice del tamaño del fichero
	datos := make([]byte, tamano)

	// con read guardamos en datos lo que hay en el fichero
	_, err1 := fd.Read(datos)
	if err1 != nil {
		log.Fatal(err1)
	}
	
	// Cerramos el fichero
	err2 := fd.Close()
	if err2 != nil {
		// Comprobamos si ha habido errores
		fmt.Println("Error al cerrar el fichero", nombre)
		log.Fatal(err2)
	}

	// Devolvemos datos en forma de string
	return strings.ToLower(string(datos))
}

// Función que devuelve en codigo morse la cadena que le pasamos como string
func Cadena_morse(cadena string) []Morse {
	// guardamos el tamaño de la cadena para crear el array de morse
	n := len(cadena)
	mi_morses := make([]Morse, n)

	// Se recorre la cadena introduciendo valores en el array de Morses
	for i, v := range cadena {
		mi_morses[i] = A_morse[string(v)]
	}

	return mi_morses
}

// Función que crea un fichero con el nombre que se le pase
// Y que guarda un mensaje en codigo morse
func Fichero_morse(morses []Morse, nombre string) {
	fd, err := os.Create(nombre)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range morses {

		cadena := string(v) + " "

		// Escribimos la cadena en el fichero
		_, err1 := fd.WriteString(cadena)
		if err1 != nil {
			log.Fatal(err1)
		}
	}

	// Cerramos el fichero que hemos creado
	err2 := fd.Close()
	if err2 != nil {
		// Comprobamos si ha habido errores
		fmt.Println("Error al cerrar el fichero", nombre)
		log.Fatal(err2)
	}
}

// Funcion que se pide una cadena al usuario y se escribe en un fichero
func Escribir_fichero() {
	// Variables locales
	var nombre string
	_ = nombre

	// Se pide al usuario que escriba el mensaje
	fmt.Println("Escribe el contenido que quieres guardar (escribe un string y se copia en Morse):")
	
	// Pasamos a leer el mensaje
	// Creamos un nuevo buffer para leer de la entrada estandar
	in := bufio.NewReader(os.Stdin)
	// leemos del teclado en el buffer
	// con readString se lee del teclado y devuelve una cadena hasta
	// el carácter pasado como parametro
	contenido, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Error en la entrada de datos por teclado")
		log.Fatal(err)
	}

	// Pasamos el mensaje a codigo morse
	mi_morse := Cadena_morse(strings.ToLower(contenido))
	fmt.Println(mi_morse)

	fmt.Println("Nombre del fichero que quieres guardar?:")
	_, err1 := fmt.Scanf("%s", &nombre )
	if err1 != nil {
		log.Fatal(err1)
	}

	// Escribimos el código en el fichero
	Fichero_morse(mi_morse, nombre)
}

// Funcion que lee un fichero en morse lo traduche y lo muestra por pantalla
func Leer_fichero() {
	//variables locales
	var nombre string // para guardar el nombre del fichero que se quiere leer

	fmt.Println("Los ficheros disponibles son los siguientes:")
	System("ls")
	fmt.Println("Introduce el nombre del fichero que quieres leer:")
	_, err := fmt.Scanf("%s", &nombre)
	if err != nil {
		fmt.Println("Error al introducir el nombre del fichero")
		log.Fatal(err)
	}

	contenido := Guardar_fichero(nombre)

	// IMPLEMENTAR
	// leer el string en morse e ir traduciendo a char normal
	
	runes := make([]rune, 0, 8)

	for i, v := range contenido {
		// Si es igual a punto o raya lo concatenamos con la variable auxiliar
		if v == '.' || v == '-' {
			if i > 1 && contenido[i-1] == ' ' {
				if i > 2 && contenido[i-2] == ' ' {
					fmt.Print(" ")
				}
				runes = runes[:0]
			}
			// Lo añadimos al slice de runes
			runes = append(runes, v)
		} else if v == ' '{
 			// Si es un espacio y el anterior es . ó - mostrar rune y poner a cero
			if contenido[i-1] == '.' || contenido[i-1] == '-' {
				fmt.Print(Morse(runes).Traducir()) // para mostrar traducido
				//Morse(runes).Imprimir()
				runes = runes[:0]
			}
				runes = append(runes, v)
			
		} else {
			// Si no es nada de lo anterior cerrar el bucle
			break
		}
	}
	// Hacemos salto de linea
	fmt.Println()
}

// Funcion que ejecuta el comando que se le pase como parametro
func System(comando string) {

	if valor ,err := exec.Command(comando).Output(); err != nil {
		fmt.Println("Error al ejecutar el comando", comando)	
	} else {
		fmt.Printf("\n%s\n",valor)
	}

}

// Declaramos el método imprimir para morse
func (m Morse)Imprimir() {
	fmt.Println(m)
}

// Método traducir de un morse
func (m Morse)Traducir() string {
	return A_char[m]
}