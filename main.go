package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var mx sync.Mutex

func main() {
	for {
		//1-pergunta o numero ao usuario e guarda em uma variavel
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("qual a posição que você deseja saber o numero da sequência fibonacci(digite 1 para sair)")
		scanner.Scan()
		numero := scanner.Text()
		//-----------------------------------------------------
		// 2-converter uma string em int
		num, err := strconv.Atoi(numero)
		if err != nil {
			log.Fatal(err)
		}
		//-------------------------------------------------------
		//3-Enquanto o numero digitado pelo usuario for diferente de 1 ele executa o programa
		if num != 1 {
			//--------------------------------------------------------------
			{
				//4-criar arquivo e escrever
				wg.Add(2)
				go func() {
					titulo := (numero + "° da sequência Fibonacci " + ".txt") ///4.1-Criar titulo
					file, err := os.Create(titulo)                            ///4.2-criar arquivo                                        //4.7-fechar arquivo
					if err != nil {
						panic(err)
					}
					defer file.Close()
					go func() {
						mx.Lock()
						fib(num) //4.3-executando a função do fibonacci
						mx.Unlock()
					}()
					total := fib(num)             //4.4-guandando o retorno da função em uma variavel
					totals := strconv.Itoa(total) //4.5-convertendo o resultado da função(tipo int) para string e colocando em variavel e tratando o erro
					if err != nil {
						log.Fatal(err)
					}
					file.WriteString(totals) ///4.6-escrever no arquivo usando a variavel ja sendo string
					wg.Wait()
				}()
				//wg.Done()
			}
		} else { //5-saindo se o numero digitado pelo usuario for 1
			fmt.Println("SAINDO...")
			break
		}
	}
	fmt.Println("FIM!!!!!!!")
}

//-------------------------------------------------------
//função do fibonnaci
func fib(n int) int {

	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
