package calculations

import (
	"math"
)

type MatchResult struct {
	oponentRating int
	matchOutcome  float64
}

func CalculateElo(initalRating int, results []MatchResult) int {
	var kFactor float64 = 32

	var expectedScore float64 = 0

	for _, result := range results {
		expectedScore += calculateExpectedScore(initalRating, result.oponentRating)
	}

	var points float64 = 0
	for _, result := range results {
		points += result.matchOutcome
	}

	return int(float64(initalRating) + kFactor*(points-expectedScore))
}

func calculateExpectedScore(ratingA int, ratingB int) float64 {
	return 1 / (1 + math.Pow(10, (float64(ratingB-ratingA))/400))
}
