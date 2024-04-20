package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product describe the items
type Product struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"product_name" bson:"product_name"`
	Price       int                `json:"price" bson:"price"`
	Currency    string             `json:"currency" bson:"currency"`
	Quantity    string             `json:"quantity" bson:"quantity"`
	Discount    int                `json:"discount,omitempty" bson:"discount,omitempty"`
	Vendor      string             `json:"vendor" bson:"vendor"`
	Accessories []string           `json:"accessories,omitempty" bson:"accessories,omitempty"`
	SkuID       string             `json:"sku_id" bson:"sku_id"`
}

// Iphone10 Sample data
var Iphone10 = Product{
	ID:          primitive.NewObjectID(),
	Name:        "Iphone14",
	Price:       850,
	Currency:    "USD",
	Quantity:    "40",
	Discount:    0,
	Vendor:      "apple",
	Accessories: []string{"charger", "headset"},
	SkuID:       "1234",
}

// Galaxy Sample data
var Galaxy = Product{
	ID:          primitive.NewObjectID(),
	Name:        "Galaxy",
	Price:       550,
	Currency:    "USD",
	Quantity:    "30",
	Discount:    0,
	Vendor:      "samasung",
	Accessories: []string{"charger", "headset"},
	SkuID:       "12",
}
