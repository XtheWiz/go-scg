package handler

import "testing"

import "net/url"

import "go-scg/internal/config"

import "strings"

func TestQueryString(t *testing.T) {
	testURL := "http://www.foo.com"
	destURL, _ := url.Parse(testURL)
	config, _ := config.LoadConfig("config/config.yml")

	t.Run("found foodtype in query string", func(t *testing.T) {
		queryString := prepareQueryString(destURL, config, "khai jeaw")
		expected := strings.Contains(queryString, "keyword=khai+jeaw")

		if !expected {
			t.Errorf("expected query string contains 'keyword=khai+jeaw', but got %s", queryString)
		}
	})
}
