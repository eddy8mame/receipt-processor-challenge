package service

import (
    "math"
    "strconv"
    "strings"
    "unicode"
    "example.com/receipt-processor/internal/models"
)

type PointsCalculator struct{}

func NewPointsCalculator() *PointsCalculator {
    return &PointsCalculator{}
}

func (pc *PointsCalculator) CalculatePoints(receipt models.Receipt) int {
    return pc.retailerNamePoints(receipt.Retailer) +
        pc.roundDollarPoints(receipt.Total) +
        pc.quarterMultiplePoints(receipt.Total) +
        pc.itemCountPoints(receipt.Items) +
        pc.itemDescriptionPoints(receipt.Items) +
        pc.oddDayPoints(receipt.PurchaseDate) +
        pc.timeOfDayPoints(receipt.PurchaseTime)
}

// Rule 1: One point for every alphanumeric character in the retailer name
func (pc *PointsCalculator) retailerNamePoints(retailer string) int {
    points := 0
    for _, char := range retailer {
        if unicode.IsLetter(char) || unicode.IsNumber(char) {
            points++
        }
    }
    return points
}

// Rule 2: 50 points if the total is a round dollar amount with no cents
func (pc *PointsCalculator) roundDollarPoints(total string) int {
    amount, err := strconv.ParseFloat(total, 64)
    if err != nil {
        return 0
    }

    if amount == float64(int(amount)) {
        return 50
    }
    return 0
}

// Rule 3: 25 points if the total is a multiple of 0.25
func (pc *PointsCalculator) quarterMultiplePoints(total string) int {
    amount, err := strconv.ParseFloat(total, 64)
    if err != nil {
        return 0
    }

    cents := int(math.Round(amount * 100))
    if cents%25 == 0 {
        return 25
    }
    return 0
}

// Rule 4: 5 points for every two items on the receipt
func (pc *PointsCalculator) itemCountPoints(items []models.Item) int {
    return (len(items) / 2) * 5
}

// Rule 5: Points based on item description length and price
func (pc *PointsCalculator) itemDescriptionPoints(items []models.Item) int {
    points := 0
    for _, item := range items {
        trimmedLen := len(strings.TrimSpace(item.ShortDescription))
        if trimmedLen%3 == 0 {
            price, err := strconv.ParseFloat(item.Price, 64)
            if err != nil {
                continue
            }
            points += int(math.Ceil(price * 0.2))
        }
    }
    return points
}

// Rule 6: 6 points if the day in the purchase date is odd
func (pc *PointsCalculator) oddDayPoints(purchaseDate string) int {
    parts := strings.Split(purchaseDate, "-")
    if len(parts) != 3 {
        return 0
    }

    day, err := strconv.Atoi(parts[2])
    if err != nil {
        return 0
    }

    if day%2 == 1 {
        return 6
    }
    return 0
}

// Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm
func (pc *PointsCalculator) timeOfDayPoints(purchaseTime string) int {
    parts := strings.Split(purchaseTime, ":")
    if len(parts) != 2 {
        return 0
    }

    hour, err := strconv.Atoi(parts[0])
    if err != nil {
        return 0
    }

    minute, err := strconv.Atoi(parts[1])
    if err != nil {
        return 0
    }

    if (hour == 14 && minute > 0) || hour == 15 {
        return 10
    }
    return 0
}