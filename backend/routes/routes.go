package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"ocean_backend/models"
	"ocean_backend/util"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func GetRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/", testData).Methods("GET")
	router.HandleFunc("/api/repositories", getRepositories).Methods("GET")
	router.HandleFunc("/api/repositories/{username}", getUsernameRepositories).Methods("GET", "POST")
	router.HandleFunc("/api/getRepoGoogle", GetReposGoogle).Methods("GET", "POST")

	controller := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(router)

	return controller
}

func testData(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got /api request from %v\n", r.RemoteAddr)

	test := &models.Repository{
		Name:        "Maor",
		Languange:   "C",
		Stars:       1,
		Description: "Testing data here",
	}

	json_response, _ := json.Marshal(test)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json_response)

}

func getRepositories(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got /api/repositories request from %v\n", r.RemoteAddr)
	json_database, err := util.GetDatabaseFile("")

	if err != nil {
		fmt.Println(err)
		error_resopsne, _ := json.Marshal(map[string]string{
			"Status": "Not OK",
			"Error":  fmt.Sprintf("Could not found Database: %v", err),
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(error_resopsne)
		return

	}
	response, _ := json.Marshal(json_database)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func getUsernameRepositories(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	log.Printf("Got /api/repositories/%s request from %v", username, r.RemoteAddr)

	reqBody, _ := ioutil.ReadAll(r.Body)

	var jsonBody map[string]interface{}
	json.Unmarshal(reqBody, &jsonBody)
	_, ok := jsonBody["force"]

	checkRepos, err := util.GetDatabaseFile(username)
	if err == nil && !ok {
		jsonResponse, _ := json.Marshal(checkRepos)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
		return
	}

	_, err = util.ScanGithub(username)
	if err != nil {
		response, _ := json.Marshal(map[string]string{
			"Status":  "NOT OK",
			"Message": fmt.Sprintf("Error %s", err),
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return
	}
	response, _ := util.GetDatabaseFile(username)
	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func GetReposGoogle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got /getRepoGoogle request from %v\n", r.RemoteAddr)
	reqBody, _ := ioutil.ReadAll(r.Body)

	var jsonBody map[string]interface{}
	json.Unmarshal(reqBody, &jsonBody)

	text := jsonBody["queryResult"].(map[string]interface{})["parameters"].(map[string]interface{})["repoName"]

	textResponse := setTextResponse(text.(string))

	jsonResponse, _ := json.Marshal(textResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func setTextResponse(text string) models.Response {

	response := models.Response{
		FulfillmentText: "This is a text response",
		FulfillmentMessages: []struct {
			Text struct {
				Text []string `json:"text"`
			} `json:"text"`
		}{
			{
				Text: struct {
					Text []string `json:"text"`
				}{
					Text: []string{text},
				},
			},
		},
		Source: "example.com",
		Payload: struct {
			Google struct {
				ExpectUserResponse bool `json:"expectUserResponse"`
				RichResponse       struct {
					Items []struct {
						SimpleResponse struct {
							TextToSpeech string `json:"textToSpeech"`
						} `json:"simpleResponse"`
					} `json:"items"`
				} `json:"richResponse"`
			} `json:"google"`
			Facebook struct {
				Text string `json:"text"`
			} `json:"facebook"`
			Slack struct {
				Text string `json:"text"`
			} `json:"slack"`
		}{
			Google: struct {
				ExpectUserResponse bool `json:"expectUserResponse"`
				RichResponse       struct {
					Items []struct {
						SimpleResponse struct {
							TextToSpeech string `json:"textToSpeech"`
						} `json:"simpleResponse"`
					} `json:"items"`
				} `json:"richResponse"`
			}{
				ExpectUserResponse: true,
				RichResponse: struct {
					Items []struct {
						SimpleResponse struct {
							TextToSpeech string `json:"textToSpeech"`
						} `json:"simpleResponse"`
					} `json:"items"`
				}{
					Items: []struct {
						SimpleResponse struct {
							TextToSpeech string `json:"textToSpeech"`
						} `json:"simpleResponse"`
					}{
						{
							SimpleResponse: struct {
								TextToSpeech string `json:"textToSpeech"`
							}{
								TextToSpeech: "this is a simple response",
							},
						},
					},
				},
			},
			Facebook: struct {
				Text string `json:"text"`
			}{
				Text: "Hello, Facebook!",
			},
			Slack: struct {
				Text string `json:"text"`
			}{
				Text: "This is a text response for Slack.",
			},
		},
	}
	return response
}
