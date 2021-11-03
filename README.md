# Twingly Search API Go

A Go package for Twingly's Blog Search API. Twingly is a blog search service that provides a searchable API known as [Twingly Blog Search API](https://developer.twingly.com/resources/search/)

I did this for fun to try and learn the Go language, figured it was a good idea to have some sort of project alongside my Go journey.

## Installation

Import the package into your code:

```Go
import (
    twinglySearchClient "github.com/chrizpy/twingly-search-api-go"
)

func main(){
    ...
}
```

And install it.

```shell
go get github.com/chrizpy/twingly-search-api-go
```

## Usage

The twingly-search package talks to a commercial API and requires an API key. Set the TWINGLY_SEARCH_KEY environment variable to the obtained key.

To learn more about the capabilities of Twingly's APIs, please read the [Blog Search API documentation](https://developer.twingly.com/resources/search/).

## Blog Search API

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"

    twinglySearchClient "github.com/twingly/twingly-search-api-go"
)

func main() {
    client, _ := twinglySearchClient.New()
    client.SetUserAgent("MyCompany/1.0")
    client.AddQuery("cop26 page-size:1")


    response, _ := client.ExecuteQuery()
}
```

## Requirements

* API key, [sign up](https://www.twingly.com/try-for-free) via [twingly.com](https://www.twingly.com/) to get one
* Go >= 1.17 (Earlier versions probably work, just developed at 1.17)

## Development

### Tests

Run all tests

```shell
go test ./...
```

## TODO

This package is still very much in its infancy and could be lacking features compared to its [sibling libraries](https://github.com/twingly)
