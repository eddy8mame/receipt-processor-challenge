package main

import (
    "example.com/receipt-processor/internal/api"
)

func main() {
    router := api.SetupRouter()
    router.Run(":8080")
}