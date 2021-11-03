package main

import (
	"encoding/json"
	"fmt"
	"os"

	twinglySearchClient "github.com/chrizpy/twingly-search-api-go"
)

func main() {
	client, err := twinglySearchClient.New()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)

		os.Exit(1)
	}

	client.AddQuery("cop26 page-size:1")

	response, _ := client.ExecuteQuery()

	jsonResponse, _ := json.Marshal(response)

	fmt.Printf("%+v\n", string(jsonResponse))
}
