package models

import (
    "sync"
)

type ReceiptStore struct {
    sync.RWMutex
    receipts map[string]Receipt 
}

type Response struct {
    ID string `json:"id"`
}

type Item struct {
    ShortDescription string `json:"shortDescription"`
    Price            string `json:"price"`
}

type Receipt struct {
    Retailer     string `json:"retailer"`
    PurchaseDate string `json:"purchaseDate"`
    PurchaseTime string `json:"purchaseTime"`
    Items        []Item `json:"items"`
    Total        string `json:"total"`
}

type PointsResponse struct {
    Points int `json:"points"`
}

func NewReceiptStore() *ReceiptStore {
    return &ReceiptStore{
        receipts: make(map[string]Receipt), 
    }
}

func (rs *ReceiptStore) Save(id string, receipt Receipt) {
    rs.Lock()
    defer rs.Unlock()
    rs.receipts[id] = receipt
}


func (rs *ReceiptStore) Get(id string) (Receipt, bool) {
    rs.RLock()
    defer rs.RUnlock()
    receipt, exists := rs.receipts[id]
    return receipt, exists
}