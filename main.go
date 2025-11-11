package main

import (
	"github.com/fabiomassucatto/gin-go-api/models"
	"github.com/fabiomassucatto/gin-go-api/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Fabio Henrique", CPF: "123.456.789-00", RG: "12.345.678-9"},
		{Nome: "Maria Silva", CPF: "987.654.321-00", RG: "98.765.432-1"},
		{Nome: "João Souza", CPF: "456.789.123-00", RG: "45.678.912-3"},
	}
	routes.HandleRequests()
}
