package types

import "errors"

// Identifier represents the unique identifier of the book
type Identifier struct {
	Type       string
	Identifier string
}

// Item represents a search result from the google book api
type Item struct {
	Kind       string
	ID         string
	VolumeInfo Volume
}

// ImageLinks represents a link to the thumbnail images
type ImageLinks struct {
	SmallThumbnail string
	Thumbnail      string
}

// Volume represents the book details
type Volume struct {
	Title               string
	Subtitle            string
	Authors             []string
	PublishedDate       string
	Description         string
	IndustryIdentifiers []Identifier
	ImageLinks          ImageLinks
}

// SearchResult represents the search result from the google api
type SearchResult struct {
	TotalItems int
	Kind       string
	Items      []Item
}

// Identifier will return the unique identifier of a book (either the ISBN_10 or ISBN_13 value)
func (v *Volume) Identifier() (string, error) {
	isbn := find(v.IndustryIdentifiers, func(i Identifier) bool {
		return "ISBN_13" == i.Type
	})

	if isbn == nil {
		isbn = find(v.IndustryIdentifiers, func(i Identifier) bool {
			return "ISBN_10" == i.Type
		})
	}

	if isbn == nil {
		return "", errors.New("Could not determine the unique identifier of the book")
	}
	return isbn.Identifier, nil
}

func find(identifiers []Identifier, fn func(i Identifier) bool) *Identifier {
	for _, indentifier := range identifiers {
		if fn(indentifier) {
			return &indentifier
		}
	}
	return nil
}
