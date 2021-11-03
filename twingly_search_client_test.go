package twingly_search_client

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	testClient, _ := New()

	got := reflect.TypeOf(testClient)
	want := reflect.TypeOf(client{})

	if got != want {
		t.Errorf("Expected a type of %q, got %q", want, got)
	}
}

func TestSetUserAgent(t *testing.T) {
	testClient, _ := New()

	testClient.SetUserAgent("Testing Company/1.0")

	got := testClient.UserAgent
	want := "Testing Company/1.0"

	if got != want {
		t.Errorf("Expected a type of %q, got %q", want, got)
	}
}

func TestAddQuery(t *testing.T) {
	testClient, _ := New()
	testClient.AddQuery("golang page-size:10")

	got := testClient.query["q"]
	want := []string{"golang page-size:10"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %q, got %q", reflect.ValueOf(want), reflect.ValueOf(got))
	}
}

func TestBuildUrl(t *testing.T) {
	t.Run("it builds url without a query", func(t *testing.T) {
		client, _ := New()
		builtUrl := client.buildUrl()

		got := builtUrl.String()
		want := "https://api.twingly.com/blog/search/api/v3/search?apikey="

		if got != want {
			t.Errorf("Expected %q, got %q", want, got)
		}
	})

	t.Run("it builds url with a query", func(t *testing.T) {
		client, _ := New()
		client.AddQuery("programming in golang page-size:1")
		builtUrl := client.buildUrl()

		got := builtUrl.String()
		want := "https://api.twingly.com/blog/search/api/v3/search?apikey=&q=programming+in+golang+page-size%3A1"

		if got != want {
			t.Errorf("Expected %q, got %q", want, got)
		}
	})
}

func TestExecuteQuery(t *testing.T) {
	client, _ := New()
	result, _ := client.ExecuteQuery()

	got := reflect.TypeOf(result)
	want := reflect.TypeOf(TwinglyResponse{})

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}
