package util

import (
	"encoding/json"
	"fmt"
	"ocean_backend/models"
	"os"
)

func GetPath() string {
	CurrentPath, err := os.Getwd()
	if err != nil {
		return ""
	}

	return CurrentPath
}

func GetDatabaseFile() (models.Database, error) {

	currentDir := GetPath()
	databaseFolder := currentDir + "/models/database.json"

	jsonDatabaseFile, err := os.ReadFile(databaseFolder)
	if err != nil {
		fmt.Println("Could not find database file!")
		return nil, err
	}

	var databaseData models.Database
	json.Unmarshal(jsonDatabaseFile, &databaseData)

	return databaseData, nil

}
