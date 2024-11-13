package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "example.com/receipt-processor/internal/api"
    "example.com/receipt-processor/internal/models"
    "github.com/stretchr/testify/assert"
)

func TestGetPointsEndpoint(t *testing.T) {
    const (
        colorRed    = "\033[31m"
        colorGreen  = "\033[32m"
        colorYellow = "\033[33m"
        colorBlue   = "\033[34m"
        colorPurple = "\033[35m"
        colorCyan   = "\033[36m"
        colorReset  = "\033[0m"
    )

    r := api.SetupRouter()

    receipt := models.Receipt{
        Retailer:     "Target",
        PurchaseDate: "2022-01-01",
        PurchaseTime: "13:01",
        Items: []models.Item{
            {
                ShortDescription: "Mountain Dew 12PK",
                Price:           "6.49",
            },
            {
                ShortDescription: "Emils Cheese Pizza",
                Price: "12.25",
            },
            {
                ShortDescription: "Knorr Creamy Chicken",
                Price: "1.26",
            },
            {
                ShortDescription: "Doritos Nacho Cheese",
                Price: "3.35",
            },
            {
                ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
                Price: "12.00",
            },
        },
        Total: "35.35",
    }

    receiptJSON, _ := json.MarshalIndent(receipt, "", "    ")
    t.Logf("%sTest Receipt:%s\n%s%s%s", 
        colorBlue, 
        colorReset,
        colorCyan, 
        string(receiptJSON), 
        colorReset)

    // Process the receipt
    jsonValue, _ := json.Marshal(receipt)
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)

    var processResponse models.Response
    err := json.Unmarshal(w.Body.Bytes(), &processResponse)
    assert.NoError(t, err)
    receiptID := processResponse.ID
    t.Logf("%sReceipt ID:%s %s%s%s",
        colorBlue, 
        colorReset,
        colorGreen, 
        receiptID, 
        colorReset)

    // Get points for the receipt
    w = httptest.NewRecorder()
    req, _ = http.NewRequest("GET", "/receipts/"+receiptID+"/points", nil)
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    var pointsResponse models.PointsResponse
    err = json.Unmarshal(w.Body.Bytes(), &pointsResponse)
    assert.NoError(t, err)
    t.Logf("%sPoints awarded:%s %s%d%s",
        colorYellow, 
        colorReset,
        colorPurple, 
        pointsResponse.Points, 
        colorReset)
    assert.NotNil(t, pointsResponse.Points)
}