package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProcessReceiptResponse struct {
	ID string `json:"id"`
}

type PointsResponse struct {
	Points int `json:"points"`
}

// ProcessReceipt handles the receipt processing and assigns an ID
func ProcessReceipt(c echo.Context) error {
	receipt := new(Receipt)
	if err := c.Bind(receipt); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid receipt format"})
	}

	// Validate the receipt
	if err := receipt.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// Generate a unique ID
	id := GenerateUUID()

	// Calculate points
	points := CalculatePoints(receipt)

	// Save receipt and points to the database
	if err := SaveReceipt(id, points); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to save receipt"})
	}

	return c.JSON(http.StatusOK, ProcessReceiptResponse{ID: id})
}

// GetPoints returns the points for a specific receipt ID
func GetPoints(c echo.Context) error {
	id := c.Param("id")

	// Retrieve points from the database
	points, err := GetPointsByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "receipt not found"})
	}

	return c.JSON(http.StatusOK, PointsResponse{Points: points})
}
