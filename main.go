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
	//-------------------------------------------------------
	fmt.Printf("%t\n", num)
	//fmt.Println(fib(num))

}

/*
func fib(n int) int {
	if n <= 1 {
		return n
	}
	total := []byte("teste")
	err := ioutil.WriteFile("C:/Users/cesar.gomes/Desktop/go teste/teste.txt", total, 0644)

	if err != nil {
		panic(err)
	}
	return fib(n-1) + fib(n-2)

}
*/
