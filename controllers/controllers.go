package controllers

import (
	"net/http"

	"github.com/fabiomassucatto/gin-go-api/database"
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

func CriaAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}
