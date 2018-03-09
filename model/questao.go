package model

import "github.com/jinzhu/gorm"

//Questao da prova
type Questao struct {
	gorm.Model
	Texto   string
	ProvaID uint
	Prova   Prova
}

func (Questao) TableName() string {
	return "questoes"
}
