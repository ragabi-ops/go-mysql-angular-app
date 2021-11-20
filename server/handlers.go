package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (f *Film) GetAllFilmsJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	sort, sort_order, page_str := r.URL.Query().Get("sort"), r.URL.Query().Get("sort_order"), r.URL.Query().Get("page")
	page, err := strconv.Atoi(page_str)

	if err != nil {
		log.Println(err)
	}

	if page > 0 {
		page = page * 20
	}
	// log.Printf("GetAllFilmsJson=sort=%s sort_order=%s limit=%d page=%d", sort, sort_order, 20, page)
	allFilms, err := f.GetFilmsOrderedByLimitSkip(sort, sort_order, 20, page)
	web := Web{
		Films:      allFilms,
		TotalCount: 1000,
	}

	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(web)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
