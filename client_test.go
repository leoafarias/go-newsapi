package newsapi_test

import (
	"strconv"
	"testing"

	"github.com/leoafarias/newsapi"
)

const (
	statusOK = "ok"
	succeed  = "\u2713"
	failed   = "\u2717"
	apikey   = ""
)

func TestTopHeadlines(t *testing.T) {
	c, err := newsapi.NewClient(apikey)

	if err != nil {
		t.Fatal(err)
	}

	params := make(map[string]string)

	params["country"] = "us"
	params["category"] = "business"

	res, err := c.TopHeadlines(params)
	if err != nil {
		t.Fatal(err)
	}

	if res.Status == statusOK {
		t.Logf("\t%s\t Should receive a %s status code", succeed, statusOK)
	} else {
		t.Errorf("\t%s\t Status was not ok, got: %s, want %s", failed, res.Status, statusOK)
	}
}

func TestEverything(t *testing.T) {
	c, err := newsapi.NewClient(apikey)

	query := "(ethereum OR litecoin OR bitcoin)"
	domains := "bbc.co.uk, techcrunch.com, engadget.com"
	pageSize := 2
	pageSizeString := strconv.Itoa(pageSize)

	if err != nil {
		t.Fatal(err)
	}

	params := make(map[string]string)

	params["q"] = query
	params["domains"] = domains
	params["pageSize"] = pageSizeString

	res, err := c.Everything(params)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Articles) == pageSize {
		t.Logf("\t%s\t Should have pageSize of %s", succeed, pageSizeString)
	} else {
		t.Errorf("\t%s\t Page size, got: %d, want %s", failed, len(res.Articles), pageSizeString)
	}

	if res.Status == statusOK {
		t.Logf("\t%s\t Should receive a %s status code", succeed, statusOK)
	} else {
		t.Errorf("\t%s\t Status was not ok, got: %s, want %s", failed, res.Status, statusOK)
	}
}

func TestSources(t *testing.T) {
	c, err := newsapi.NewClient(apikey)

	if err != nil {
		t.Fatal(err)
	}

	params := make(map[string]string)

	params["category"] = "technology"

	res, err := c.Sources(params)
	if err != nil {
		t.Fatal(err)
	}

	if res.Status == statusOK {
		t.Logf("\t%s\t Should receive a %s status code", succeed, statusOK)
	} else {
		t.Errorf("\t%s\t Status was not ok, got: %s, want %s", failed, res.Status, statusOK)
	}
}
