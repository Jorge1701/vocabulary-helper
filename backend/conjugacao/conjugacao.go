package conjugacao

import (
	"fmt"
	"strings"
	"vocabulary-helper/model"
	"vocabulary-helper/utils"

	"github.com/gocolly/colly/v2"
)

const (
	CONJUGACAO_SEARCH_URL = "https://www.conjugacao.com.br/busca.php?q="
	CONJUGACAO_DIRECT_URL = "https://www.conjugacao.com.br/verbo-"
)

type ConjugacaoSearch struct {
	Found      bool
	SearchWord string
	Source     string
	VerbInfo   model.VerbInfo
}

func FindConjugacaoInfo(word string) ConjugacaoSearch {
	return searhForConjugacaoInfo(word, fmt.Sprint(CONJUGACAO_SEARCH_URL, word), true)
}

func searhForConjugacaoInfo(word, url string, deepSearch bool) ConjugacaoSearch {
	c := utils.CreateCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Searching for verb info ...", r.URL.String())
	})

	verbInfo := ConjugacaoSearch{
		Found:      false,
		SearchWord: word,
	}

	c.OnHTML("html", func(e *colly.HTMLElement) {
		info := e.DOM.Find("h1.page-title ~ div.verb-info").First()
		conjugacao := e.DOM.Find("h1.page-title ~ div#conjugacao").First()

		if info.Length() > 0 && conjugacao.Length() > 0 {
			verbInfo.Found = true
			verbInfo.Source = url

			// Obtain info
			verbInfo.VerbInfo = model.VerbInfo{
				Type:             parseColon(info.Find("p.verb-info--sec > span:nth-child(1)").First().Text()),
				Infinitive:       parseColon(info.Find("p.verb-info--main > span:nth-child(3)").First().Text()),
				PresentPaticiple: parseColon(info.Find("p.verb-info--main > span:nth-child(1) > span > span").First().Text()),
				PastParticiple:   parseColon(info.Find("p.verb-info--main > span:nth-child(2) > span > span.f").First().Text()),
			}

			// Obtain conjugacao
			verbInfo.VerbInfo.SimplePresent = model.VerbConjugations{
				FirstPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(1) > p > span:nth-child(1) > span.f").First().Text(),
				SecondPersonSingular: conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(1) > p > span:nth-child(3) > span.f").First().Text(),
				ThirdPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(1) > p > span:nth-child(5) > span.f").First().Text(),
				FirstPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(1) > p > span:nth-child(7) > span.f").First().Text(),
				SecondPersonPlural:   conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(1) > p > span:nth-child(9) > span.f").First().Text(),
				ThirdPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(1) > p > span:nth-child(11) > span.f").First().Text(),
			}

			verbInfo.VerbInfo.ImperfectPast = model.VerbConjugations{
				FirstPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(2) > p > span:nth-child(1) > span.f").First().Text(),
				SecondPersonSingular: conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(2) > p > span:nth-child(3) > span.f").First().Text(),
				ThirdPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(2) > p > span:nth-child(5) > span.f").First().Text(),
				FirstPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(2) > p > span:nth-child(7) > span.f").First().Text(),
				SecondPersonPlural:   conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(2) > p > span:nth-child(9) > span.f").First().Text(),
				ThirdPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(2) > p > span:nth-child(11) > span.f").First().Text(),
			}
			verbInfo.VerbInfo.SimplePast = model.VerbConjugations{
				FirstPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(3) > p > span:nth-child(1) > span.f").First().Text(),
				SecondPersonSingular: conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(3) > p > span:nth-child(3) > span.f").First().Text(),
				ThirdPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(3) > p > span:nth-child(5) > span.f").First().Text(),
				FirstPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(3) > p > span:nth-child(7) > span.f").First().Text(),
				SecondPersonPlural:   conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(3) > p > span:nth-child(9) > span.f").First().Text(),
				ThirdPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(3) > p > span:nth-child(11) > span.f").First().Text(),
			}
			verbInfo.VerbInfo.PerfectPast = model.VerbConjugations{
				FirstPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(4) > p > span:nth-child(1) > span.f").First().Text(),
				SecondPersonSingular: conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(4) > p > span:nth-child(3) > span.f").First().Text(),
				ThirdPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(4) > p > span:nth-child(5) > span.f").First().Text(),
				FirstPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(4) > p > span:nth-child(7) > span.f").First().Text(),
				SecondPersonPlural:   conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(4) > p > span:nth-child(9) > span.f").First().Text(),
				ThirdPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(4) > p > span:nth-child(11) > span.f").First().Text(),
			}
			verbInfo.VerbInfo.SimpleFuture = model.VerbConjugations{
				FirstPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(5) > p > span:nth-child(1) > span.f").First().Text(),
				SecondPersonSingular: conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(5) > p > span:nth-child(3) > span.f").First().Text(),
				ThirdPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(5) > p > span:nth-child(5) > span.f").First().Text(),
				FirstPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(5) > p > span:nth-child(7) > span.f").First().Text(),
				SecondPersonPlural:   conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(5) > p > span:nth-child(9) > span.f").First().Text(),
				ThirdPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(5) > p > span:nth-child(11) > span.f").First().Text(),
			}
			verbInfo.VerbInfo.Conditional = model.VerbConjugations{
				FirstPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(6) > p > span:nth-child(1) > span.f").First().Text(),
				SecondPersonSingular: conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(6) > p > span:nth-child(3) > span.f").First().Text(),
				ThirdPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(6) > p > span:nth-child(5) > span.f").First().Text(),
				FirstPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(6) > p > span:nth-child(7) > span.f").First().Text(),
				SecondPersonPlural:   conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(6) > p > span:nth-child(9) > span.f").First().Text(),
				ThirdPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(6) > p > span:nth-child(11) > span.f").First().Text(),
			}
		} else if deepSearch {
			linkToVerb := e.DOM.Find("#content div > h2 > a")

			if linkToVerb.Length() > 0 {
				verbInfo = searhForConjugacaoInfo(word, fmt.Sprint(CONJUGACAO_DIRECT_URL, linkToVerb.Text()), false)
			}
		}
	})

	c.Visit(url)

	return verbInfo
}

func parseColon(text string) string {
	if strings.Contains(text, ":") {
		return strings.TrimSpace(strings.Split(text, ":")[1])
	} else {
		return strings.TrimSpace(text)
	}
}
