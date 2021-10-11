package handlers

import (
	"log"
	"net/http"

	"github.com/damuel90/twitgo/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func RunHandlers() {
	PORT := utils.Getenv("PORT", ":8080")
	router := mux.NewRouter()

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(PORT, handler))
}
