package controllers

import (
	"github.com/fabiomassucatto/gin-go-api/models"
	"github.com/gin-gonic/gin"
)

func ExibeAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API Diz": "E ai " + nome + ", tudo bomm?",
	})
}
