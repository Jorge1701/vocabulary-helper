package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"vocabulary-helper/conjugations"
	"vocabulary-helper/dictionary"
)

type WordInfo struct {
	DictionarySearch  *dictionary.DictionarySearch    `json:"dictionary_search,omitempty"`
	ConjugationSearch *conjugations.ConjugationSearch `json:"conjugation_search,omitempty"`
}

func main() {
	http.HandleFunc("/word/{word}", func(w http.ResponseWriter, r *http.Request) {
		word := strings.TrimSpace(strings.ToLower(r.PathValue("word")))
		if word == "" {
			http.Error(w, `{"error":"word is required"}`, http.StatusBadRequest)
			return
		}

		verbInfo := conjugations.FindVerbInfo(word)
		dictionaryInfo := dictionary.DictionarySearch{}

		if verbInfo.Found {
			dictionaryInfo = dictionary.FindDictionaryInfo(verbInfo.VerbInfo.Infinitivo)
		} else {
			dictionaryInfo = dictionary.FindDictionaryInfo(word)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(WordInfo{
			DictionarySearch:  &dictionaryInfo,
			ConjugationSearch: &verbInfo,
		})
	})

	http.ListenAndServe(":8080", nil)
}
