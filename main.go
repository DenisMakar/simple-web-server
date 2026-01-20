package main

type Branch struct{
	id int `json: "id"`
	name string `json: "name"`
}

type Category struct{
	id int `json:" id"`
	branchId int `json:" branchId"`  
	name string `json: "name"`

	Branch *Branch `json: "BranchId"`
}

type Product struct{
	id string `json: "id"`
	categoryId int `json: "categoryId"`
	name string `json: "name"`
	price int `json: "price"`
	description string `json: "description"`

	Category *Category `json: "CategoryId`
}