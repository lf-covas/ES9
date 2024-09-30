package main

import (
    "testing"
)

// TestHelloWorld verifica se a função helloWorld retorna a string esperada.
func TestHelloWorld(t *testing.T) {
    expected := "Hello, OpenTelemetry!"
    result := helloWorld()
    if result != expected {
        t.Errorf("Resultado esperado '%s', mas obteve '%s'", expected, result)
    }
}