package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Branch struct {
	ID   string    `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID       string    `json:"id"`
	BranchId string    `json:"branchId"`
	Name     string `json:"name"`

	Branch *Branch `json:"Branch"`
}

type Product struct {
	ID          string    `json:"id"`
	CategoryId  string    `json:"categoryId"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`

	Category *Category `json:"Category"`
}

var productMap map[string]Product
// var products []Product
var categories map[string]Category
var branches map[string]Branch


// Функция создания Branch (ветки)
func CreateBranch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var branch Branch
	json.NewDecoder(r.Body).Decode(&branch)
	branch.ID = uuid.New().String()
	branches[branch.ID] = branch
	json.NewEncoder(w).Encode(branch)

	defer r.Body.Close()
}

// Функция создания Category (категории)
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var category Category
	json.NewDecoder(r.Body).Decode(&category)

	// проверка существует ли branch с ID
	branch, exist := branches[category.BranchId]
	if !exist{
		http.Error(w, "Branch с ID "+category.BranchId+"не найдено", http.StatusBadRequest)
		return
	}

	category.Branch = &branch
	category.ID = uuid.New().String()
	categories[category.ID] = category

	// w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)

	defer r.Body.Close()
}

// Функция создания Product
func CreatProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	json.NewDecoder(r.Body).Decode(&product)

	category, exist := categories[product.CategoryId]
	if !exist{
		http.Error(w, "Category с ID "+product.CategoryId+"не найдено", http.StatusBadRequest)
		return
	}


	product.ID = uuid.New().String()
	product.Category = &category
	productMap[product.ID] = product
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)

	defer r.Body.Close()
}


func getBranch(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	branch, exist := branches[id]
	if !exist{
		http.Error(w, "Branch с ID "+id+"не найдено", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(branch)
	json.NewEncoder(w).Encode(branches)

}

func getCatagory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	category, exist := categories[id]
	if !exist{
		http.Error(w, "Category с ID "+id+"не найдено", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(category)
	json.NewEncoder(w).Encode(categories)
}

// Функция getProduct будет показывать книгу по опрделенному ID
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	product, exist := categories[id]
	if !exist{
		http.Error(w, "Product с ID "+id+"не найдено", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(product)
	json.NewEncoder(w).Encode(productMap)
}


func UpdateBranch(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	
	var updateBranch Branch
	if err := json.NewDecoder(r.Body).Decode(&updateBranch); err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid json"})
		return
	} 
	_, exist := branches[id]
	if !exist{
		http.Error(w, "Branch с ID "+id+"не найдено", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(updateBranch)
	
}

func UpdateCategory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	var updateCategory Category
	if err := json.NewDecoder(r.Body).Decode(&updateCategory); err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid json"})
		return
	} 
	
	_, exist := categories[id]
	if !exist{
		http.Error(w, "Category с ID "+id+"не найдено", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(updateCategory)
			
}


func UpdateProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	var updateProduct Product
	if err := json.NewDecoder(r.Body).Decode(&updateProduct); err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid json"})
		return
	} 
	
	_, exist := categories[id]
	if !exist{
		http.Error(w, "Product с ID "+id+"не найдено", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(updateProduct)
	
}

func DeleteBranch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	id := mux.Vars(r)["id"]
	delete(branches, id)

}

// Удаленеие категории
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	id := mux.Vars(r)["id"]
	delete(categories, id)

}

// Удаление продукта по ID 
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	id := (mux.Vars(r)["id"])
	delete(productMap, id)
	
}	




func main() {
	branches = make(map[string]Branch)
	categories = make(map[string]Category)
	productMap = make(map[string]Product)
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
	r.HandleFunc("/categories/{id}", DeleteCategory).Methods("DELETE")
	r.HandleFunc("/products/{id}", DeleteCategory).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
