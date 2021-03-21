package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
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

	//Data Manipulation: Post Request
	router.Route("/data", func (r chi.Router){
		r.Post("/receive",routes.ReceiveData)
		r.Post("/send",routes.SendData)
	})
	return router
}
