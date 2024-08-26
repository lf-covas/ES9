package main

import (
	"os"
	"testing"

    // Importar "github.com/stretchr/testify/suite" pode ser útil para organizar testes em suites, 
    // mas por enquanto está comentado até que seja necessário utilizá-lo.
    // "github.com/stretchr/testify/suite"
)

// TestMain é uma função especial que o pacote de testes do Go reconhece.
// Ela é usada para configurar e limpar recursos antes e depois que todos os testes sejam executados.
// Essa função é ideal para ser usada em um ambiente de TDD, onde podemos precisar de uma configuração e desmontagem repetida de recursos ao escrever e executar testes continuamente.
func TestMain(m *testing.M) {
    // Aqui você pode adicionar código para configurar qualquer recurso necessário antes de executar os testes.
    // Por exemplo, você pode iniciar uma conexão com o banco de dados ou definir variáveis de ambiente.

    // m.Run() executa todos os testes definidos no pacote.
    // No contexto de TDD, cada iteração de desenvolvimento deve ser seguida pela execução dos testes.

    code := m.Run()

    // Após os testes serem executados, você pode adicionar código de desmontagem aqui.
    // Isso pode incluir liberar recursos, como fechar conexões de banco de dados, remover arquivos temporários, etc.
    // Em TDD, esta seção é importante para garantir que cada ciclo de testes comece com um estado limpo.

    // os.Exit(code) encerra o programa com o código de saída resultante dos testes.
    os.Exit(code)
}