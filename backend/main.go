package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"vocabulary-helper/conjugacao"
	"vocabulary-helper/database"
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
	err := database.NormalizeDatabase()
	if err != nil {
		fmt.Println("Could not normilize database:", err)
	} else {
		fmt.Println("Database normilized")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/words/{words}", func(w http.ResponseWriter, r *http.Request) {
		wordsParam := strings.TrimSpace(strings.ToLower(r.PathValue("words")))
		if wordsParam == "" {
			http.Error(w, `{"error":"word is required"}`, http.StatusBadRequest)
			return
		}

		results := []model.SearchResult{}

		words := strings.SplitSeq(wordsParam, ",")

		for word := range words {
			searchResult := getSearchResultFor(strings.TrimSpace(word))
			results = append(results, searchResult)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	})

	mux.HandleFunc("/word/{word}", func(w http.ResponseWriter, r *http.Request) {
		word := strings.TrimSpace(strings.ToLower(r.PathValue("word")))
		if word == "" {
			http.Error(w, `{"error":"word is required"}`, http.StatusBadRequest)
			return
		}

		searchResult := getSearchResultFor(word)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(searchResult)
	})

	handler := enableCORS(mux)

	http.ListenAndServe(":8080", handler)
}

func getSearchResultFor(word string) model.SearchResult {
	searchResult := model.SearchResult{
		SearchWord: word,
		Type:       "Unknown",
		Meanings:   []model.Meaning{},
		Sources:    map[string]string{},
	}

	conjugacaoResult := conjugacao.FindInConjugacao(word)
	lingueeResult := linguee.FindInLinguee(word)
	databaseSearch := database.FindInDatabase(word)
	dicioResult := dicio.FindInDicio(word)

	if conjugacaoResult.Found {
		searchResult.Type = "Verb"
		searchResult.VerbInfo = &conjugacaoResult.VerbInfo

		searchResult.Sources["Conjugacao"] = conjugacaoResult.Source
	}

	if dicioResult.FoundWord != word && conjugacaoResult.Found {
		dicioResult = dicio.FindInDicio(conjugacaoResult.VerbInfo.Infinitive)
	}

	if lingueeResult.Found {
		searchResult.FoundWord = lingueeResult.FoundWord
		searchResult.Translation = lingueeResult.Translation
		searchResult.Examples = lingueeResult.Examples

		searchResult.Sources["Linguee"] = lingueeResult.Source
	}

	if dicioResult.Found {
		for _, meaning := range dicioResult.Meanings {
			searchResult.Meanings = append(searchResult.Meanings, model.Meaning{
				Text: meaning,
			})
		}
		searchResult.Synonyms = dicioResult.Synonyms

		searchResult.Sources["Dicio"] = dicioResult.Source
	}

	if databaseSearch.Found {
		searchResult.Meanings = append(searchResult.Meanings, databaseSearch.Meaning...)

		searchResult.Sources["WikDict"] = databaseSearch.Source
	}

	if _, exists := searchResult.Sources["Dicio"]; !exists {
		searchResult.Sources["Dicio"] = fmt.Sprint(dicio.DICIO_SEARCH_URL, word)
	}
	if _, exists := searchResult.Sources["Linguee"]; !exists {
		searchResult.Sources["Linguee"] = fmt.Sprint(linguee.LINGUEE_URL, word)
	}
	if _, exists := searchResult.Sources["Conjugacao"]; !exists {
		searchResult.Sources["Conjugacao"] = fmt.Sprint(conjugacao.CONJUGACAO_SEARCH_URL, word)
	}
	if _, exists := searchResult.Sources["WikDict"]; !exists {
		searchResult.Sources["WikDict"] = fmt.Sprint(database.WIKDICT_URL, word)
	}

	// TODO ADD https://pt.wiktionary.org

	return searchResult
}
