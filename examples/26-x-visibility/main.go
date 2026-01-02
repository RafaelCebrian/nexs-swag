package main

import (
	"encoding/json"
	"net/http"
)

// @title API de Teste de Visibilidade
// @version 1.0
// @description Demonstra separação de endpoints públicos e privados
// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

type Product struct {
	ID    int     `json:"id" example:"1"`
	Name  string  `json:"name" example:"Notebook"`
	Price float64 `json:"price" example:"2500.00"`
}

type User struct {
	ID    int    `json:"id" example:"1"`
	Name  string `json:"name" example:"João Silva"`
	Email string `json:"email" example:"joao@example.com"`
	Role  string `json:"role" example:"admin"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"Error message"`
}

// ListProductsPublic lista produtos (PÚBLICO - sem annotation)
// @Summary Lista produtos
// @Description Retorna lista de produtos disponíveis publicamente
// @Tags products
// @Produce json
// @Success 200 {array} Product
// @Router /products [get]
func ListProductsPublic(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{ID: 1, Name: "Notebook", Price: 2500.00},
		{ID: 2, Name: "Mouse", Price: 50.00},
	}
	json.NewEncoder(w).Encode(products)
}

// GetProductPublic busca produto (PÚBLICO - explícito)
// @Summary Buscar produto
// @Description Retorna detalhes de um produto específico
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} Product
// @Failure 404 {object} ErrorResponse
// @Router /products/{id} [get]
// @x-visibility public
func GetProductPublic(w http.ResponseWriter, r *http.Request) {
	product := Product{ID: 1, Name: "Notebook", Price: 2500.00}
	json.NewEncoder(w).Encode(product)
}

// CreateProduct cria produto (PRIVADO)
// @Summary Criar produto
// @Description Cria um novo produto no sistema (requer autenticação)
// @Tags products
// @Accept json
// @Produce json
// @Param product body Product true "Dados do produto"
// @Success 201 {object} Product
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Security Bearer
// @Router /products [post]
// @x-visibility private
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	json.NewDecoder(r.Body).Decode(&product)
	product.ID = 3
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// UpdateProduct atualiza produto (PRIVADO)
// @Summary Atualizar produto
// @Description Atualiza dados de um produto existente
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body Product true "Dados atualizados"
// @Success 200 {object} Product
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Security Bearer
// @Router /products/{id} [put]
// @x-visibility private
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	json.NewDecoder(r.Body).Decode(&product)
	product.ID = 1
	json.NewEncoder(w).Encode(product)
}

// DeleteProduct deleta produto (PRIVADO)
// @Summary Deletar produto
// @Description Remove um produto do sistema
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 204 "No Content"
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Security Bearer
// @Router /products/{id} [delete]
// @x-visibility private
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// PublicHealth verifica status da API (PÚBLICO)
// @Summary Health check público
// @Description Verifica se a API está respondendo
// @Tags system
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func PublicHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// ListUsers lista usuários (PRIVADO - admin)
// @Summary Listar usuários
// @Description Retorna lista de todos os usuários (apenas admin)
// @Tags admin
// @Produce json
// @Success 200 {array} User
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Security Bearer
// @Router /admin/users [get]
// @x-visibility private
func ListUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "João Silva", Email: "joao@example.com", Role: "admin"},
		{ID: 2, Name: "Maria Santos", Email: "maria@example.com", Role: "user"},
	}
	json.NewEncoder(w).Encode(users)
}

// DeleteUser deleta usuário (PRIVADO - admin)
// @Summary Deletar usuário
// @Description Remove um usuário do sistema (apenas admin)
// @Tags admin
// @Produce json
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Security Bearer
// @Router /admin/users/{id} [delete]
// @x-visibility private
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	// Rotas públicas
	http.HandleFunc("/api/v1/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			ListProductsPublic(w, r)
		} else if r.Method == http.MethodPost {
			CreateProduct(w, r)
		}
	})
	http.HandleFunc("/api/v1/products/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetProductPublic(w, r)
		} else if r.Method == http.MethodPut {
			UpdateProduct(w, r)
		} else if r.Method == http.MethodDelete {
			DeleteProduct(w, r)
		}
	})
	http.HandleFunc("/api/v1/health", PublicHealth)
	http.HandleFunc("/api/v1/admin/users", ListUsers)
	http.HandleFunc("/api/v1/admin/users/", DeleteUser)

	http.ListenAndServe(":8080", nil)
}
