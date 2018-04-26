# NewsAPI Go Client
[![Go Report Card](https://goreportcard.com/badge/github.com/leoafarias/go-newsapi)](https://goreportcard.com/report/github.com/leoafarias/go-newsapi) [![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)
[![stability-unstable](https://img.shields.io/badge/stability-stable-green.svg)](https://github.com/emersion/stability-badges#unstable) 
[![Maintainability](https://api.codeclimate.com/v1/badges/1603870f2a43c27639e6/maintainability)](https://codeclimate.com/github/leoafarias/go-newsapi/maintainability)



A Go client for [NewsAPI v2](https://newsapi.org/docs)

## Installation

```bash
$ go get github.com/leoafarias/go-newsapi
```

## Usage

This library is a GO client you can use to interact with the [NewsAPI v2](https://newsapi.org/docs). Here are some examples

### TopHeadlines
This method provides live top and breaking headlines for a country, specific category in a country, single source, or multiple sources. You can also search with keywords. Articles are sorted by the earliest date published first.


```go
package main

import (
    "fmt"
    "log"

    "github.com/leoafarias/go-newsapi"
)

const apikey = "API_KEY"

func main() {

    c, _ := newsapi.NewClient(apikey)

    params := make(map[string]string)

    params["country"] = "us"
    params["category"] = "technology"

    res, err := c.TopHeadlines(params)

    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }

    for _, a := range res.Articles {
        fmt.Println(a)
    }

}
```

### Everything

Search through millions of articles from over 30,000 large and small news sources and blogs. This includes breaking news as well as lesser articles.


```go
package main

import (
    "fmt"
    "log"

    "github.com/leoafarias/go-newsapi"
)

const apikey = "API_KEY"

func main() {

    c, _ := newsapi.NewClient(apikey)

    params := make(map[string]string)

    params["q"] = "(ethereum OR litecoin OR bitcoin)"
    params["domains"] = "bbc.co.uk, techcrunch.com, engadget.com"
    params["pageSize"] = "10"

    res, err := c.Everything(params)

    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }

    for _, a := range res.Articles {
        fmt.Println(a)
    }

}
```

### Sources

This method returns the subset of news publishers that top headlines (TopHeadlines) are available from. It's mainly a convenience endpoint that you can use to keep track of the publishers available on the API, and you can pipe it straight through to your users.


```go
package main

import (
    "fmt"
    "log"

    "github.com/leoafarias/go-newsapi"
)

const apikey = "API_KEY"

func main() {

    c, _ := newsapi.NewClient(apikey)

    params := make(map[string]string)

    params["category"] = "technology"

    res, err := c.Sources(params)

    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }

    for _, a := range res.Sources {
        fmt.Println(a)
    }

}
```

## To Do
- [x] Implement TopHeadlines
- [x] Implement Everything
- [x] Implement Source
- [x] Write tests
- [ ] Implement Cancelable requests