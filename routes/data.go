package routes

/*
data.go
built by Lim Kyu Min.
=============================================
Data.go represents routers dedicated to manipulating data.
Following endpoints are responsible for...
	- receiving data from frontend
	- send data to frontend
Both endpoint will be receive POST request.
*/

import (
	//"github.com/go-chi/chi"
	"log"
	"net/http"
)

// ReceiveData: It receives answers from frontend. (user answer)
func ReceiveData(w http.ResponseWriter, r *http.Request){
	log.Println("SYSTEM: ReceiveData Called")
	//TODO: receive data from frontend
}

// SendData: It sends mbti results to frontend.
func SendData(w http.ResponseWriter, r *http.Request){
	log.Println("SYSTEM: SendData Called")
	//TODO: send data to frontend.
}


