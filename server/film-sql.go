package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

// CHECK IS FILM EXISTS
func (f *Film) IsFilmExists(id int) bool {

	con := GetDBConnection()
	defer con.Close()

	sqlStatement := `SELECT * FROM film WHERE film_id = ?`

	var film Film

	row := con.QueryRow(sqlStatement, id)

	err := row.Scan(&film.FilmId, &film.Title, &film.Description,
		&film.ReleaseYear, &film.LanguageId, &film.OriginalLanguagId,
		&film.RentalDuration, &film.RenatalRate, &film.Length,
		&film.ReplacementCost, &film.Rating, &film.SpecialFeatures, &film.LastUpdate)

	if err == sql.ErrNoRows {
		log.Println("NO Such film")
		return false
	} else if err != nil {
		log.Println(err)
	}

	log.Println("Film exits")
	return true

}

// ADD FILM
func (f *Film) AddFilm(film Film) error {
	con := GetDBConnection()
	defer con.Close()

	// Prepare the Context Runtime to Prevent Crashing and Handle Errors
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	// SQL Statement to execute
	sqlStatement := "INSERT INTO film(title, description," +
		"release_year, language_id, original_language_id, rental_duration, " +
		"rental_rate, length, replacement_cost, rating)" +
		"VALUES(?, ?, ?, ?, ? ,? , ?, ?, ?, ?)"

	// Prepare Context check SQL statment is valid
	stmt, err := con.PrepareContext(ctx, sqlStatement)

	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	// Execute the SQL statment
	res, err := stmt.ExecContext(ctx, film.Title,
		film.Description, film.ReleaseYear, film.LanguageId,
		film.OriginalLanguagId, film.RentalDuration, film.RenatalRate,
		film.Length, film.ReplacementCost, film.Rating)

	if err != nil {
		log.Printf("Error %s when inserting row into films table", err)
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}

	log.Printf("%d film created ", rows)
	return nil
}

// DELETE FILM BY ID
func (f *Film) DeleteFilm(id int) error {
	con := GetDBConnection()
	defer con.Close()

	// Prepare the Context Runtime to Prevent Crashing and Handle Errors
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	sqlStatement := `DELETE FROM film WHERE film_id = ?`

	// Prepare Context check SQL statment is valid
	stmt, err := con.PrepareContext(ctx, sqlStatement)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		log.Printf("Error %s when inserting row into films table", err)
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}

	log.Printf("%d film delete ", rows)
	return nil
}

// GET FILM BY ID
func (f *Film) GetFilmById(id int) Film {

	con := GetDBConnection()
	defer con.Close()

	sqlStatement := `SELECT * FROM film WHERE film_id = ?`

	var film Film

	row := con.QueryRow(sqlStatement, id)

	switch err := row.Scan(&film.FilmId, &film.Title, &film.Description,
		&film.ReleaseYear, &film.LanguageId, &film.OriginalLanguagId,
		&film.RentalDuration, &film.RenatalRate, &film.Length,
		&film.ReplacementCost, &film.Rating, &film.SpecialFeatures, &film.LastUpdate); err {

	case sql.ErrNoRows:
		log.Println("No rows were returned!")
	case nil:
		log.Println("Get Film by Id Operearion was Succsusful")
	default:
		log.Println(err)
	}
	return film
}

// GET ALL FILMS WITH LIMIT
func (f *Film) GetAllFilms(limit int) ([]Film, error) {
	con := GetDBConnection()
	defer con.Close()

	var film Film
	var allFilms []Film

	sqlStatement := `SELECT * FROM film limit ?`
	q, err := con.Query(sqlStatement, limit)

	if err != nil {
		return allFilms, err
	}
	for q.Next() {

		err = q.Scan(&film.FilmId, &film.Title, &film.Description,
			&film.ReleaseYear, &film.LanguageId, &film.OriginalLanguagId,
			&film.RentalDuration, &film.RenatalRate, &film.Length,
			&film.ReplacementCost, &film.Rating, &film.SpecialFeatures, &film.LastUpdate)

		if err != nil {
			return allFilms, err
		}
		allFilms = append(allFilms, film)
	}

	return allFilms, nil
}

//TODO: func update operation not working
func (f *Film) UpdateFilm(film Film) error {
	con := GetDBConnection()
	defer con.Close()

	// Prepare the Context Runtime to Prevent Crashing and Handle Errors
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	// SQL Statement to execute
	sqlStatement := "UPDATE film set title = ?, description = ?," +
		"release_year = ?, language_id = ?, original_language_id = ?, rental_duration = ?, " +
		"rental_rate = ?, length = ?, replacement_cost = ?, rating = ?  WHERE film_id = ?"

	// Prepare Context check SQL statment is valid
	stmt, err := con.PrepareContext(ctx, sqlStatement)

	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	// Execute the SQL statment
	res, err := stmt.ExecContext(ctx, film.FilmId, film.Title,
		film.Description, film.ReleaseYear, film.LanguageId,
		film.OriginalLanguagId, film.RentalDuration, film.RenatalRate,
		film.Length, film.ReplacementCost, film.Rating)

	if err != nil {
		log.Printf("Error %s when updating row into films table", err)
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when updating rows affected", err)
		return err
	}

	log.Printf("%d film updated ", rows)
	return nil
}

func (f *Film) GetAllFilmsOrderedByCategoryDesc(category string) []Film {
	con := GetDBConnection()
	defer con.Close()

	var film Film
	var allFilms []Film

	sqlStatement := fmt.Sprintf("SELECT * FROM film ORDER BY %s DESC LIMIT 50", category)
	q, err := con.Query(sqlStatement)

	errorHandler(err)

	for q.Next() {

		err = q.Scan(&film.FilmId, &film.Title, &film.Description,
			&film.ReleaseYear, &film.LanguageId, &film.OriginalLanguagId,
			&film.RentalDuration, &film.RenatalRate, &film.Length,
			&film.ReplacementCost, &film.Rating, &film.SpecialFeatures, &film.LastUpdate)

		errorHandler(err)

		allFilms = append(allFilms, film)
	}

	return allFilms
}

func (f *Film) GetAllFilmsOrderedByCategoryAsc(category string) []Film {
	con := GetDBConnection()
	defer con.Close()

	var film Film
	var allFilms []Film

	sqlStatement := fmt.Sprintf("SELECT * FROM film ORDER BY %s ASC LIMIT 50", category)
	q, err := con.Query(sqlStatement)

	errorHandler(err)

	for q.Next() {

		err = q.Scan(&film.FilmId, &film.Title, &film.Description,
			&film.ReleaseYear, &film.LanguageId, &film.OriginalLanguagId,
			&film.RentalDuration, &film.RenatalRate, &film.Length,
			&film.ReplacementCost, &film.Rating, &film.SpecialFeatures, &film.LastUpdate)

		errorHandler(err)

		allFilms = append(allFilms, film)
	}

	return allFilms
}

func (f *Film) GetFilmsOrderedByLimitSkip(category string, order string, limit int, skip int) ([]Film, error) {
	con := GetDBConnection()
	defer con.Close()

	var film Film
	var allFilms []Film
	// log.Printf("GetFilmsOrderedByLimitSkip=category=%s orderby=%s limit=%d ofset=%d", category, order, limit, skip)
	sqlStatement := fmt.Sprintf("SELECT * FROM film ORDER BY %s %s LIMIT %d OFFSET %d", category, order, limit, skip)
	log.Println(sqlStatement)
	q, err := con.Query(sqlStatement)

	if err != nil {
		log.Println("Error: ", err)
		return allFilms, err
	}

	for q.Next() {

		err = q.Scan(&film.FilmId, &film.Title, &film.Description,
			&film.ReleaseYear, &film.LanguageId, &film.OriginalLanguagId,
			&film.RentalDuration, &film.RenatalRate, &film.Length,
			&film.ReplacementCost, &film.Rating, &film.SpecialFeatures, &film.LastUpdate)

		if err != nil {
			log.Println("Error: ", err)
			return allFilms, err
		}

		allFilms = append(allFilms, film)
	}

	return allFilms, nil
}

// ERROR HANDLE FUNC
func errorHandler(e error) {
	if e != nil {
		log.Println("Error: ", e)
	}
}
