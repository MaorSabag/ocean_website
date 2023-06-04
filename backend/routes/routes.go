package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"ocean_backend/models"
	"ocean_backend/util"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func GetRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", testData).Methods("GET")
	router.HandleFunc("/database", getDatabase).Methods("GET")

	controller := handlers.CORS(
		handlers.AllowedOrigins([]string{"https://*.blunun.com"}),
		handlers.AllowedHeaders([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(router)

	return controller
}

func testData(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got / request from %v\n", r.RemoteAddr)

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

func getDatabase(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got /database request from %v\n", r.RemoteAddr)
	json_database, err := util.GetDatabaseFile()

	if err != nil {
		fmt.Println(err)
		error_resopsne, _ := json.Marshal(map[string]interface{}{
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
