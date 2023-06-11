package utils

import (
	"fmt"
	"math/rand"
)

var (
	adjectives = []string{
		"conventional",
		"fashionable",
		"democratic",
		"environmental",
		"disposable",
		"conservative",
		"autome",
		"professional",
		"diplomatic",
		"economical",
		"fantastical",
		"fastidious",
		"accessible",
		"deficient",
		"amplifying",
		"generous",
		"greedy",
		"satisfying",
		"separable",
		"cooperative",
		"anticipated",
		"relaxing",
		"reliable",
		"charismatic",
		"ordinary",
		"integrated",
		"original",
	}
	nouns = []string{
		"helicopter",
		"plane",
		"material",
		"community",
		"philosophy",
		"computer",
		"motorcycle",
		"vegetation",
		"competition",
		"economy",
		"aluminium",
		"factory",
		"architecture",
		"refrigerator",
		"dishwasher",
		"stovetop",
		"radiation",
		"education",
		"technology",
		"hypothesis",
		"curriculum",
		"publication",
		"newsletter",
		"stereotype",
		"machinery",
		"emergency",
		"excavation",
	}
)

func GenerateRandomName() string {
	adjectiveIx := getRandomIndex(len(adjectives))
	adjective := adjectives[adjectiveIx]

	nounIx := getRandomIndex(len(nouns))
	noun := nouns[nounIx]

	return fmt.Sprintf("%s-%s", adjective, noun)
}

func getRandomIndex(maxExcl int) int {
	index := rand.Int()

	return index % maxExcl
}
