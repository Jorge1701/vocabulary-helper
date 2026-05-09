package main

import (
	"encoding/json"
	"net/http"
	"vocabulary-helper/conjugations"
)

func main() {
	http.HandleFunc("/word/{word}", func(w http.ResponseWriter, r *http.Request) {
		word := r.PathValue("word")
		if word == "" {
			http.Error(w, `{"error":"word is required"}`, http.StatusBadRequest)
			return
		}

		verbInfo := conjugations.FindVerbInfo(word)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(verbInfo)
	})

	http.ListenAndServe(":8080", nil)
}
