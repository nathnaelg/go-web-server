package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ALPHACOD3RS/go-web-server/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)


type apiConfig struct{
	DB *database.Queries
}

func main(){


	godotenv.Load()

	portString := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")

	if portString == "" {
		log.Fatal("port is not geting from the env")
	}

	if dbUrl == "" {
		log.Fatal("database url is not geting from the env")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil{
		log.Fatal("There is a connection error: ", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELTET", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/health", handlerReadines)
	v1Router.Get("/error", handlerError)
	v1Router.Post("/user", apiCfg.handlerCreateUser)
	v1Router.Get("/user", apiCfg.handlerGetUser)

	router.Mount("/v1", v1Router)




	srver := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}


	log.Printf("the server is runing in port %v", portString);



	err = srver.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(portString)
}
