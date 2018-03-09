package model

import "github.com/jinzhu/gorm"

//Prova ----
type Prova struct {
	gorm.Model
	Titulo        string
	Ano           int
	Mes           int
	InstituicaoID uint
	Instituicao   Instituicao
	Provas        []Prova
}
