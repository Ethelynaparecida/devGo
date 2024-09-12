package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hellooo world")

	var verdadeiro, falso bool = true, false
	if verdadeiro && falso {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("FALSO")
	}

	fmt.Printf("Type: %T - value: %v\n", true, true)

	fmt.Printf("Type: %T value: %v\n", "stp", "step")
	fmt.Printf("Type: %T value: %v\n", 1, 1)
	fmt.Printf("Type: %T value: %v\n", "1", "1")
	fmt.Printf("Type: %T value: %v\n", 1.23, 1.23)

	/*ls, err := net.Listen("tcp", "5000")
		if err != nil {
			panic(err)

		}
		defer ls.Close()

		for{
			con, erro := ls.Accept()
			if erro != nil {
	           panic(err)
	        }
			go func(con net.Conn){

			}(con)
		}*/
}
