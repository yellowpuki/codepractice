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
	PARTS_URL_PATH     = "https://gb.ru/lessons/"
	PARTS_JSON_PATH    = "./parts.json"
	DIV_MAIN_CONTAINER = "div.main-content-wrapper"
	PART_NUMBER        = "h2"
	DESCRIPTION        = "h3"
	IMAGE              = "div.article-image"
	VIDEO              = "div.vjs-gbui-player-container"
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
		video, _ := s.Find(VIDEO).Find("video").Attr("src")

		g.Exports <- map[string]interface{}{
			"description": strings.TrimSpace(description),
			"partnumber":  strings.TrimSpace(partNumber),
			"video":       video,
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
