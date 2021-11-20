package main

type FilmDAO interface {
	IsFilmExists(id int) bool                                               // DONE
	AddFilm(film Film)                                                      // DONE
	UpdateFilm(film Film)                                                   // FIX-LATER
	DeleteFilm(id int)                                                      // DONE
	GetAllFilms(limit int) []Film                                           // DONE
	GetFilmById(id int)                                                     // DONE
	GetAllFilmsOrderedByCategoryDesc(category string) []Film                // DONE
	GetAllFilmsOrderedByCategoryAsc(category string) []Film                 // DONE
	GetFilmsOrderedByLimitSkip(category string, limit int, skip int) []Film // DONE
}
