package models

import "errors"

type product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ForeignName string `json:"foreignName"`
}

// Mock data for test purpose, will be fetched from db
var productList = []product{
	{ID: "CLP", Name: "CLIPS", ForeignName: "CLIPS hot galvanized"},
	{ID: "CN0017", Name: "Concertina", ForeignName: "Barbed Tape Concertina BTO 22 - Profil Ultra - 700mm - HDG Z180"},
	{ID: " EC0017A", Name: "Lama dentata per nastratrice - singolo bisello - passo 2,5 mm - 3 asole (12x5)", ForeignName: "Tape sealer - pitch 2,5 mm."},
}

// Return all products
func GetAllProducts() []product {
	return productList
}

// Return selected (by id) product
func GetProductById(id string) (*product, error) {
	for _, a := range productList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Product not found")
}
