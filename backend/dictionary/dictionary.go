package dictionary

import (
	"fmt"
	"vocabulary-helper/utils"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

const (
	DICIO_SEARCH_URL        = "https://www.dicio.com.br/pesquisa.php?q="
	DICIO_DIRECT_URL        = "https://www.dicio.com.br"
	MAX_MEANINGS_TO_EXTRACT = 2
)

type DictionarySearch struct {
	Found      bool
	SearchWord string
	Source     string
	Meanings   []string
	Synonyms   []string
}

func FindDictionaryInfo(word string) DictionarySearch {
	return searchForDictionaryInfo(word, fmt.Sprint(DICIO_SEARCH_URL, word), true)
}

func searchForDictionaryInfo(word, url string, deepSearch bool) DictionarySearch {
	c := utils.CreateCollector()

	dictionaryInfo := DictionarySearch{
		Found:      false,
		SearchWord: word,
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Searching for dictionary info ...", r.URL.String())
	})

	c.OnHTML("html", func(e *colly.HTMLElement) {
		mainContent := e.DOM.Find("#content > div > div.card.card-main").First()
		resultados := e.DOM.Find("#content > div > ul.resultados").First()

		if mainContent.Length() > 0 {
			dictionaryInfo.Found = true
			dictionaryInfo.Source = url

			// Obtain meanings
			dictionaryInfo.Meanings = []string{}
			spans := mainContent.Find("p.significado span:not(.cl)")
			maxMeanings := min(spans.Length(), MAX_MEANINGS_TO_EXTRACT)
			spans.Slice(0, maxMeanings).Each(func(i int, s *goquery.Selection) {
				dictionaryInfo.Meanings = append(dictionaryInfo.Meanings, s.Text())
			})

			// Obtaing sinonimos
			dictionaryInfo.Synonyms = []string{}
			mainContent.Find("div.wrap-section > h2.subtitle-significado ~ p.adicional.sinonimos a").Each(func(i int, s *goquery.Selection) {
				dictionaryInfo.Synonyms = append(dictionaryInfo.Synonyms, s.Text())
			})
		} else if resultados.Length() > 0 && deepSearch {
			match := resultados.Find("li a").FilterFunction(func(i int, s *goquery.Selection) bool {
				return s.Find(".list-link").First().Text() == word
			})

			if match.Length() == 0 {
				match = resultados.Find("li a").First()
			}

			if href, exists := match.Attr("href"); exists {
				dictionaryInfo = searchForDictionaryInfo(word, fmt.Sprint(DICIO_DIRECT_URL, href), false)
			}
		}
	})

	c.Visit(url)

	return dictionaryInfo
}
