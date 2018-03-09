package model

import "github.com/jinzhu/gorm"

//Instituicao organizadora da prova
type Instituicao struct {
	gorm.Model
	Nome   string
	UF     string
	Cidade string
	Provas []Prova
}

func (Instituicao) TableName() string {
	return "instituicoes"
}
