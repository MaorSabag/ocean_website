package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"ocean_backend/routes"
	"os"
)

func main() {
	fmt.Println("Hello World!")

	log.Println("Started serving backend server at http://localhost:1122")
	router := routes.GetRoutes()
	err := http.ListenAndServe(":1122", router)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Failed to start server at localhost:1122")
		os.Exit(1)
	} else if err != nil {
		log.Printf("Error starting server %v\n", err)
		os.Exit(1)
	}

	// maorGithub := util.GithubParse{
	// 	Url: "https://github.com/maorsabag",
	// }
	// body, err := maorGithub.ParseUrl("/sideLoadingDLL")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(body)

	// hrefs, err := maorGithub.GetHrefs("/sideLoadingDLL")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(strings.Join(hrefs, ", "))
	// stars, err := maorGithub.GetStars("/sideLoadingDLL")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Stars ", stars)

	// languages, err := maorGithub.GetLanguages("/sideLoadingDLL")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(languages)

}
