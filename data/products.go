package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = Products{
	&Product{1, "Latte", "Frothy milky coffee", 2.45, "abc123", time.Now().UTC().String(), time.Now().UTC().String(), time.Now().UTC().String()},
	&Product{2, "Espresso", "Short and strong coffee without milks", 1.99, "fjd34", time.Now().UTC().String(), time.Now().UTC().String(), time.Now().UTC().String()},
}
