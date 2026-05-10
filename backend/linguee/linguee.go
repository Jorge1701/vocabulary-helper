package linguee

import (
	"fmt"
	"vocabulary-helper/model"
	"vocabulary-helper/utils"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

const (
	LINGUEE_URL = "https://www.linguee.es/espanol-portugues/search?query="
)

type LingueeSearch struct {
	Found       bool
	SearchWord  string
	Source      string
	Translation string
	Examples    []model.Example
}

func FindInLinguee(word string) LingueeSearch {
	return fetchAndParseLingueeInfo(word, fmt.Sprint(LINGUEE_URL, word))
}

func fetchAndParseLingueeInfo(word, url string) LingueeSearch {
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
			lingueeSearch.Examples = []model.Example{}
			info.Find("div.translation div.example_lines span.tag_e").Each(func(i int, s *goquery.Selection) {
				lingueeSearch.Examples = append(lingueeSearch.Examples, model.Example{
					Source: s.Find("span.tag_s").First().Text(),
					Target: s.Find("span.tag_t").First().Text(),
				})
			})
		}
	})

	c.Visit(url)

	return lingueeSearch
}
