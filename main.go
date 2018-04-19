package main

import (
	"bytes"
	"encoding/json"
	"errors"
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
	apiURL         = "https://newsapi.org/v2"
	apikey         = ""
)

var (
	// ErrUnauthorized can be returned on any call on response status code 401.
	ErrUnauthorized = errors.New("newsapi: unauthorized")
)

// Client exports
type Client struct {
	BaseURL    *url.URL
	APIKey     string
	httpClient *http.Client
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
	Status       string
	TotalResults int
	Articles     []article
}

type sourcesResponse struct {
	Status       string
	TotalResults int
	Sources      []source
}

type params map[string]string

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {

	url, err := url.Parse(apiURL)
	if err != nil {
		log.Fatal(err)
	}
	rel, err := url.Parse(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", rel)
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
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("X-Api-Key", apikey)
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

func (c *Client) topHeadlines(p params) (articlesResponse, error) {
	v := url.Values{}
	fmt.Printf("%v\n", v)
	// for k, v := range p {

	// }

	req, err := c.newRequest("GET", "/top-headlines", nil)
	var res articlesResponse
	if err != nil {
		return res, err
	}
	_, err = c.do(req, &res)
	return res, nil
}
func (c *Client) everything() {}
func (c *Client) sources()    {}

func main() {

	client := &Client{
		APIKey: apikey,
	}

	params := make(map[string]string)
	params["country"] = "us"
	res, _ := client.topHeadlines(params)
	fmt.Printf("%v\n", res)
}