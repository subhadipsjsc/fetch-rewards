package main

import (
	"errors"
	"regexp"
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

func (r *Receipt) Validate() error {
	if r.Retailer == "" || !regexp.MustCompile(`^[\w\s\-&]+$`).MatchString(r.Retailer) {
		return errors.New("invalid retailer name")
	}
	if len(r.Items) == 0 {
		return errors.New("receipt must have at least one item")
	}
	if !regexp.MustCompile(`^\d+\.\d{2}$`).MatchString(r.Total) {
		return errors.New("invalid total format")
	}
	return nil
}
