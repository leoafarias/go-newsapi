# NewsAPI Go Client
[![Go Report Card](https://goreportcard.com/badge/github.com/leoafarias/go-newsapi)](https://goreportcard.com/report/github.com/leoafarias/go-newsapi) [![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)
[![stability-unstable](https://img.shields.io/badge/stability-unstable-red.svg)](https://github.com/emersion/stability-badges#unstable) 
[![Maintainability](https://api.codeclimate.com/v1/badges/1603870f2a43c27639e6/maintainability)](https://codeclimate.com/github/leoafarias/go-newsapi/maintainability)



A Go client for [NewsAPI v2](https://newsapi.org/docs)

## Installation

```bash
$ go get github.com/leoafarias/go-newsapi
```

## Usage

This library is a GO client you can use to interact with the [NewsAPI v2](https://newsapi.org/docs). Here are some examples

* TopHeadlines()
* Everything()
* Sources()


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
    params["category"] = "business"

    res, err := c.TopHeadlines(params)

    if err != nil {
        log.Fatal(err)
    }

    for _, a := range res.Articles {
        fmt.Println(a)
    }

}
```

## To Do
- [x] Implement TopHeadlines
- [x] Implement Everything
- [x] Implement Source
- [ ] Write tests
- [ ] Implement Cancelable requests