package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-crud-api/internal/models"
	"go-crud-api/internal/services"
	"go-crud-api/pkg"

	"github.com/gorilla/mux"
)

type ArticleHandler struct {
	service *services.ArticleService
}

func NewArticleHandler(service *services.ArticleService) *ArticleHandler {
	return &ArticleHandler{service: service}
}

func (h *ArticleHandler) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := h.service.GetAllArticles()
	if err != nil {
		pkg.Logger.Errorw("Failed to get all articles", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func (h *ArticleHandler) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	// idStr := r.URL.Query().Get("id")

	if idStr == "" {
		pkg.Logger.Errorw("ID is required", "error", "missing ID parameter")
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.Logger.Errorw("Invalid ID format", "error", err)
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	article, err := h.service.GetArticleByID(uint(id))
	if err != nil {
		pkg.Logger.Errorw("Failed to get article by ID", "id", id, "error", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

func (h *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article models.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		pkg.Logger.Errorw("Failed to decode request body", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateArticle(&article); err != nil {
		pkg.Logger.Errorw("Failed to create article", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ArticleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		pkg.Logger.Errorw("ID is required", "error", "missing ID parameter")
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.Logger.Errorw("Invalid ID format", "error", err)
		http.Error(w, "Invalid ID Format", http.StatusBadRequest)
		return
	}

	var article models.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		pkg.Logger.Errorw("Failed to decode request body", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateArticle(uint(id), &article); err != nil {
		pkg.Logger.Errorw("Failed to update article", "id", id, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ArticleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		pkg.Logger.Errorw("ID is required", "error", "missing ID parameter")
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.Logger.Errorw("Invalid ID format", "error", err)
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteArticle(uint(id)); err != nil {
		pkg.Logger.Errorw("Failed to delete article", "id", id, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
