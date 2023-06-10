package util

import (
	"encoding/json"
	"fmt"
	"ocean_backend/models"
	"os"
	"strconv"
	"strings"

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
	c.OnHTML("span.color-fg-default.text-bold.mr-1", func(e *colly.HTMLElement) {
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

func (g *GithubParse) GetDescription(endpoint string) (string, error) {

	c := colly.NewCollector()

	var description string
	c.OnHTML("p.f4.my-3", func(e *colly.HTMLElement) {
		description = e.Text
	})

	uri := fmt.Sprintf("%s%s", g.Url, endpoint)
	fmt.Printf("Sending get request to %s\n", uri)
	err := c.Visit(uri)
	if err != nil {
		fmt.Println("Error occur, ", err)
		return "", err
	}

	return description, nil
}

func (g *GithubParse) GetRepos() ([]string, error) {

	c := colly.NewCollector()

	var repositories []string
	c.OnHTML("div.d-inline-block.mb-1", func(e *colly.HTMLElement) {
		repoName := strings.TrimSpace(e.ChildText("a[itemprop='name codeRepository']"))
		forkedFrom := strings.TrimSpace(e.ChildText("span.f6.color-fg-muted.mb-1"))
		if strings.HasPrefix(forkedFrom, "Forked from") {
			// forkedFrom = strings.TrimPrefix(forkedFrom, "Forked from ")
			// fmt.Println("Forked From:", forkedFrom)
			fmt.Println("Skipping this repo..")
		} else {
			fmt.Println("Repository:", repoName)
			repositories = append(repositories, repoName)
		}
	})

	uri := fmt.Sprintf("%s%s", g.Url, "?tab=repositories")
	fmt.Printf("Sending get request to %s\n", uri)
	err := c.Visit(uri)
	if err != nil {
		fmt.Println("Error occur, ", err)
		return nil, err
	}

	return repositories, nil
}

func ScanGithub() {
	maorGithub := GithubParse{
		Url: "https://github.com/maorsabag",
	}

	repos, err := maorGithub.GetRepos()
	if err != nil {
		panic(err)
	}

	for _, repoName := range repos {
		var repo models.Repository
		repo.Name = repoName

		repoEndpoint := fmt.Sprintf("/%s", repoName)

		languages, err := maorGithub.GetLanguages(repoEndpoint)
		if err != nil {
			panic(err)
		}
		allLanguages := strings.Trim(strings.Join(languages, ", "), ", ")
		// fmt.Println(allLanguages)
		repo.Languange = allLanguages

		repoDescription, err := maorGithub.GetDescription(repoEndpoint)
		if err != nil {
			panic(err)
		}
		repoDescription = strings.TrimSpace(repoDescription)
		// fmt.Println(repoDescription)
		repo.Description = repoDescription

		stars, err := maorGithub.GetStars(repoEndpoint)
		if err != nil {
			panic(err)
		}
		// fmt.Println(stars)
		repo.Stars, _ = strconv.Atoi(stars)

		repo.Link = fmt.Sprintf("%s%s", maorGithub.Url, repoEndpoint)

		maorGithub.Repositories = append(maorGithub.Repositories, repo)
	}
	fileContent, _ := json.Marshal(maorGithub.Repositories)

	databasePath := GetPath() + "/models/database.json"
	os.WriteFile(databasePath, fileContent, 0644)
}
