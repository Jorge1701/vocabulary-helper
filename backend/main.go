package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"vocabulary-helper/conjugacao"
	"vocabulary-helper/dicio"
	"vocabulary-helper/linguee"
	"vocabulary-helper/model"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // TODO fix domain
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/word/{word}", func(w http.ResponseWriter, r *http.Request) {
		word := strings.TrimSpace(strings.ToLower(r.PathValue("word")))
		if word == "" {
			http.Error(w, `{"error":"word is required"}`, http.StatusBadRequest)
			return
		}

		lingueeResult := linguee.FindInLinguee(word)
		dicioResult := dicio.FindInDicio(word)
		conjugacaoResult := conjugacao.FindInConjugacao(word)

		if !dicioResult.Found && !conjugacaoResult.Found && !lingueeResult.Found {
			http.Error(w, `{"error":"word not found"}`, http.StatusNotFound)
			return
		}

		searchResult := model.SearchResult{
			SearchWord: word,
			Sources:    map[string]string{},
		}

		if lingueeResult.Found {
			searchResult.FoundWord = lingueeResult.SearchWord
			searchResult.Translation = lingueeResult.Translation
			searchResult.Examples = lingueeResult.Examples

			searchResult.Sources["Linguee"] = lingueeResult.Source
		}

		if dicioResult.Found {
			searchResult.Meanings = dicioResult.Meanings
			searchResult.Synonyms = dicioResult.Synonyms

			searchResult.Sources["Dicio"] = dicioResult.Source
		}

		if conjugacaoResult.Found {
			searchResult.VerbInfo = &conjugacaoResult.VerbInfo

			searchResult.Sources["Conjugacao"] = conjugacaoResult.Source
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(searchResult)
	})

	handler := enableCORS(mux)

	http.ListenAndServe(":8080", handler)
}
