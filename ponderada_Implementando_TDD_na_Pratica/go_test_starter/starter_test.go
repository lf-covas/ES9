package starter_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	starter "github.com/williaminfante/go_test_starter"
)

// TestSayHello testa a função SayHello para verificar se a saudação está correta ao utilizar a variavel inserida na impressão
func TestSayHello(t *testing.T) {

// RED: Teste falhará inicialmente porque a função SayHello ainda não está implementada corretamente
  greeting := starter.SayHello("Romualdo")

// GREEN: Após implementar a função SayHello corretamente, este teste deve passar, mostrando que a nossa função está correta ;)
// O teste verifica se a função retorna a saudação esperada "Hello Romualdo. Welcome!".
  assert.Equal(t, "Hello Romualdo. Welcome!", greeting)

// REFACTOR: Após garantir que o teste passa, o código pode ser refatorado para melhorar a implementação, sem alterar a lógica ou quebrar o teste

}

// TestOddOrEven testa a função OddOrEven para verificar se ela classifica corretamente números como ímpares ou pares
func TestOddOrEven(t *testing.T) {
	// Teste para números não negativos
	// RED: Testes inicialmente falharão porque a função OddOrEven não estará completa.
	t.Run("Check Non Negative Numbers", func(t *testing.T) {
	  assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
	  assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
	  assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
	  // GREEN: Após implementar a função OddOrEven para tratar números pares, ímpares e zero, estes testes devem passar, confirmando que a função funciona corretamente.
	  })
	
	  // Teste para números negativos
	t.Run("Check Negative Numbers", func(t *testing.T) {
	// RED: Testes inicialmente falharão porque a função OddOrEven pode não lidar corretamente com números negativos.
	  assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
	  assert.Equal(t, "-42 is an even number", starter.OddOrEven(-42))

	// GREEN: Após ajustar a função para lidar com números negativos, estes testes devem passar.
	// REFACTOR: A refatoração pode incluir melhorias no cálculo ou tratamento dos números, mantendo a lógica intacta e garantindo que todos os testes continuem a passar.
	})
  }

  // TestCheckHealth testa a função CheckHealth para verificar se a resposta HTTP está correta.
  func TestCheckHealth(t *testing.T) {
	t.Run("Check health status", func(t *testing.T) {
	// RED: Teste falhará inicialmente se a função CheckHealth não retornar a resposta correta.
	  req := httptest.NewRequest("GET", "http://mysite.com/example", nil)
	  writer := httptest.NewRecorder()
	  starter.CheckHealth(writer, req)
	  response := writer.Result()
	  body, _ := io.ReadAll(response.Body)
  
	// GREEN: Após implementar a função CheckHealth corretamente, o teste deve passar, verificando que a resposta contém "health check passed".
	  assert.Equal(t, "health check passed", string(body))

	// REFACTOR: Após garantir que o teste passa, a função pode ser refatorada para melhorar o código sem alterar a lógica ou quebrar o teste.
	})
  }