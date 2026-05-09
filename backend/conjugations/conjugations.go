package conjugations

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

const (
	COLLY_CACHE_DIR       = "./colly_cache"
	CONJUGACAO_SEARCH_URL = "https://www.conjugacao.com.br/busca.php?q="
	CONJUGACAO_DIRECT_URL = "https://www.conjugacao.com.br/verbo-"
)

type Conjugations struct {
	FirstPersonSingular  string `json:"first_person_singular,omitempty"`
	SecondPersonSingular string `json:"second_person_singular,omitempty"`
	ThirdPersonSingular  string `json:"third_person_singular,omitempty"`
	FirstPersonPlural    string `json:"first_person_plural,omitempty"`
	SecondPersonPlural   string `json:"second_person_plural,omitempty"`
	ThirdPersonPlural    string `json:"third_person_plural,omitempty"`
}

type VerbInfo struct {
	Found                    bool          `json:"found"`
	Source                   string        `json:"source_url,omitempty"`
	Infinitivo               string        `json:"infinitivo,omitempty"`
	TipoDeVerbo              string        `json:"tipo_de_verbo,omitempty"`
	Gerundio                 string        `json:"gerundio,omitempty"`
	ParticipioPassado        string        `json:"participio_passado,omitempty"`
	Presente                 *Conjugations `json:"presente,omitempty"`
	PreteritoImperfeito      *Conjugations `json:"preterito_imperfeito,omitempty"`
	PreteritoPerfeito        *Conjugations `json:"preterito_perfeito,omitempty"`
	PreteritoMaisQuePerfeito *Conjugations `json:"preterito_mais_que_perfeito,omitempty"`
	FuturoDoPresente         *Conjugations `json:"futuro_do_presente,omitempty"`
	FuturoDoPreterito        *Conjugations `json:"futuro_do_preterito,omitempty"`
}

func parseColon(text string) string {
	if strings.Contains(text, ":") {
		return strings.TrimSpace(strings.Split(text, ":")[1])
	} else {
		return strings.TrimSpace(text)
	}
}

func searchForVerbInfo(url string, deepSearch bool) VerbInfo {
	c := createCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Searching for verb info ...", r.URL.String())
	})

	verbInfo := VerbInfo{
		Found: false,
	}

	c.OnHTML("html", func(e *colly.HTMLElement) {
		info := e.DOM.Find("h1.page-title ~ div.verb-info").First()
		conjugacao := e.DOM.Find("h1.page-title ~ div#conjugacao").First()

		if info.Length() > 0 && conjugacao.Length() > 0 {
			verbInfo.Found = true
			verbInfo.Source = url

			// Obtain info
			verbInfo.Gerundio = parseColon(info.Find("p.verb-info--main > span:nth-child(1) > span > span").First().Text())
			verbInfo.ParticipioPassado = parseColon(info.Find("p.verb-info--main > span:nth-child(2) > span > span.f").First().Text())
			verbInfo.Infinitivo = parseColon(info.Find("p.verb-info--main > span:nth-child(3)").First().Text())
			verbInfo.TipoDeVerbo = parseColon(info.Find("p.verb-info--sec > span:nth-child(1)").First().Text())

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
				verbInfo = searchForVerbInfo(fmt.Sprint(CONJUGACAO_DIRECT_URL, linkToVerb.Text()), false)
			}
		}
	})

	c.Visit(url)

	return verbInfo
}

func FindVerbInfo(word string) VerbInfo {
	return searchForVerbInfo(fmt.Sprint(CONJUGACAO_SEARCH_URL, word), true)
}

func do() {
	words := []string{
		"átimo",
		"pulou",
		"retardatário",
		"sumiu",
		"comprido",
		"atirava",
		"macaquices",
		"tocou",
		"fêmea",
		"ajoelhou",
		"voou",
		"despencou",
		"demorada",
		"estatelar",
		"estendida",
	}

	for _, word := range words {
		verbInfo := FindVerbInfo(word)
		PrintVerbInfo(word, verbInfo)
	}
}

func PrintVerbInfo(word string, verbInfo VerbInfo) {
	if verbInfo.Found {
		fmt.Printf("\n\n---- | '%s'\n\n", word)
		fmt.Printf("Gerúndio: '%s'\n", verbInfo.Gerundio)
		fmt.Printf("Particípio passado: '%s'\n", verbInfo.ParticipioPassado)
		fmt.Printf("Infinitivo: '%s'\n", verbInfo.Infinitivo)
		fmt.Printf("Tipo de verbo: '%s'\n", verbInfo.TipoDeVerbo)

		fmt.Printf("\n--| ")
		fmt.Printf("Presente:\n")
		printConjugations(verbInfo.Presente)
		fmt.Printf("--| ")
		fmt.Printf("Pretérito Imperfeito:\n")
		printConjugations(verbInfo.PreteritoImperfeito)
		fmt.Printf("--| ")
		fmt.Printf("Pretérito Perfeito:\n")
		printConjugations(verbInfo.PreteritoPerfeito)
		fmt.Printf("--| ")
		fmt.Printf("Pretérito Máis-que-perfeito:\n")
		printConjugations(verbInfo.PreteritoMaisQuePerfeito)
		fmt.Printf("--| ")
		fmt.Printf("Futuro do Presente:\n")
		printConjugations(verbInfo.FuturoDoPresente)
		fmt.Printf("--| ")
		fmt.Printf("Futuro do Pretérifo:\n")
		printConjugations(verbInfo.FuturoDoPreterito)
	} else {
		fmt.Printf("\n\n---- | '%s'\n\n", word)
		fmt.Printf("\nNOT FOUND\n\n")
	}

	fmt.Printf("--------------------------------------------------------------------\n")
}

func printConjugations(c *Conjugations) {
	fmt.Printf("Eu %s\n", c.FirstPersonSingular)
	fmt.Printf("Tu %s\n", c.SecondPersonSingular)
	fmt.Printf("Você/ele/ela %s\n", c.ThirdPersonSingular)
	fmt.Printf("Nós %s\n", c.FirstPersonPlural)
	fmt.Printf("Vós %s\n", c.SecondPersonPlural)
	fmt.Printf("Vocês/eles/elas %s\n", c.ThirdPersonPlural)
}

func createCollector() *colly.Collector {
	return colly.NewCollector(
		colly.CacheDir(COLLY_CACHE_DIR),
		colly.CacheExpiration(24*time.Hour),
	)
}
