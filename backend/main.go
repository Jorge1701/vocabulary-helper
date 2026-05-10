package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"vocabulary-helper/conjugations"
	"vocabulary-helper/dictionary"
	"vocabulary-helper/linguee"
	"vocabulary-helper/model"
)

type WordInfo struct {
	DictionarySearch  *dictionary.DictionarySearch    `json:"dictionary_search,omitempty"`
	ConjugationSearch *conjugations.ConjugationSearch `json:"conjugation_search,omitempty"`
	LingueeSearch     *linguee.LingueeSearch          `json:"linguee_search,omitempty"`
}

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

		lingueeSearch := linguee.FindLingueeSearch(word)
		dictionaryInfo := dictionary.FindDictionaryInfo(word)
		verbInfo := conjugations.FindVerbInfo(word)

		if !dictionaryInfo.Found && !verbInfo.Found && !lingueeSearch.Found {
			http.Error(w, `{"error":"word not found"}`, http.StatusNotFound)
			return
		}

		searchResult := model.SearchResult{
			SearchWord: word,
			Examples:   []model.Example{},
			Sources:    map[string]string{},
		}

		if lingueeSearch.Found {
			searchResult.FoundWord = lingueeSearch.SearchWord
			searchResult.Translation = lingueeSearch.Translation
			searchResult.Examples = lingueeSearch.Examples

			searchResult.Sources["Linguee"] = lingueeSearch.Source
		}

		if dictionaryInfo.Found {
			searchResult.Meanings = dictionaryInfo.Meanings
			searchResult.Synonyms = dictionaryInfo.Synonyms

			searchResult.Sources["Dicio"] = dictionaryInfo.Source
		}

		if verbInfo.Found {
			searchResult.VerbInfo = &verbInfo.VerbInfo

			searchResult.Sources["Conjugacao"] = verbInfo.Source
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(searchResult)
	})

	handler := enableCORS(mux)

	http.ListenAndServe(":8080", handler)
}
