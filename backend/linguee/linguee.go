package linguee

import (
	"fmt"
	"vocabulary-helper/utils"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

const (
	LINGUEE_URL = "https://www.linguee.es/espanol-portugues/search?query="
)

type Example struct {
	Source string `json:"source,omitempty"`
	Target string `json:"target,omitempty"`
}

type LingueeSearch struct {
	Found       bool      `json:"found"`
	SearchWord  string    `json:"search_word"`
	Source      string    `json:"source_url,omitempty"`
	Translation string    `json:"translation,omitempty"`
	Examples    []Example `json:"examples,omitempty"`
}

func FindLingueeSearch(word string) LingueeSearch {
	return searchForLingueeInfo(word, fmt.Sprint(LINGUEE_URL, word))
}

func searchForLingueeInfo(word, url string) LingueeSearch {
	c := utils.CreateCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Searching for linguee info ...", r.URL.String())
	})

	lingueeSearch := LingueeSearch{
		Found:      false,
		SearchWord: word,
	}

	c.OnHTML("html", func(e *colly.HTMLElement) {
		info := e.DOM.Find("#dictionary .exact")

		if info.Length() > 0 {
			lingueeSearch.Found = true
			lingueeSearch.Source = url

			// Get translation
			lingueeSearch.Translation = info.Find("div.translation h3.translation_desc a.featured").First().Text()

			// Get examples
			lingueeSearch.Examples = []Example{}
			info.Find("div.translation div.example_lines span.tag_e").Each(func(i int, s *goquery.Selection) {
				lingueeSearch.Examples = append(lingueeSearch.Examples, Example{
					Source: s.Find("span.tag_s").First().Text(),
					Target: s.Find("span.tag_t").First().Text(),
				})
			})
		}
	})

	c.Visit(url)

	return lingueeSearch
}
