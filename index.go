package main

import(
	"log"
	"net/http"
	"github.com/go-chi/chi"
	"routes"

)

func main(){
	router := setRouter()
	log.Println("http://localhost:8080")
	//Open Server at port 8080
	log.Fatal(http.ListenAndServe(":8080",router))
}

func setRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/",routes.TestMain)
	return router
}
