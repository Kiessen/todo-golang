package main

import (
	"fmt"
	"github.com/ichtrojan/go-todo/controllers"
	"github.com/ichtrojan/go-todo/models"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&models.Note{})
	if err != nil {
		panic(err.Error())
	}
	noteController := &controllers.NoteControllers{}
	router := httprouter.New()
	router.GET("/", noteController.Index)
	router.GET("/create", noteController.Create)
	router.POST("/create", noteController.Store)
	router.GET("/edit/:id", noteController.Edit)
	router.POST("/edit/:id", noteController.Update)
	router.POST("/done/:id", noteController.Done)
	router.POST("/delete/:id", noteController.Delete)

	port := ":1234"
	fmt.Println("Aplikasi jalan di http://localhost:1234")
	log.Fatal(http.ListenAndServe(port, router))
	fmt.Println("aman")
}
