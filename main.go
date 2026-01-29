package main

import (
	"encoding/json"
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
	branchExist := false
	var foundBranch *Branch
	for i := range branchs {
		if branchs[i].ID == category.BranchId {
			branchExist = true
			foundBranch = &branchs[i]
			break
		}
	}

	if !branchExist {
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

	categoryExist := false
	for _, category := range categorys {
		if category.ID == product.ID {
			categoryExist = true
			break
		}
	}
	if !categoryExist {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Категории с ID нет"})
	}

	product.ID = (rand.Intn(1000000))
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
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

func getCatagory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, item := range categorys {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// Функция CreatProduc будет создавать книгу

// func UpdateProduct(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range products{
// 		if item.ID == params["id"]{
// 			products = append(products[:index], products[index+1:]...)
// 			var product  Product
// 			_ = json.NewDecoder(r.Body).Decode(&product)
// 			product.ID = params["id"]
// 			products = append(products, product)
// 			json.NewEncoder(w).Encode(product)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(products)
// }

// func DeleteProduct(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range products{
// 		if item.ID == params["id"]{
// 			products = append(products[:index], products[index+1:]...)
// 			break

// 		}
// 	}
// 	json.NewEncoder(w).Encode(products)
// }

func main() {
	r := mux.NewRouter()

	products = append(products, Product{ID: 1, CategoryId: 2, Name: "iphone", Price: 5000, Description: "Small", Category: &Category{ID: 2, BranchId: 5, Name: "Phone", Branch: &Branch{ID: 5, Name: "main"}}})
	products = append(products, Product{ID: 2, CategoryId: 2, Name: "iphone2", Price: 5000, Description: "Big", Category: &Category{ID: 2, BranchId: 5, Name: "Phone", Branch: &Branch{ID: 5, Name: "main"}}})
	r.HandleFunc("/products", getProduct).Methods("GET")
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/branch", CreateBranch).Methods("POST")
	r.HandleFunc("/category", CreateCategory).Methods("POST")
	// r.HandleFunc("/products", CreatProduct).Methods("POST")
	// r.HandleFunc("/products/{id}", UpdateProduct).Methods("PUT")
	// r.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
