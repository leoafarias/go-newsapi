package newsapi_test

import (
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
