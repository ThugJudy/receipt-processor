package main

import (
	"sync"
)

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// maps receipt ID to points
type ReceiptStore struct {
	sync.Mutex
	Receipts map[string]int
}

var store = ReceiptStore{
	Receipts: make(map[string]int),
}
