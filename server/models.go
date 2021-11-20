package main

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type App struct {
	r *mux.Router
}

type Film struct {
	FilmId            int           `json:"film_id"`
	Title             string        `json:"title"`
	Description       string        `json:"description"`
	ReleaseYear       string        `json:"release_year"`
	LanguageId        int           `json:"language_id"`
	OriginalLanguagId sql.NullInt32 `json:"original_language_id"`
	RentalDuration    string        `json:"rental_duration"`
	RenatalRate       float32       `json:"rental_rate"`
	Length            string        `json:"length"`
	ReplacementCost   float32       `json:"replacement_cost"`
	Rating            string        `json:"rating"`
	SpecialFeatures   string        `json:"special_features"`
	LastUpdate        string        `json:"last_update"`
}

type Web struct {
	Films      []Film `json:"items"`
	TotalCount int    `json:"total_count"`
}
