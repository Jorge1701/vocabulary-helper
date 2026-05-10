package conjugations

import (
	"fmt"
	"strings"
	"vocabulary-helper/utils"

	"github.com/gocolly/colly/v2"
)

const (
	COLLY_CACHE_DIR       = "./colly_cache"
	CONJUGACAO_SEARCH_URL = "https://www.conjugacao.com.br/busca.php?q="
	CONJUGACAO_DIRECT_URL = "https://www.conjugacao.com.br/verbo-"
)

type Conjugations struct {
	FirstPersonSingular  string `json:"first_per_sin,omitempty"`
	SecondPersonSingular string `json:"second_per_sin,omitempty"`
	ThirdPersonSingular  string `json:"third_per_sin,omitempty"`
	FirstPersonPlural    string `json:"first_per_plu,omitempty"`
	SecondPersonPlural   string `json:"second_per_plu,omitempty"`
	ThirdPersonPlural    string `json:"third_per_plu,omitempty"`
}

type VerbInfo struct {
	Infinitivo        string `json:"infinitivo,omitempty"`
	TipoDeVerbo       string `json:"tipo_de_verbo,omitempty"`
	Gerundio          string `json:"gerundio,omitempty"`
	ParticipioPassado string `json:"participio_passado,omitempty"`
}

type ConjugationSearch struct {
	Found                    bool          `json:"found"`
	Source                   string        `json:"source_url,omitempty"`
	SearchWord               string        `json:"search_word"`
	VerbInfo                 *VerbInfo     `json:"verb_info,omitempty"`
	Presente                 *Conjugations `json:"presente,omitempty"`
	PreteritoImperfeito      *Conjugations `json:"preterito_imperfeito,omitempty"`
	PreteritoPerfeito        *Conjugations `json:"preterito_perfeito,omitempty"`
	PreteritoMaisQuePerfeito *Conjugations `json:"preterito_mais_que_perfeito,omitempty"`
	FuturoDoPresente         *Conjugations `json:"futuro_do_presente,omitempty"`
	FuturoDoPreterito        *Conjugations `json:"futuro_do_preterito,omitempty"`
}

func FindVerbInfo(word string) ConjugationSearch {
	return searchForVerbInfo(word, fmt.Sprint(CONJUGACAO_SEARCH_URL, word), true)
}

func searchForVerbInfo(word, url string, deepSearch bool) ConjugationSearch {
	c := utils.CreateCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Searching for verb info ...", r.URL.String())
	})

	verbInfo := ConjugationSearch{
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
			verbInfo.VerbInfo = &VerbInfo{
				Gerundio:          parseColon(info.Find("p.verb-info--main > span:nth-child(1) > span > span").First().Text()),
				ParticipioPassado: parseColon(info.Find("p.verb-info--main > span:nth-child(2) > span > span.f").First().Text()),
				Infinitivo:        parseColon(info.Find("p.verb-info--main > span:nth-child(3)").First().Text()),
				TipoDeVerbo:       parseColon(info.Find("p.verb-info--sec > span:nth-child(1)").First().Text()),
			}

			// Obtain conjugacao
			verbInfo.Presente = &Conjugations{
				FirstPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(1) > p > span:nth-child(1) > span.f").First().Text(),
				SecondPersonSingular: conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(1) > p > span:nth-child(3) > span.f").First().Text(),
				ThirdPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(1) > p > span:nth-child(5) > span.f").First().Text(),
				FirstPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(1) > p > span:nth-child(7) > span.f").First().Text(),
				SecondPersonPlural:   conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(1) > p > span:nth-child(9) > span.f").First().Text(),
				ThirdPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(1) > p > span:nth-child(11) > span.f").First().Text(),
			}

			verbInfo.PreteritoImperfeito = &Conjugations{
				FirstPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(2) > p > span:nth-child(1) > span.f").First().Text(),
				SecondPersonSingular: conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(2) > p > span:nth-child(3) > span.f").First().Text(),
				ThirdPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(2) > p > span:nth-child(5) > span.f").First().Text(),
				FirstPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(2) > p > span:nth-child(7) > span.f").First().Text(),
				SecondPersonPlural:   conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(2) > p > span:nth-child(9) > span.f").First().Text(),
				ThirdPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(2) > p > span:nth-child(11) > span.f").First().Text(),
			}
			verbInfo.PreteritoPerfeito = &Conjugations{
				FirstPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(3) > p > span:nth-child(1) > span.f").First().Text(),
				SecondPersonSingular: conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(3) > p > span:nth-child(3) > span.f").First().Text(),
				ThirdPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(3) > p > span:nth-child(5) > span.f").First().Text(),
				FirstPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(3) > p > span:nth-child(7) > span.f").First().Text(),
				SecondPersonPlural:   conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(3) > p > span:nth-child(9) > span.f").First().Text(),
				ThirdPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(3) > p > span:nth-child(11) > span.f").First().Text(),
			}
			verbInfo.PreteritoMaisQuePerfeito = &Conjugations{
				FirstPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(4) > p > span:nth-child(1) > span.f").First().Text(),
				SecondPersonSingular: conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(4) > p > span:nth-child(3) > span.f").First().Text(),
				ThirdPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(4) > p > span:nth-child(5) > span.f").First().Text(),
				FirstPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(4) > p > span:nth-child(7) > span.f").First().Text(),
				SecondPersonPlural:   conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(4) > p > span:nth-child(9) > span.f").First().Text(),
				ThirdPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(4) > p > span:nth-child(11) > span.f").First().Text(),
			}
			verbInfo.FuturoDoPresente = &Conjugations{
				FirstPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(5) > p > span:nth-child(1) > span.f").First().Text(),
				SecondPersonSingular: conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(5) > p > span:nth-child(3) > span.f").First().Text(),
				ThirdPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(5) > p > span:nth-child(5) > span.f").First().Text(),
				FirstPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(5) > p > span:nth-child(7) > span.f").First().Text(),
				SecondPersonPlural:   conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(5) > p > span:nth-child(9) > span.f").First().Text(),
				ThirdPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(5) > p > span:nth-child(11) > span.f").First().Text(),
			}
			verbInfo.FuturoDoPreterito = &Conjugations{
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
				verbInfo = searchForVerbInfo(word, fmt.Sprint(CONJUGACAO_DIRECT_URL, linkToVerb.Text()), false)
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
