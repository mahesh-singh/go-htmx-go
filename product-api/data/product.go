package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Price      float32 `json:"price"`
	SKU        string  `json:"sku"`
	CreationOn string  `json:"-"`
	UpdatedOn  string  `json:"-"`
	DeletedOn  string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:         1,
		Name:       "Coffee",
		Desc:       "A coffee",
		Price:      23.00,
		SKU:        "abc1234",
		CreationOn: time.Now().UTC().String(),
		UpdatedOn:  time.Now().UTC().String(),
	},
	&Product{
		ID:         2,
		Name:       "Tea",
		Desc:       "A tea",
		Price:      20.00,
		SKU:        "ab123",
		CreationOn: time.Now().UTC().String(),
		UpdatedOn:  time.Now().UTC().String(),
	},
}
