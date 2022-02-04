package main

import (
	"fmt"
	"time"
)


func main()  {
	canal1 := make(chan time.Duration) //Creamos un canal de comunicación de tipo time.Duration
	go bucle(canal1)
	fmt.Println("Llegué hasta aquí")

	mensaje := <- canal1 //Detiene la ejecución hasta que recibe un mensaje en canal1. El valor de canal1 se va a grabar en la variable mensaje
	fmt.Println(mensaje)
}

func bucle(canal chan time.Duration)  {
	inicio := time.Now()
	for i := 0; i<1000000000; i++ {

	}
	final := time.Now()
	canal <- final.Sub(inicio) //guardamos en el channel canal el resultado de final-inicio. La función final.Sub() devuelve la duración de final-inicio
}
