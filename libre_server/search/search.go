package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	pb "github.com/crbaker/libre/libre"
	T "github.com/crbaker/libre/libre_server/search/types"
)

// Search will find search results for a keyword
func Search(keyword string) []*pb.Book {
	searchResult, _ := searchGoogleAPI(keyword)

	var books []*pb.Book

	for _, item := range searchResult.Items {
		if item.Kind == "books#volume" {
			book := volumeToBook(item.VolumeInfo)
			books = append(books, &book)
		}
	}

	return books
}

func volumeToBook(volume T.Volume) pb.Book {

	links := pb.ImageLink{
		SmallThumbnail: volume.ImageLinks.SmallThumbnail,
		Thumbnail:      volume.ImageLinks.Thumbnail,
	}

	var identifiers []*pb.Identifier

	for _, ident := range volume.IndustryIdentifiers {
		identifier := pb.Identifier{
			Identifier: ident.Identifier,
			Type:       ident.Type,
		}

		identifiers = append(identifiers, &identifier)
	}

	book := pb.Book{
		Title:               volume.Title,
		Description:         volume.Description,
		SubTitle:            volume.Subtitle,
		Authors:             volume.Authors,
		ImageLinks:          &links,
		IndustryIdentifiers: identifiers,
	}

	return book
}

func searchGoogleAPI(keyword string) (*T.SearchResult, error) {
	apikey := "AIzaSyAobwxOZis0BWKeSmSpEKJWdb3Nc0TAwtE"
	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=%s&key=%s", keyword, apikey)

	resp, err := http.Get(url)

	if err == nil {

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		return byteToSearchResult(body), nil

	}

	return new(T.SearchResult), err
}

func byteToSearchResult(b []byte) *T.SearchResult {
	var s T.SearchResult
	err := json.Unmarshal(b, &s)

	if err != nil {
		panic(err)
	}
	return &s
}
