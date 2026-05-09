package main

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

type WordInfo struct {
	meaning string
}

type Conjugations struct {
	firstPersonSingular  string
	secondPersonSingular string
	thirdPersonSingular  string
	firstPersonPlural    string
	secondPersonPlural   string
	thirdPersonPlugar    string
}

type VerbInfo struct {
	found                    bool
	infinitivo               string
	tipoDeVerbo              string
	gerundio                 string
	participioPassado        string
	presente                 Conjugations
	preteritoImperfeito      Conjugations
	preteritoPerfeito        Conjugations
	preteritoMaisQuePerfeito Conjugations
	futuroDoPresente         Conjugations
	futuroDoPreterito        Conjugations
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
		found: false,
	}

	c.OnHTML("html", func(e *colly.HTMLElement) {
		info := e.DOM.Find("h1.page-title ~ div.verb-info").First()
		conjugacao := e.DOM.Find("h1.page-title ~ div#conjugacao").First()

		if info.Length() > 0 && conjugacao.Length() > 0 {
			verbInfo.found = true

			// Obtain info
			verbInfo.gerundio = parseColon(info.Find("p.verb-info--main > span:nth-child(1) > span > span").First().Text())
			verbInfo.participioPassado = parseColon(info.Find("p.verb-info--main > span:nth-child(2) > span > span.f").First().Text())
			verbInfo.infinitivo = parseColon(info.Find("p.verb-info--main > span:nth-child(3)").First().Text())
			verbInfo.tipoDeVerbo = parseColon(info.Find("p.verb-info--sec > span:nth-child(1)").First().Text())

			// Obtain conjugacao
			verbInfo.presente = Conjugations{
				firstPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(1) > p > span:nth-child(1) > span.f").First().Text(),
				secondPersonSingular: conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(1) > p > span:nth-child(3) > span.f").First().Text(),
				thirdPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(1) > p > span:nth-child(5) > span.f").First().Text(),
				firstPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(1) > p > span:nth-child(7) > span.f").First().Text(),
				secondPersonPlural:   conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(1) > p > span:nth-child(9) > span.f").First().Text(),
				thirdPersonPlugar:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(1) > p > span:nth-child(11) > span.f").First().Text(),
			}

			verbInfo.preteritoImperfeito = Conjugations{
				firstPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(2) > p > span:nth-child(1) > span.f").First().Text(),
				secondPersonSingular: conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(2) > p > span:nth-child(3) > span.f").First().Text(),
				thirdPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(2) > p > span:nth-child(5) > span.f").First().Text(),
				firstPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(2) > p > span:nth-child(7) > span.f").First().Text(),
				secondPersonPlural:   conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(2) > p > span:nth-child(9) > span.f").First().Text(),
				thirdPersonPlugar:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(2) > p > span:nth-child(11) > span.f").First().Text(),
			}
			verbInfo.preteritoPerfeito = Conjugations{
				firstPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(3) > p > span:nth-child(1) > span.f").First().Text(),
				secondPersonSingular: conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(3) > p > span:nth-child(3) > span.f").First().Text(),
				thirdPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(3) > p > span:nth-child(5) > span.f").First().Text(),
				firstPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(3) > p > span:nth-child(7) > span.f").First().Text(),
				secondPersonPlural:   conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(3) > p > span:nth-child(9) > span.f").First().Text(),
				thirdPersonPlugar:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(3) > p > span:nth-child(11) > span.f").First().Text(),
			}
			verbInfo.preteritoMaisQuePerfeito = Conjugations{
				firstPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(4) > p > span:nth-child(1) > span.f").First().Text(),
				secondPersonSingular: conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(4) > p > span:nth-child(3) > span.f").First().Text(),
				thirdPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(4) > p > span:nth-child(5) > span.f").First().Text(),
				firstPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(4) > p > span:nth-child(7) > span.f").First().Text(),
				secondPersonPlural:   conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(4) > p > span:nth-child(9) > span.f").First().Text(),
				thirdPersonPlugar:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(4) > p > span:nth-child(11) > span.f").First().Text(),
			}
			verbInfo.futuroDoPresente = Conjugations{
				firstPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(5) > p > span:nth-child(1) > span.f").First().Text(),
				secondPersonSingular: conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(5) > p > span:nth-child(3) > span.f").First().Text(),
				thirdPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(5) > p > span:nth-child(5) > span.f").First().Text(),
				firstPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(5) > p > span:nth-child(7) > span.f").First().Text(),
				secondPersonPlural:   conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(5) > p > span:nth-child(9) > span.f").First().Text(),
				thirdPersonPlugar:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(5) > p > span:nth-child(11) > span.f").First().Text(),
			}
			verbInfo.futuroDoPreterito = Conjugations{
				firstPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(6) > p > span:nth-child(1) > span.f").First().Text(),
				secondPersonSingular: conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(6) > p > span:nth-child(3) > span.f").First().Text(),
				thirdPersonSingular:  conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(6) > p > span:nth-child(5) > span.f").First().Text(),
				firstPersonPlural:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(6) > p > span:nth-child(7) > span.f").First().Text(),
				secondPersonPlural:   conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(6) > p > span:nth-child(9) > span.f").First().Text(),
				thirdPersonPlugar:    conjugacao.Find("div:nth-child(1) > div > div > div:nth-child(6) > p > span:nth-child(11) > span.f").First().Text(),
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

func findVerbInfo(word string) VerbInfo {
	return searchForVerbInfo(fmt.Sprint(CONJUGACAO_SEARCH_URL, word), true)
}

func main() {
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
		fmt.Printf("--------------------------------------------------------------------\n")

		verbInfo := findVerbInfo(word)

		if verbInfo.found {
			fmt.Printf("\n\n---- | '%s'\n\n", word)
			fmt.Printf("Gerúndio: '%s'\n", verbInfo.gerundio)
			fmt.Printf("Particípio passado: '%s'\n", verbInfo.participioPassado)
			fmt.Printf("Infinitivo: '%s'\n", verbInfo.infinitivo)
			fmt.Printf("Tipo de verbo: '%s'\n", verbInfo.tipoDeVerbo)

			fmt.Printf("\n--| ")
			fmt.Printf("Presente:\n")
			printConjugations(verbInfo.presente)
			fmt.Printf("--| ")
			fmt.Printf("Pretérito Imperfeito:\n")
			printConjugations(verbInfo.preteritoImperfeito)
			fmt.Printf("--| ")
			fmt.Printf("Pretérito Perfeito:\n")
			printConjugations(verbInfo.preteritoPerfeito)
			fmt.Printf("--| ")
			fmt.Printf("Pretérito Máis-que-perfeito:\n")
			printConjugations(verbInfo.preteritoMaisQuePerfeito)
			fmt.Printf("--| ")
			fmt.Printf("Futuro do Presente:\n")
			printConjugations(verbInfo.futuroDoPresente)
			fmt.Printf("--| ")
			fmt.Printf("Futuro do Pretérifo:\n")
			printConjugations(verbInfo.futuroDoPreterito)
		} else {
			fmt.Printf("\n\n---- | '%s'\n\n", word)
			fmt.Printf("\nNOT FOUND\n\n")
		}

		fmt.Printf("--------------------------------------------------------------------\n")
	}
}

func printConjugations(c Conjugations) {
	fmt.Printf("Eu %s\n", c.firstPersonSingular)
	fmt.Printf("Tu %s\n", c.secondPersonSingular)
	fmt.Printf("Você/ele/ela %s\n", c.thirdPersonSingular)
	fmt.Printf("Nós %s\n", c.firstPersonPlural)
	fmt.Printf("Vós %s\n", c.secondPersonPlural)
	fmt.Printf("Vocês/eles/elas %s\n", c.thirdPersonPlugar)
}

func createCollector() *colly.Collector {
	return colly.NewCollector(
		colly.CacheDir(COLLY_CACHE_DIR),
		colly.CacheExpiration(24*time.Hour),
	)
}
