package project

import (
	"math/rand"
	"time"
)

//go:generate go run tweets.go

func PrintRandomTweet() string {
	rand.Seed(time.Now().Unix())
	randomIndex := rand.Intn(maxItem)
	return allTweets[randomIndex]
}
