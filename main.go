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
	ID      int    `json:"id"`
	BranchId int    `json:"branchId"`
	Name string `json:"name"`

	Branch *Branch `json:"BranchId"`
}

type Product struct {
	ID          string `json:"id"`
	CategoryId  int    `json:"categoryId"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`

	Category *Category `json:"CategoryId"`
}

var products []Product

// Функция getProduct будет показывать книгу по опрделенному ID
func getProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// Функция CreatProduc будет создавать книгу
func CreatProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var product Product 
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = strconv.Itoa(rand.Intn(1000000))
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range products{
		if item.ID == params["id"]{
			products = append(products[:index], products[index+1:]...)
			var product  Product
			_ = json.NewDecoder(r.Body).Decode(&product)
			product.ID = params["id"]
			products = append(products, product)
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	json.NewEncoder(w).Encode(products)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range products{
		if item.ID == params["id"]{
			products = append(products[:index], products[index+1:]...)
			break
			
		}
	}
	json.NewEncoder(w).Encode(products)
}

func main(){
	r :=  mux.NewRouter()
	
	products = append(products, Product{ID: "1", CategoryId: 2, Name: "iphone", Price: 5000, Description: "Small", Category: &Category{ID: 2, BranchId: 5, Name: "Phone", Branch: &Branch{ID: 5, Name: "main"}}})
	r.HandleFunc("/products", getProduct).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}	