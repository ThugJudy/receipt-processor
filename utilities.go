package main

import (
	"math"
	"strconv"
	"strings"
	"time"
)

func isAlphanumeric(r rune) bool {
	return ('A' <= r && r <= 'Z') || ('a' <= r && r <= 'z') || ('0' <= r && r <= '9')
}

func CalculatePoints(receipt Receipt) int {
	points := 0

	// Points for alphanumeric characters in the retailer name
	for _, r := range receipt.Retailer {
		if isAlphanumeric(r) {
			points++
		}
	}

	totalFloat, _ := strconv.ParseFloat(receipt.Total, 64)

	// 50 points if the total is a round dollar amount
	if float64(totalFloat) == math.Floor(totalFloat) {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25
	if math.Mod(totalFloat*100, 25) == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt
	points += (len(receipt.Items) / 2) * 5

	// Points for item descriptions and price calculation
	for _, item := range receipt.Items {
		itemPrice, _ := strconv.ParseFloat(item.Price, 64)

		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			points += int(math.Ceil(itemPrice * 0.2))
		}
	}

	// 6 points if the day in the purchase date is odd
	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 != 0 {
		points += 6
	}

	// 10 points if the time of purchase is between 2:00pm and 4:00pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}
