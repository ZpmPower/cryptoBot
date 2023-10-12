package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type NewInfo struct {
	title   string
	link    string
	source  string
	timeAgo string
	content string
}

func GetNews() ([]NewInfo, error) {

	url := "https://ru.investing.com/news/cryptocurrency-news"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status code: %d", resp.StatusCode)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var news []NewInfo

	doc.Find("article").Each(func(index int, article *goquery.Selection) {
		aElement := article.Find("a.title")
		href, _ := aElement.Attr("href")
		title := aElement.Text()
		articleDetails := article.Find("span.articleDetails")
		source := strings.TrimSpace(articleDetails.Find("span").First().Text())
		timeAgo := strings.TrimSpace(articleDetails.Find("span.date").Text())
		content := strings.TrimSpace(article.Find("p").Text())

		if title != "" && href != "" && source != "" && timeAgo != "" && content != "" {

			news = append(news, NewInfo{title: title, link: href, source: source, timeAgo: timeAgo, content: content})
			fmt.Printf("Title: %s\n", title)
			fmt.Printf("Href: %s\n", href)
			fmt.Printf("Source: %s\n", source)
			fmt.Printf("TimeAgo: %s\n", timeAgo)
			fmt.Printf("Content: %s\n", content)
			fmt.Println("----------")
		}

	})

	return news, nil
}
