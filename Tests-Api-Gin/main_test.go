package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestFalha(t *testing.T) {
	t.Fatalf("Teste Falhou de Proposito, não se preocupe")
}
