// Code is from https://gist.github.com/chadlung/3b4a12e03ce721632019.  Using to understand and keep reference
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Movie Struct
type Movie struct {
	Title  string `json:"title"`
	Rating string `json:"rating"`
	Year   string `json:"year"`
}

var movies = map[string]*Movie{
	"tt0076759": &Movie{Title: "Star Wars: A New Hope", Rating: "8.7", Year: "1977"},
	"tt0082971": &Movie{Title: "Indiana Jones: Raiders of the Lost Ark", Rating: "8.6", Year: "1981"},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", handleMovies).Methods("GET")
	router.HandleFunc("/movie/{imdbKey}", getMovie).Methods("GET")
	router.HandleFunc("/movie/{imdbKey}", deleteMovie).Methods("DELETE")
	router.HandleFunc("/movie/{imdbKey}", addMovie).Methods("POST")
	http.ListenAndServe(":8080", router)
}

func findMovie(res http.ResponseWriter, req *http.Request) (*Movie, string) {
	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	imdbKey := vars["imdbKey"]

	log.Println("Request for:", imdbKey)

	movie, ok := movies[imdbKey]
	if !ok {
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprint(res, string("Movie not found"))
	}
	return movie, imdbKey
}

func getMovie(res http.ResponseWriter, req *http.Request) {
	movie, _ := findMovie(res, req)
	outgoingJSON, error := json.Marshal(movie)
	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(res, string(outgoingJSON))
}

func deleteMovie(res http.ResponseWriter, req *http.Request) {
	_, imdbKey := findMovie(res, req)
	delete(movies, imdbKey)
	res.WriteHeader(http.StatusNoContent)
}

func handleMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	outgoingJSON, error := json.Marshal(movies)

	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(res, string(outgoingJSON))
}

func addMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	movie := new(Movie)
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&movie); err != nil {
		log.Println(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(req)
	imdbKey := vars["imdbKey"]

	movies[imdbKey] = movie
	outgoingJSON, err := json.Marshal(movie)
	if err != nil {
		log.Println(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, string(outgoingJSON))
}
