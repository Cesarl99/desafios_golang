package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var total []int

func main() {
	//pergunta o numero ao usuario e guarda em uma variavel
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("digite um numero")
	scanner.Scan()
	numero := scanner.Text()
	fmt.Printf("%t\n", numero)
	//---------------------------------------------------------------
	// converter uma string em int
	num, err := strconv.Atoi(numero)
	if err != nil {
		log.Fatal(err)
	}
	//---------------------------------------------------------------
	{
		//criar arquivo e escrever
		titulo := ("resultado do fibonnaci " + numero + ".txt") ///criar arquivo
		file, err := os.Create(titulo)                          ///criar arquivo
		if err != nil {
			panic(err)
		}
		data := []byte("bom dia\n")           //o que vai escrever no arquivo
		file.Write(data)                      //escrever no arquivo
		file.WriteString("isto e uma string") ///escrever no arquivo quando string
		file.Close()                          //fechar arquivo
	}

	//-------------------------------------------------------
	fmt.Printf("%t\n", num)
	//fmt.Println(fib(num))

}

/*
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
*/
