package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// MySuite é uma suíte de testes que contém múltiplos testes relacionados.
type MySuite struct {
	suite.Suite
}

// SetupTest é executado antes de cada teste.
func (suite *MySuite) SetupTest() {
	// Não há necessidade de configurar nada específico para estes testes simples.
}

// TestAddition verifica se a função Add funciona corretamente.
func (suite *MySuite) TestAddition() {
	result := Add(1, 1) // Chama a função Add
	assert.Equal(suite.T(), 2, result, "1 + 1 deve ser igual a 2")
}

// TestSubtraction verifica se a função Subtract funciona corretamente.
func (suite *MySuite) TestSubtraction() {
	result := Subtract(2, 1) // Chama a função Subtract
	assert.Equal(suite.T(), 1, result, "2 - 1 deve ser igual a 1")
}

// TestSuite executa a suíte de testes.
func TestSuite(t *testing.T) {
	suite.Run(t, new(MySuite))
}
