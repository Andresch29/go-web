package models

type Product struct {
	Id 			int
	Name 		string
	Quantity 	int
	CodeValue 	string
	IsPublished bool
	Expiration 	string
	Price 		float64
}

type Products []Product