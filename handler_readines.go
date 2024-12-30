package main

import "net/http"

func handlerReadines(w http.ResponseWriter, r *http.Request){
	respondWithJson(w, 200 , struct{}{})
}