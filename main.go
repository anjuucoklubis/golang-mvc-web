package main

import (
	"golang-web-crud/config"
	categorycontroller "golang-web-crud/controllers/categoriesController"
	homecontroller "golang-web-crud/controllers/homeController"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	//1. homepage
	http.HandleFunc("/", homecontroller.Welcome)

	//2. categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)
	log.Println("startingserver is running on port 8000")
	http.ListenAndServe(":8000", nil)
}
