package api

import (
    "example.com/receipt-processor/internal/models"
    "example.com/receipt-processor/internal/service"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()
    store := models.NewReceiptStore()
    pointsCalculator := service.NewPointsCalculator()

    router.POST("/receipts/process", func(c *gin.Context) {
        var receipt models.Receipt

        if err := c.ShouldBindJSON(&receipt); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt format"})
            return
        }

        if receipt.Retailer == "" || receipt.PurchaseDate == "" || receipt.PurchaseTime == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
            return
        }

        id := uuid.New().String()
        store.Save(id, receipt)
        c.JSON(http.StatusOK, models.Response{ID: id})
    })

    router.GET("/receipts/:id/points", func(c *gin.Context) {
        id := c.Param("id")

        receipt, exists := store.Get(id)
        if !exists {
            c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that id"})
            return
        }

        points := pointsCalculator.CalculatePoints(receipt)
        c.JSON(http.StatusOK, models.PointsResponse{Points: points})
    })

    return router
}