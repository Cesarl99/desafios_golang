package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//pergunta o numero ao usuario e guarda em uma variavel
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("digite um numero")
	scanner.Scan()
	numero := scanner.Text()
	//---------------------------------------------------------------
	// converter uma string em int
	num, err := strconv.Atoi(numero)
	if err != nil {
		log.Fatal(err)
	}
	//---------------------------------------------------------------
	{
		fib(num)                      //executando a função do fibonacci
		total := fib(num)             //guandando o returno da função em uma variavel
		totals := strconv.Itoa(total) //convertendo o resultado da função(tipo int) para string e colocando em variavel e tratando o erro
		if err != nil {
			log.Fatal(err)
		}
		//criar arquivo e escrever
		titulo := ("resultado do fibonnaci " + numero + ".txt") ///criar arquivo
		file, err := os.Create(titulo)                          ///criar arquivo
		if err != nil {
			panic(err)
		}
		file.WriteString(totals) ///escrever no arquivo quando string
		file.Close()             //fechar arquivo
	}

} //-------------------------------------------------------
//função do fibonnaci
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
