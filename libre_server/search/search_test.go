package search

import (
	"testing"
)

func Test_SearchForBookWithKeyword(t *testing.T) {

	searchResults, _ := Search("978-0393338102")

	assertIntGreaterThan(len(searchResults.Items), 0, t)
}

func assertIntGreaterThan(actual int, expected int, t *testing.T) {
	if actual <= expected {
		t.Error(actual, expected)
	}
}

func assertIntEquals(actual int, expected int, t *testing.T) {
	if actual != expected {
		t.Error(actual, expected)
	}
}
