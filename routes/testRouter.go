package routes

import (
	"net/http"
)

func TestMain(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("TEST MAIN PAGE: IF YOU ARE NOT THE ADMIN BUT STILL YOU CAN SEE THIS MESSAGE, SOMETHING MUST BE WRONG."))
}
