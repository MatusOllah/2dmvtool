package rui

import (
	"math/rand/v2"
)

var quotes []string = []string{
	"High-pressure washing machine.",
	"Art is an explosion, unless it's a controlled hydraulic device.",
	"fufufu",
	"yoyoyo",
	"oyaoyaoya",
	"Yaa Tsukasa-kun~",
	"Go is better than Rust.",
	"Normal is overrated.",
	"With a snow making machine, we could turn the stage into a winter wonderland.",
	"My goal's always the same. Entertaining the crowd!",
	"A token of my appreciation. Although, it is nothing extravagant.",
	"I've added a bean-collecting feature to Robo-Nene that can shoot beans indefinitely.",
	"Frame rate isn't real. It's a social construct.",
}

var bdayQuotes []string = []string{
	"When was the last time I celebrated my own birthday in the company of friends...? Thank you, everyone.",
	"It's been quite some time since I last had a birthday party that included anyone other than my family. Hehe. I can't help but be excited about it.",
}

func RandomQuote() string {
	if IsBirthday() {
		return bdayQuotes[rand.IntN(len(bdayQuotes))]
	}

	return quotes[rand.IntN(len(quotes))]
}
