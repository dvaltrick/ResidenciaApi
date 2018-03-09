package database

import (
	"../../residencia-api/model"
	"github.com/jinzhu/gorm"
)

//Migrate - Atualiza o banco de dados da aplicação
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Instituicao{})
	db.AutoMigrate(&model.Questao{})
	db.AutoMigrate(&model.Prova{})
}
