package main

import(
    "fmt"
    "log"
    "encoding/json"
    "math/rand"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
)


type Movie struct {
    ID string `json: "ID"`
    Isbn string `json: "isbn"`
    Title string `json: "title"`
    Director *Director `json: "director"`
}

type Director struct { 

    Firstname string `json: "firstname"`
    Lastname string `json: "lastname"`
}

var movies []Movie 

// How to get movie list
func getMovies(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content - Type", "Application/json")
    json.NewEncoder(w).Encode(movies)
}

//How to delete movies
func deleteMovie(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content - Type", "Application/json")
    
    params := mux.Vars(r)
    for index, item := range movies{

        if item.ID == params["id"]{
            movies = append(movies[:index], movies[index+1:]...)
            break
        }       
    }
    json.NewEncoder(w).Encode(movies)
}

//How to get particular movie
func getMovie(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content - Type", "Application/json")
    params := mux.Vars(r)

    // getting back one particular movie instead of all
    for _, item  := range movies{
        if item.ID == params["id"]{
            json.NewEncoder(w).Encode(item)
            return
        }   
    }
}

//How to add movie 
func createMovie(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content - Type", "Application/json")
    var movie Movie
    _ = json.NewDecoder(r.Body).Decode(&movie)
    movie.ID = strconv.Itoa(rand.Intn(100000000))
    movies = append(movies, movie)
    json.NewEncoder(w).Encode(movie)
}

//How to update Movie
func updateMovie(w http.ResponseWriter, r *http.Request){

    //set json content type
    w.Header().Set("Content Type", "Application/json")
    //params
    params := mux.Vars(r)
 
     //loop over the movies range
    for index, item := range movies{
        if item.ID == params["ID"]{
    //delete the movie with the i/d that you've sent        
            movies = append(movies[:index], movies[index+1:]... )
            var movie Movie
    //add a new movie - the movie that we send in the body of postman        
            _ = json.NewDecoder(r.Body).Decode(&movie)
            movie.ID = params["id"]
            movies = append(movies, movie)
            json.NewEncoder(w).Encode(movie)
            return
        }
    }
}

func main(){
    // CRUD ROUTE 
    r := mux.NewRouter()

    // Main feature of program 
    movies = append(movies, Movie{ID: "1", Isbn: "438227", Title:"Movie One", Director: &Director{Firstname: "john", Lastname: "Doe"}})
    movies = append(movies, Movie{ID:"2", Isbn: "438231", Title: "Movie two", Director: &Director{Firstname: "steve", Lastname: "Smith"}})
    
    // CRUD ROUTE 
    r .HandleFunc("/movies", getMovies).Methods("GET")
    r.HandleFunc("/movies/(id)",getMovie).Methods("GET")
    r.HandleFunc("/movies",createMovie).Methods("POST")
    r.HandleFunc("/movies/(id)", updateMovie).Methods("PUT")
    r.HandleFunc("/movies/(id)",deleteMovie).Methods("DELETE")

    fmt.Printf("Starting server at port 8000\n")
    //To start server 
    log.Fatal(http.ListenAndServe(":8000", r))
}

