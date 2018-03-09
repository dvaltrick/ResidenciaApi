package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../residencia-api/controller"
	"../residencia-api/database"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbprovider *gorm.DB

func main() {
	fmt.Println("Teste")

	db, err := gorm.Open("mysql", "admin:admin@tcp(192.168.0.2:3306)/dbresidencia?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error)
	}

	dbprovider = db

	database.Migrate(db)

	router := mux.NewRouter()
	router.HandleFunc("/instituicao", controller.GravarInstituicao).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))

	defer db.Close()
}

func GravarInstituicao(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode("")
}
