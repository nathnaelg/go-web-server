package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ALPHACOD3RS/go-web-server/internal/auth"
	"github.com/ALPHACOD3RS/go-web-server/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request){

	type params struct{
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	para := params{}

	err := decoder.Decode(&para)
	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Eror json: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.NullUUID{UUID: uuid.New(), Valid: true },
		Name: para.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("couldnt create user error: %v", err))
		// log.Fatalf("couldnt create user error: %v", err)
		return
	}



	respondWithJson(w, 201 , dbUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request){

	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldnt get apikey: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldnt get user: %v", err))
		return
	}

	respondWithJson(w, 200 , dbUserToUser(user))

}