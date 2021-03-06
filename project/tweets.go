// +build ignore

package main

import (
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
	"strconv"
	"strings"
	"text/template"
)

func main() {
	anaconda.SetConsumerKey("")
	anaconda.SetConsumerSecret("")
	api := anaconda.NewTwitterApi("", "")

	values := url.Values{}
	values.Set("count", "1")
	values.Set("screen_name", "CodeWisdom")

	searchResult, _ := api.GetUserTimeline(values)
	maxId := searchResult[0].Id
	var quotes [2000]string
	maxQuotesIndex := 0

	values.Set("count", "200")

	for i := 0; i < 10; i++ {

		values.Set("max_id", strconv.FormatInt(maxId, 10))
		searchResult, _ := api.GetUserTimeline(values)

		for _, tweet := range searchResult {
			text := tweet.FullText

			if !strings.Contains(text, "RT") && strings.Contains(text, '"') && !strings.Contains(text, "http") && !strings.Contains(text, "@") && len(text) != 0 {
				quotes[maxQuotesIndex] = tweet.FullText
				maxQuotesIndex += 1
			}
			maxId = tweet.Id
		}
	}
	f, _ := os.Create("all_tweets.go")
	defer f.Close()
	packageTemplate.Execute(f, struct {
		Quotes [2000]string
		Max    int
	}{
		Quotes: quotes,
		Max:    maxQuotesIndex,
	})
}

var packageTemplate = template.Must(template.New("").Parse(`
package project

var maxItem = {{ .Max }}

var allTweets = []string{
{{- range .Quotes }}
  {{ printf "%q" . }},
{{- end }}
}
`))
