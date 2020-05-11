package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getJSON(rawMap map[string]interface{}) string {
	json, err := json.Marshal(rawMap)

	if err != nil {
		return ""
	}

	return string(json)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	x := map[string]interface{}{
		"firstName": "Kees",
		"lastName":  "van Voorthuizen",
		"skills":    []string{"Node.js", "TypeScript", "Vue", "Python"},
		"projects":  []string{"Tribecamp"},
	}
	fmt.Fprintf(w, getJSON(x))
}

func main() {
	// Create mux router
	r := mux.NewRouter()
	r.HandleFunc("/", handleRoot)

	// Listen on port 8080
	go http.ListenAndServe(":8080", r)

	// Print status
	fmt.Println("Server is running at http://localhost:8080!")
	fmt.Scanln()
}
