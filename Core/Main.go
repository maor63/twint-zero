package Core

import (
	"net/url"
	"strconv"
)

var (
	condition bool   = true
	cursor    string = ""
	tweet_count int = 0
	n_tweets int = 0
)

func Main(Query *string, Instance *string, Format *string, Limit *string) {
	(*Query) = url.QueryEscape(*Query)
	for condition {
		condition, n_tweets = Scrape(Request(Query, Instance, &cursor), Instance, Format, &cursor)
		tweet_count = tweet_count + n_tweets

        limit, err := strconv.Atoi(*Limit)
        if err == nil {
            if tweet_count >= limit{
                condition = false
            }
        }

	}
}
