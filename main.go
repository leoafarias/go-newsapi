package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	libraryVersion = "0.1"
	userAgent      = "go-newsapi/" + libraryVersion
	apiURL         = "https://newsapi.org"
	apikey         = ""
)

// Client exports
type Client struct {
	BaseURL    *url.URL
	APIKey     string
	httpClient http.Client
}

type articles []*article

type article struct {
	Source      source    `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
}

type source struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	Category    string `json:"category,omitempty"`
	Language    string `json:"language,omitempty"`
	Country     string `json:"country,omitempty"`
}

type articlesResponse struct {
	Status       string `json:"status"`
	Code         string `json:"code"`
	Message      string `json:"message"`
	TotalResults int    `json:"totalResults"`
	Articles     []article
}

type sourcesResponse struct {
	Status       string `json:"status"`
	Code         string `json:"code"`
	Message      string `json:"message"`
	TotalResults int    `json:"totalResults"`
	Sources      []source
}

type params map[string]string

const (
	statusOK    = "ok"
	statusError = "error"
)

// NewClient - Instantiates a new Client Struct
func NewClient(key string) (*Client, error) {

	url, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	client := &Client{
		APIKey:  key,
		BaseURL: url,
	}

	return client, nil
}

func (c *Client) newRequest(method, path string, p params, body interface{}) (*http.Request, error) {

	rel, err := url.Parse(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter

	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()

	for k, v := range p {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("X-Api-Key", c.APIKey)
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Printf("%v\n", resp.StatusCode)
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

func (c *Client) topHeadlines(p params) (articlesResponse, error) {

	req, err := c.newRequest("GET", "/v2/top-headlines", p, nil)

	var res articlesResponse
	if err != nil {
		return res, err
	}
	_, err = c.do(req, &res)
	if err != nil {

	}
	return res, nil
}
func (c *Client) everything() {}
func (c *Client) sources()    {}

func main() {
	c, _ := NewClient(apikey)
	params := make(map[string]string)
	// params["country"] = "us"
	res, err := c.topHeadlines(params)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", res)
}
