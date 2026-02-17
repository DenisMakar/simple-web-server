package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Branch struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID       int    `json:"id"`
	BranchId int    `json:"branchId"`
	Name     string `json:"name"`

	Branch *Branch `json:"BranchId"`
}

type Product struct {
	ID          int    `json:"id"`
	CategoryId  int    `json:"categoryId"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`

	Category *Category `json:"CategoryId"`
}

var products []Product
var categorys []Category
var branchs []Branch

// Функция создания Branch (ветки)
func CreateBranch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var branch Branch
	json.NewDecoder(r.Body).Decode(&branch)
	branch.ID = rand.Intn(1000000)
	branchs = append(branchs, branch)
	json.NewEncoder(w).Encode(branch)
}

// Функция создания Category (категории)
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var category Category
	json.NewDecoder(r.Body).Decode(&category)

	// проверка существует ли branch с ID
	var foundBranch *Branch
	for i := range branchs {
		if branchs[i].ID == category.BranchId {
			foundBranch = &branchs[i]
			break
		}
	}
// Проверка существования Branch
	if foundBranch == nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Ветки с ID нет",
		})
		return
	}

	category.ID = rand.Intn(1000000)
	category.Branch = foundBranch
	categorys = append(categorys, category)
	json.NewEncoder(w).Encode(category)

}

// Функция создания Product
func CreatProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	json.NewDecoder(r.Body).Decode(&product)

	
	var foundCategory *Category
	for i:= range categorys {
		if categorys[i].ID == product.CategoryId{
			foundCategory = &categorys[i]
			break
		}
	}

	// Проверка существования Category
	if foundCategory ==  nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Категории с ID нет"})
		return
	}

	product.ID = (rand.Intn(1000000))
	product.Category = foundCategory
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}


func getBranch(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, item := range branchs{
		if item.ID == id{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(branchs)

}

func getCatagory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, item := range categorys{
		if item.ID == id{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(categorys)

}

// Функция getProduct будет показывать книгу по опрделенному ID
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, item := range products {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(products)
}


func UpdateBranch(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	var updateBranch Branch
	if err := json.NewDecoder(r.Body).Decode(&updateBranch); err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid json"})
		return
	} 
	
	found := false
	for i, branch := range branchs{
		if branch.ID == id{
			updateBranch.ID = id

			branchs[i] = updateBranch
			found = true

			json.NewEncoder(w).Encode(updateBranch)
			break
		}
	}

	if !found{
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Branch with ID %d not found", id)})
	}

	
}

func UpdateCategory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	var updateCategory Category
	if err := json.NewDecoder(r.Body).Decode(&updateCategory); err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid json"})
		return
	} 
	
	found := false
	for i, category := range categorys{
		if category.ID == id{
			updateCategory.ID = id

			categorys[i] = updateCategory
			found = true

			json.NewEncoder(w).Encode(updateCategory)
			break
		}
	}

	if !found{
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Category with ID %d not found", id)})
	}

	
}


func UpdateProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	var updateProduct Product
	if err := json.NewDecoder(r.Body).Decode(&updateProduct); err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid json"})
		return
	} 
	
	found := false
	for i, product := range products{
		if product.ID == id{
			updateProduct.ID = id

			products[i] = updateProduct
			found = true

			json.NewEncoder(w).Encode(updateProduct)
			break
		}
	}

	if !found{
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Product with ID %d not found", id)})
	}

	
}

func DeleteBranch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}
	
	// Проверяем существование ветки
	found := false
	for _, b := range branchs {
		if b.ID == id {
			found = true
			break
		}
	}
	
	if !found {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Branch not found"})
		return
	}
	
	// 1. Удаляем продукты
	for i := len(products) - 1; i >= 0; i-- {
		for _, cat := range categorys {
			if cat.ID == products[i].CategoryId && cat.BranchId == id {
				products = append(products[:i], products[i+1:]...)
				break
			}
		}
	}
	
	// 2. Удаляем категории  
	for i := len(categorys) - 1; i >= 0; i-- {
		if categorys[i].BranchId == id {
			categorys = append(categorys[:i], categorys[i+1:]...)
		}
	}
	
	// 3. Удаляем ветку
	for i := len(branchs) - 1; i >= 0; i-- {
		if branchs[i].ID == id {
			branchs = append(branchs[:i], branchs[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Branch and all related data deleted",
			})
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/branch", getBranch).Methods("GET")
	r.HandleFunc("/branch/{id}", getBranch).Methods("GET")
	r.HandleFunc("/category", getCatagory).Methods("GET")
	r.HandleFunc("/category/{id}", getCatagory).Methods("GET")
	r.HandleFunc("/products", getProduct).Methods("GET")
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/branch/{id}/category", getBranch).Methods("GET")
	r.HandleFunc("/categories/{id}/products", getCatagory).Methods("GET")

	r.HandleFunc("/branch", CreateBranch).Methods("POST")
	r.HandleFunc("/category", CreateCategory).Methods("POST")
	r.HandleFunc("/products", CreatProduct).Methods("POST")

	r.HandleFunc("/branch/{id}", UpdateBranch).Methods("PUT")
	r.HandleFunc("/categories/{id}", UpdateCategory).Methods("PUT")
	r.HandleFunc("/products/{id}", UpdateProduct).Methods("PUT")
	

	r.HandleFunc("/branch/{id}", DeleteBranch).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Старт программы")
}
