package clickup

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Get the absolute path to the .env file, relative to the current file
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	envPath := filepath.Join(basePath, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file from %s", envPath)
	}
}

var authenticationToken string = os.Getenv("CLICKUP_AUTH_TOKEN")

func GetWorkspaces() {
	url := "https://api.clickup.com/api/v2/team"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", authenticationToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

func GetSpaces(space_id string) {

	url := fmt.Sprintf("https://api.clickup.com/api/v2/team/%s/space", space_id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", authenticationToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
