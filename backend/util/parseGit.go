package util

import (
	"encoding/json"
	"fmt"
	"ocean_backend/models"
	"os"
	"strconv"
	"strings"
	"sync"

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

func (g *GithubParse) GetReleaseDate(endpoint string) (string, error) {

	c := colly.NewCollector()

	var releaseDates []string
	c.OnHTML("h2.f5.text-normal", func(e *colly.HTMLElement) {
		releaseDates = append(releaseDates, e.Text)
	})

	uri := fmt.Sprintf("%s%s/commits/main", g.Url, endpoint)
	fmt.Printf("Sending get request to %s\n", uri)
	err := c.Visit(uri)
	if err != nil {
		fmt.Println("Error occur, ", err)
		return "", err
	}
	var lastReleaseDate string
	if len(releaseDates) > 0 {
		lastReleaseDate = strings.Trim(releaseDates[len(releaseDates)-1], "Commits on")
	} else {
		lastReleaseDate = ""
	}

	return lastReleaseDate, nil
}

func ScanGithub(username string) (bool, error) {
	if username == "" {
		username = "maorsabag"
	}
	maorGithub := GithubParse{
		Url: fmt.Sprintf("https://github.com/%s", username),
	}

	repos, err := maorGithub.GetRepos()
	if err != nil {
		return false, err
	}

	var wg sync.WaitGroup

	wg.Add(len(repos))

	for i := 0; i < len(repos); i++ {
		go func(i int) {
			defer wg.Done()
			var repo models.Repository
			repo.Name = repos[i]

			repoEndpoint := fmt.Sprintf("/%s", repo.Name)

			languages, err := maorGithub.GetLanguages(repoEndpoint)
			if err != nil {
				repo.Languange = "Could not find languages"
			} else {
				allLanguages := strings.Trim(strings.Join(languages, ", "), ", ")
				repo.Languange = allLanguages
			}

			repoDescription, err := maorGithub.GetDescription(repoEndpoint)
			if err != nil {
				repo.Description = "Could not get description"
			} else {
				repoDescription = strings.TrimSpace(repoDescription)
				repo.Description = repoDescription

			}

			stars, err := maorGithub.GetStars(repoEndpoint)
			if err != nil {
				repo.Stars = 0
			} else {
				repo.Stars, _ = strconv.Atoi(stars)

			}

			releaseDate, err := maorGithub.GetReleaseDate(repoEndpoint)
			if err != nil {
				repo.ReleaseDate = "Could not get release date"
			} else {
				repo.ReleaseDate = releaseDate
				repo.Link = fmt.Sprintf("%s%s", maorGithub.Url, repoEndpoint)

			}

			maorGithub.Repositories = append(maorGithub.Repositories, repo)
		}(i)
	}
	wg.Wait()
	fileContent, _ := json.Marshal(maorGithub.Repositories)

	databasePath := GetPath() + fmt.Sprintf("/models/%s.json", username)
	err = os.WriteFile(databasePath, fileContent, 0644)
	if err != nil {
		fmt.Printf("Error writing to %s.json!\n", username)
		return false, err
	}
	fmt.Printf("Wrote database.json\n%d repositories were written\n", len(maorGithub.Repositories))

	return true, nil

}
