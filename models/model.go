package models

import "time"

type ResponseAPI struct {
	SearchName      string    `json:"SearchName,omitempty"`
	MetaphoneSearch string    `json:"MetaphoneSearch,omitempty"`
	ID              int       `json:"ID"`
	CreatedAt       time.Time `json:"CreatedAt"`
	UpdatedAt       time.Time `json:"UpdatedAt"`
	DeletedAt       time.Time `json:"DeletedAt"`
	Name            string    `json:"Name"`
	Classification  string    `json:"Classification"`
	Metaphone       string    `json:"Metaphone"`
	NameVariations  []string  `json:"NameVariations"`
}

type WriteStruct struct {
	SearchName     string   `json:"SearchName,omitempty"`
	ID             string   `json:"ID"`
	CreatedAt      string   `json:"CreatedAt"`
	UpdatedAt      string   `json:"UpdatedAt"`
	DeletedAt      string   `json:"DeletedAt"`
	Name           string   `json:"Name"`
	Classification string   `json:"Classification"`
	Metaphone      string   `json:"Metaphone"`
	NameVariations []string `json:"NameVariations"`
	Err            error    `json:"err"`
}
