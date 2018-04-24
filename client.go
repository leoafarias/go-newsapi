package newsapi

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1"
	userAgent      = "go-newsapi/" + libraryVersion
	apiURL         = "https://newsapi.org"
)

// Client exports
type Client struct {
	BaseURL    *url.URL
	APIKey     string
	httpClient http.Client
}

type params map[string]string

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

	err = json.NewDecoder(resp.Body).Decode(v)

	return resp, err
}

// TopHeadlines is great for retrieving headlines for display on news tickers or similar
func (c *Client) TopHeadlines(p params) (ArticlesResponse, error) {

	req, err := c.newRequest("GET", "/v2/top-headlines", p, nil)

	var res ArticlesResponse
	if err != nil {
		return res, err
	}

	_, err = c.do(req, &res)
	if err != nil {
		return res, err
	}

	if res.Status == "error" {
		return res, &apiError{res.Code, res.Message}
	}

	return res, nil
}

// Everything endpoint suits article discovery and analysis, but can be used to retrieve articles for display, too.
func (c *Client) Everything(p params) (ArticlesResponse, error) {
	req, err := c.newRequest("GET", "/v2/everything", p, nil)

	var res ArticlesResponse
	if err != nil {
		return res, err
	}
	_, err = c.do(req, &res)
	if err != nil {
		return res, err
	}

	if res.Status == "error" {
		return res, &apiError{res.Code, res.Message}
	}

	return res, nil
}

// Sources returns the subset of news publishers that top headlines
func (c *Client) Sources(p params) (SourcesResponse, error) {
	req, err := c.newRequest("GET", "/v2/sources", p, nil)

	var res SourcesResponse
	if err != nil {
		return res, err
	}
	_, err = c.do(req, &res)
	if err != nil {
		return res, err
	}

	if res.Status == "error" {
		return res, &apiError{res.Code, res.Message}
	}

	return res, nil
}
