package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, i := range movies {
		if i.ID == params["id"] {
			json.NewEncoder(w).Encode(i)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "982429462",
		Title: "One Pice",
		Director: &Director{
			FirstName: "Bayu",
			LastName:  "Segara",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	//r.HandleFunc("/movies", createMovie).Methods("POST")
	//r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
