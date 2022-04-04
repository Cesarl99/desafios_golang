package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var total []int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("digite um numero")
	scanner.Scan()
	numero := scanner.Text()
	num, err := strconv.Atoi(numero)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%t\n", num)
	fmt.Println(fib(num))

}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	total = fib(n-1) + fib(n-2)
	err := ioutil.WriteFile("C:/Users/cesar.gomes/Desktop/go teste/teste.txt", total, 0644)

	if err != nil {
		panic(err)
	}
	return fib(n-1) + fib(n-2)

}
