package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"httpmg/cmp"
	"log"
	"net/http"
)

//UserJSON is a struct that gets created and marchseld into a json
type UserJSON struct {
	time string
	id   string
}

func main() {
	r := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
