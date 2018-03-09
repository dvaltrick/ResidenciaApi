package service

import (
	"../../residencia-api/model"
	"github.com/jinzhu/gorm"
)

func Gravar(instituicao model.Instituicao, db *gorm.DB) {
	db.Create(instituicao)
}
