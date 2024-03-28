package models

import "time"

type Transaction struct {
	ID           int64     `json:"id"`
	Type         string    `json:"type"`
	Name         string    `json:"name"`
	Category     string    `json:"category"`
	Description  string    `json:"description"`
	Amount       float64   `json:"amount"`
	Date         time.Time `json:"date"`
	RegistryTime time.Time `json:"registryTime"`
}
