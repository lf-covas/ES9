package starter

import (
	"fmt"
	"math"
	"net/http"
  )

// SayHello retorna uma saudação personalizada para o nome fornecido.
// O nome é inserido na string formatada "Hello [nome]. Welcome!".
func SayHello(name string) string {
  return fmt.Sprintf("Hello %v. Welcome!", name)
}

// OddOrEven determina se um número é ímpar ou par.
// Utiliza a função math.Mod para calcular o resto da divisão por 2.
// Se o resto for 1 ou -1, o número é ímpar; caso contrário, é par.
// Retorna uma string indicando se o número é par ou ímpar.
func OddOrEven(num int) string {
	criteria := math.Mod(float64(num), 2)
	if criteria == 1 || criteria == -1 {
	  return fmt.Sprintf("%v is an odd number", num)
	}
	return fmt.Sprintf("%v is an even number", num)
  }

// CheckHealth escreve uma resposta HTTP simples indicando que o "health check" passou.
// É usada para verificar o status de uma página web através de uma requisição HTTP.
func CheckHealth(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "health check passed")
  }