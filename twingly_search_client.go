package twingly_search_client

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Twinglydata struct {
	XMLName                 xml.Name `xml:"twinglydata"`
	Text                    string   `xml:",chardata"`
	NumberOfMatchesReturned string   `xml:"numberOfMatchesReturned,attr"`
	SecondsElapsed          string   `xml:"secondsElapsed,attr"`
	NumberOfMatchesTotal    string   `xml:"numberOfMatchesTotal,attr"`
	IncompleteResult        string   `xml:"incompleteResult,attr"`
	Post                    []struct {
		Chardata     string `xml:",chardata"`
		ID           string `xml:"id"`
		Author       string `xml:"author"`
		URL          string `xml:"url"`
		Title        string `xml:"title"`
		Text         string `xml:"text"`
		LanguageCode string `xml:"languageCode"`
		LocationCode string `xml:"locationCode"`
		Coordinates  string `xml:"coordinates"`
		Links        string `xml:"links"`
		Tags         struct {
			Text string   `xml:",chardata"`
			Tag  []string `xml:"tag"`
		} `xml:"tags"`
		Images       string `xml:"images"`
		IndexedAt    string `xml:"indexedAt"`
		PublishedAt  string `xml:"publishedAt"`
		ReindexedAt  string `xml:"reindexedAt"`
		InlinksCount string `xml:"inlinksCount"`
		BlogId       string `xml:"blogId"`
		BlogName     string `xml:"blogName"`
		BlogUrl      string `xml:"blogUrl"`
		BlogRank     string `xml:"blogRank"`
		Authority    string `xml:"authority"`
	} `xml:"post"`
}

type TwinglyResponse struct {
	NumberOfMatchesReturned string
	SecondsElapsed          string
	NumberOfMatchesTotal    string
	IncompleteResult        string
	Post                    []Post
}

type Post struct {
	ID           string
	Author       string
	URL          string
	Title        string
	Text         string
	LanguageCode string
	LocationCode string
	Coordinates  string
	Links        string
	Tags         []Tag
	Images       string
	IndexedAt    string
	PublishedAt  string
	ReindexedAt  string
	InlinksCount string
	BlogId       string
	BlogName     string
	BlogUrl      string
	BlogRank     string
	Authority    string
}

type Tag struct {
	Text string
	Tag  []string
}

type client struct {
	scheme    string
	host      string
	path      string
	UserAgent string

	query url.Values
}

const (
	SCHEME             = "https"
	PATH               = "/blog/search/api/v3/search"
	DEFAULT_USER_AGENT = "Twingly Search Go Client/1.0"
)

func getOrigin() string {
	host := os.Getenv("ORIGIN")

	if host == "" {
		host = "api.twingly.com"
	}

	return host
}

func getApiKey() (string, error) {
	apikey := os.Getenv("TWINGLY_SEARCH_KEY")

	if apikey == "" {
		return apikey, errors.New("no API key has been provided")
	}

	return apikey, nil
}

func New() (client, error) {
	query := url.Values{}
	apikey, err := getApiKey()

	query.Add("apikey", apikey)

	return client{SCHEME, getOrigin(), PATH, DEFAULT_USER_AGENT, query}, err
}

func (c *client) SetUserAgent(userAgentName string) {
	c.UserAgent = userAgentName
}

func (c client) AddQuery(query string) {
	c.query.Add("q", query)
}

func (c client) buildUrl() url.URL {
	return url.URL{
		Scheme:   c.scheme,
		Host:     c.host,
		Path:     c.path,
		RawQuery: c.query.Encode(),
	}
}

func (c client) parseResponse(response Twinglydata) TwinglyResponse {
	var post []Post

	for _, postData := range response.Post {
		var aPost Post
		var aTag Tag
		var tags []Tag

		aPost.ID = postData.ID
		aPost.Author = postData.Author
		aPost.URL = postData.URL
		aPost.Title = postData.Title
		aPost.Text = postData.Text
		aPost.LanguageCode = postData.LanguageCode
		aPost.LocationCode = postData.LocationCode
		aPost.Coordinates = postData.Coordinates
		aPost.Links = postData.Links

		aTag.Text = postData.Tags.Text
		aTag.Tag = postData.Tags.Tag

		tags = append(tags, aTag)

		aPost.Tags = tags

		aPost.Images = postData.Images
		aPost.IndexedAt = postData.IndexedAt
		aPost.PublishedAt = postData.PublishedAt
		aPost.ReindexedAt = postData.ReindexedAt
		aPost.InlinksCount = postData.InlinksCount
		aPost.BlogId = postData.BlogId
		aPost.BlogName = postData.BlogName
		aPost.BlogUrl = postData.BlogUrl
		aPost.BlogRank = postData.BlogRank
		aPost.Authority = postData.Authority

		post = append(post, aPost)
	}

	var parsedResponse TwinglyResponse

	parsedResponse.NumberOfMatchesReturned = response.NumberOfMatchesReturned
	parsedResponse.SecondsElapsed = response.SecondsElapsed
	parsedResponse.NumberOfMatchesTotal = response.NumberOfMatchesTotal
	parsedResponse.IncompleteResult = response.IncompleteResult
	parsedResponse.Post = post

	return parsedResponse
}

func (c client) ExecuteQuery() (TwinglyResponse, error) {
	http_client := http.Client{}
	url := c.buildUrl()

	request, err := http.NewRequest("GET", url.String(), nil)

	request.Header.Set("User-Agent", c.UserAgent)

	if err != nil {
		return TwinglyResponse{}, fmt.Errorf("got error: %v", err)
	}

	response, err := http_client.Do(request)

	if err != nil {
		return TwinglyResponse{}, fmt.Errorf("got error: %v", err)
	}

	body, _ := ioutil.ReadAll(response.Body)

	if err != nil {
		return TwinglyResponse{}, fmt.Errorf("got error: %v", err)
	}

	var posts Twinglydata
	xml.Unmarshal(body, &posts)

	parsedResponse := c.parseResponse(posts)

	return parsedResponse, err
}
