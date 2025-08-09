package scrapper

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ScrapDataWithAttribute(url, tagFilter, attributeSearched string) []string {
	findValues := make([]string, 0)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Request Error: ", err)
	}
	defer response.Body.Close()

	pageDocument, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("pageDocument Analyze Error: ", err)
	}

	pageDocument.Find(tagFilter).Each(
		func(_ int, item *goquery.Selection) {
			findValues = append(findValues, findTagWithAttribute(item, attributeSearched))
		},
	)

	return findValues
}

func findTagWithAttribute(item *goquery.Selection, attributedSearched string) string {

	attributeTag, ok := item.Attr(attributedSearched)
	if ok {
		return attributeTag
	}

	return ""
}

func findTag(item *goquery.Selection) string {
	tag := item.Text()

	return tag
}
