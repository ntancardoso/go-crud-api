package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-crud-api/internal/handlers"
	"go-crud-api/internal/models"
	"go-crud-api/internal/repositories"
	"go-crud-api/internal/services"
	"go-crud-api/pkg"
)

func main() {

	config := pkg.NewConfig()

	dsn := config.DBUser + ":" + config.DBPassword + "@tcp(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName + "?charset=utf8mb4&parseTime=True&loc=UTC"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.Article{})
	if err != nil {
		log.Fatal("Failed to auto migrate schema:", err)
	}

	repo := repositories.NewArticleRepository(db)
	service := services.NewArticleService(repo)
	handler := handlers.NewArticleHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/articles", handler.GetAllArticles).Methods("GET")
	r.HandleFunc("/articles/{id}", handler.GetArticleByID).Methods("GET")
	r.HandleFunc("/articles", handler.CreateArticle).Methods("POST")
	r.HandleFunc("/articles/{id}", handler.UpdateArticle).Methods("PUT")
	r.HandleFunc("/articles/{id}", handler.DeleteArticle).Methods("DELETE")

	log.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
