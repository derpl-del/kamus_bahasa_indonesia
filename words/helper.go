package words

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetWordKbbi(word string) Words {
	var row []string
	var resp []string
	url := "https://kbbi.kemdikbud.go.id/entri/" + word
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".body-content").Find("li").Each(func(i int, sel *goquery.Selection) {
		row = append(row, sel.Text())
	})
	for x := 0; x <= len(row)-4; x++ {
		resp = append(resp, strings.TrimSpace(strings.Split(row[x], "    ")[len(strings.Split(row[x], "    "))-1]))
	}
	return Words{Kata: word, Arti: resp}
}
