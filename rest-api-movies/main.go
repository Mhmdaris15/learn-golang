package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	Year     int64     `json:"year"`
	Director *Director `json:"director"`
}

type Director struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

var movies = []Movie{
	{ID: 1, Title: "The Shawshank Redemption", Year: 1994, Director: &Director{ID: 1, Name: "Frank Darabont"}},
	{ID: 2, Title: "The Godfather", Year: 1972, Director: &Director{ID: 2, Name: "Francis Ford Coppola"}},
	{ID: 3, Title: "The Godfather: Part II", Year: 1974, Director: &Director{ID: 2, Name: "Francis Ford Coppola"}},
	{ID: 4, Title: "The Dark Knight", Year: 2008, Director: &Director{ID: 3, Name: "Christopher Nolan"}},
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getMovies")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
		Data interface{} `json:"data"`
		Status int `json:"status"`
	}{
		Message: "Data retrieved successfully",
		Data: movies,
		Status: http.StatusOK,
	})
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteMovie")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		// parser ID data type into integer
		paramsID, _ := strconv.Atoi(params["id"])
		if movie.ID == int64(paramsID) {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
			Data interface{} `json:"data"`
			Status int `json:"status"`
		}{
			Message: "Data deleted successfully",
			Data: movies,
			Status: http.StatusOK,
		},)
			return
		}
	}
	// Send Message to server, that the data is not found
	
	json.NewEncoder(w).Encode(
		struct {
			Message string `json:"message"`
			Data interface{} `json:"data"`
			Status int `json:"status"`
		}{
			Message: "Data not found",
			Data: movies,
			Status: http.StatusNotFound,
		},
	)
}


func getMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getMovie")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		// parser ID data type into integer
		paramsID, _ := strconv.Atoi(params["id"])
		if movie.ID == int64(paramsID) {
			json.NewEncoder(w).Encode(
				struct {
					Message string `json:"message"`
					Data interface{} `json:"data"`
					Status int `json:"status"`
				}{
					Message: "Data found successfully",
					Data: movies,
					Status: http.StatusOK,
				})
			return
		}
	}
	json.NewEncoder(w).Encode(
		struct {
			Message string `json:"message"`
			Data interface{} `json:"data"`
			Status int `json:"status"`
		}{
			Message: "Data not found",
			Data: movies,
			Status: http.StatusNotFound,
		},
	)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createMovie")
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) // get body request and decode to movie struct
	movie.ID = rand.Int63n(100000000)         // generate random ID
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
			Data interface{} `json:"data"`
			Status int `json:"status"`
		}{
			Message: "Data Created successfully",
			Data: movies,
			Status: http.StatusOK,
		},)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateMovie")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) // get body request and decode to movie struct
	for index, item := range movies {
		// parser ID data type into integer
		paramsID, _ := strconv.Atoi(params["id"])
		if item.ID == int64(paramsID) {
			movies = append(movies[:index], movies[index+1:]...)
			movie.ID = item.ID
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(
				struct {
					Message string `json:"message"`
					Data interface{} `json:"data"`
					Status int `json:"status"`
				}{
					Message: "Data updated successfully",
					Data: movies,
					Status: http.StatusOK,
				},
			)
			return
		}
	}
	json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
			Data interface{} `json:"data"`
			Status int `json:"status"`
		}{
			Message: "Data not found",
			Data: movies,
			Status: http.StatusNotFound,
		},)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Server started on port %d\n", 3000)
	}

}