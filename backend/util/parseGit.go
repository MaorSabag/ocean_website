package util

import (
	"fmt"
	"ocean_backend/models"

	"github.com/gocolly/colly"
)

type GithubParse struct {
	Url          string
	Repositories models.Database
}

func (g *GithubParse) ParseUrl(endpoint string) (string, error) {

	c := colly.NewCollector()
	var result string
	uri := fmt.Sprintf("%s%s", g.Url, endpoint)
	fmt.Printf("Sending Get Request to %s\n", uri)

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got results!")
		result = string(r.Body)
	})

	err := c.Visit(uri)
	fmt.Println("Visiting..")
	if err != nil {
		fmt.Println("Error occur, ", err)
		return "", err
	}

	fmt.Println(result)

	return result, nil

}

func (g *GithubParse) GetHrefs(endpoint string) ([]string, error) {

	c := colly.NewCollector()
	var hrefs []string
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		hrefs = append(hrefs, href)
	})

	uri := fmt.Sprintf("%s%s", g.Url, endpoint)
	fmt.Printf("Sending get request to %s\n", uri)

	err := c.Visit(uri)
	if err != nil {
		fmt.Println("Error occur, ", err)
		return nil, err
	}

	return hrefs, nil
}

func (g *GithubParse) GetStars(endpoint string) (string, error) {

	c := colly.NewCollector()

	var stars string
	c.OnHTML("span#repo-stars-counter-star", func(e *colly.HTMLElement) {
		stars = e.Text
	})

	uri := fmt.Sprintf("%s%s", g.Url, endpoint)
	fmt.Printf("Sending get request to %s\n", uri)
	err := c.Visit(uri)
	if err != nil {
		fmt.Println("Error occur, ", err)
		return "", err
	}

	return stars, nil
}

func (g *GithubParse) GetLanguages(endpoint string) ([]string, error) {

	c := colly.NewCollector()

	var languages []string
	c.OnHTML("div.mb-2", func(e *colly.HTMLElement) {
		languages = append(languages, e.Text)
	})

	uri := fmt.Sprintf("%s%s", g.Url, endpoint)
	fmt.Printf("Sending get request to %s\n", uri)
	err := c.Visit(uri)
	if err != nil {
		fmt.Println("Error occur, ", err)
		return nil, err
	}

	return languages, nil
}
