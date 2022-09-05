package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

const (
	PARTS_URL_PATH     = "https://lr-parts.com.ua/parts/LAND%20ROVER/"
	PARTS_JSON_PATH    = "./parts.json"
	DIV_MAIN_CONTAINER = "div.wGoodsGroupInfo"
	PART_NUMBER        = "span.article-number"
	DESCRIPTION        = "h2"
	IMAGE              = "div.article-image"
)

type PartNumbers []struct {
	PartNumber string `json:"part_number"`
}

// -----------------------------------------------------------
func checkErr(err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

// -----------------------------------------------------------
func getPartsLinks(partFilePath string) []string {
	var partNumbers PartNumbers

	file, err := os.ReadFile(partFilePath)
	checkErr(err)

	err = json.Unmarshal(file, &partNumbers)
	checkErr(err)

	var links []string
	for i := range partNumbers {
		links = append(links, PARTS_URL_PATH+partNumbers[i].PartNumber)
	}
	return links
}

// -----------------------------------------------------------
func parseParts(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find(DIV_MAIN_CONTAINER).Each(func(_ int, s *goquery.Selection) {

		partNumber := s.Find(PART_NUMBER).Text()
		description := s.Find(DESCRIPTION).Text()
		img, _ := s.Find(IMAGE).Find("a").Attr("href")

		g.Exports <- map[string]interface{}{
			"description": strings.TrimSpace(description),
			"partnumber":  strings.TrimSpace(partNumber),
			"img":         img,
		}
	})
}

// ------------------------------------------------------------
func main() {
	links := getPartsLinks(PARTS_JSON_PATH)

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: links,
		ParseFunc: parseParts,
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()
}
