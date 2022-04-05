package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	for {
		//pergunta o numero ao usuario e guarda em uma variavel
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("qual a posição que você deseja saber o numero da sequência fibonacci(digite 1 para sair)")
		scanner.Scan()
		numero := scanner.Text()
		// converter uma string em int
		num, err := strconv.Atoi(numero)
		if err != nil {
			log.Fatal(err)
		}
		if num != 1 {
			//--------------------------------------------------------------
			//---------------------------------------------------------------
			{
				fib(num)                      //executando a função do fibonacci
				total := fib(num)             //guandando o retorno da função em uma variavel
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
		} else {
			fmt.Println("VOCE DIGITOU 1!!!!!!!!!")
			break
		}
	}
}

//-------------------------------------------------------
//função do fibonnaci
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
