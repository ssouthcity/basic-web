package main

import (
	"os"
	"fmt"
	"net/http"
	"encoding/json"
)

type Config struct {
	Name string `json:"name"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	http.HandleFunc("/", handle)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	conf := Config{}

	f, err := os.Open("config.json")
	if err != nil {
		fmt.Fprint(w, "unexpected open error")
		return
	}

	if err := json.NewDecoder(f).Decode(&conf); err != nil {
		fmt.Fprint(w, "unexpected parse error")
		return
	}

	fmt.Fprintf(w, "hello %s", conf.Name)
}
