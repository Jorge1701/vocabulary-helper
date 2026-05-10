package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"vocabulary-helper/conjugations"
	"vocabulary-helper/dictionary"
	"vocabulary-helper/linguee"
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

		verbInfo := conjugations.FindVerbInfo(word)
		dictionaryInfo := dictionary.DictionarySearch{}
		lingueeSearch := linguee.LingueeSearch{}

		if verbInfo.Found {
			dictionaryInfo = dictionary.FindDictionaryInfo(verbInfo.VerbInfo.Infinitivo)
			lingueeSearch = linguee.FindLingueeSearch(verbInfo.VerbInfo.Infinitivo)
		} else {
			dictionaryInfo = dictionary.FindDictionaryInfo(word)
			lingueeSearch = linguee.FindLingueeSearch(word)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(WordInfo{
			DictionarySearch:  &dictionaryInfo,
			ConjugationSearch: &verbInfo,
			LingueeSearch:     &lingueeSearch,
		})
	})

	handler := enableCORS(mux)

	http.ListenAndServe(":8080", handler)
}
