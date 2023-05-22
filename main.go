package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

func JSONError(w http.ResponseWriter, err Error, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

func Sum(x int, y int) int {
	return x + y
}

func SumHandler(w http.ResponseWriter, r *http.Request) {
	x, err := strconv.Atoi((r.URL.Query().Get("x")))
	if err != nil {
		JSONError(w, Error{Code: 400, Msg: "Invalid x number"}, 400)
		return
	}

	y, err := strconv.Atoi(r.URL.Query().Get("y"))
	if err != nil {
		JSONError(w, Error{Code: 400, Msg: "Invalid y number"}, 400)
		return
	}

	json.NewEncoder(w).Encode(x + y)
}

func main() {

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	http.HandleFunc("/calculate", SumHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))

}
