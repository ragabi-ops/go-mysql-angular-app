package main

import (
	"fmt"
	"log"
)

func CheckCode() {
	// -------------------
	// INSERT FILM
	// var film Film = Film{
	// 	FilmId:            1004,
	// 	Title:             "Avengers End Game",
	// 	Description:       sql.NullString{String: "Go", Valid: true},
	// 	ReleaseYear:       sql.NullString{String: "2019", Valid: true},
	// 	LanguageId:        1,
	// 	OriginalLanguagId: sql.NullInt32{},
	// 	RentalDuration:    "6",
	// 	RenatalRate:       9.99,
	// 	Length:            sql.NullString{String: "182", Valid: true},
	// 	ReplacementCost:   5.99,
	// 	Rating:            sql.NullString{},
	// }

	// err := film.UpdateFilm(film)

	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// -------------------

	// -------------------
	//  DELETE FILM BY ID
	// var film Film
	// err := film.DeleteFilm(1002)

	// if err != nil {
	// 	log.Println(err)
	// }

	// -------------------

	// GET FILM BY ID
	// var f Film
	// film := f.GetFilmById(2)
	// log.Println(film)

	// -------------------

	// get all films
	var f Film
	l, err := f.GetFilmsOrderedByLimitSkip("title", "ASC", 50, 50)
	if err != nil {
		log.Println(err)
	}
	for _, film := range l {
		fmt.Println(film)
	}
	// -------------------

	// var f Film
	// film := f.IsFilmExists(9999)
	// log.Println(film)

}
