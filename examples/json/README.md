# Twingly Search API Go - JSON Example

## Example code

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"

	twinglySearchClient "github.com/twingly/twingly-search-api-go"
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

```

## Example response

```json
{
  "NumberOfMatchesReturned": "1",
  "SecondsElapsed": "0.233",
  "NumberOfMatchesTotal": "110772",
  "IncompleteResult": "false",
  "Post": [
    {
      "ID": "16696306691441095642",
      "Author": "davidtaglieri",
      "URL": "https://dinamodinamo.wordpress.com/2021/11/03/cop-26-laccordo-sulle-foreste-ed-il-metano-di-cristina-calzecchi-onesti/",
      "Title": "Cop 26: l’accordo sulle foreste ed il metano( di Cristina Calzecchi Onesti)",
      "Text": "Il progetto della banca europea per la ricostruzione e lo sviluppo e dagli ...",
      "LanguageCode": "it",
      "LocationCode": "jp",
      "Coordinates": "",
      "Links": "",
      "Tags": [
        {
          "Text": "",
          "Tag": [
            "ambiente e amore per la natura",
            "geopolitica",
            "schema della redazione",
            "ambiente",
            "sensibilità ambientale",
            "sensibilità planetaria"
          ]
        }
      ],
      "Images": "",
      "IndexedAt": "2021-11-03T15:57:00Z",
      "PublishedAt": "2021-11-03T15:56:14Z",
      "ReindexedAt": "2021-11-03T15:57:00Z",
      "InlinksCount": "0",
      "BlogId": "14683730639065325819",
      "BlogName": "RADIO PENSIERO FM La Fonte quotidiana",
      "BlogUrl": "https://dinamodinamo.wordpress.com",
      "BlogRank": "1",
      "Authority": "1"
    }
  ]
}
```
