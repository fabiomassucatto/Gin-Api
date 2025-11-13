package database

import (
	"log"

	"github.com/fabiomassucatto/gin-go-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Não foi possível conectar com o banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
