package main

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func CalculatePoints(receipt *Receipt) int {
	points := 0

	// Rule 1: One point per alphanumeric character in retailer name
	for _, char := range receipt.Retailer {
		if isAlphanumeric(char) {
			points++
		}
	}

	// Rule 2: 50 points for round dollar totals
	if isRoundDollar(receipt.Total) {
		points += 50
	}

	// Rule 3: 25 points if total is a multiple of 0.25
	if isMultipleOfQuarter(receipt.Total) {
		points += 25
	}

	// Rule 4: 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Points based on item description length
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price := parsePrice(item.Price)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 6 points for odd purchase day
	if isOddDay(receipt.PurchaseDate) {
		points += 6
	}

	// Rule 7: 10 points for purchases between 2:00 PM and 4:00 PM
	if isAfternoonPurchase(receipt.PurchaseTime) {
		points += 10
	}

	return points
}

func isAlphanumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func isRoundDollar(total string) bool {
	return strings.HasSuffix(total, ".00")
}

func isMultipleOfQuarter(total string) bool {
	price := parsePrice(total)
	return math.Mod(price, 0.25) == 0
}

func parsePrice(price string) float64 {
	value, _ := strconv.ParseFloat(price, 64)
	return value
}

func isOddDay(date string) bool {
	t, _ := time.Parse("2006-01-02", date)
	return t.Day()%2 != 0
}

func isAfternoonPurchase(timeStr string) bool {
	t, _ := time.Parse("15:04", timeStr)
	return t.Hour() >= 14 && t.Hour() < 16
}
