package utils

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

// Document is the XML structure
type Document struct {
	Title string `xml:"title"`
	Url   string `xml:"url"`
	Text  string `xml:"abstract"`
	Id    int    // Set manually later
}

func LoadDocuments(path string) ([]Document, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	dec := xml.NewDecoder(gz)

	dump := struct {
		Documents []Document `xml:"doc"`
	}{}

	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}

	docs := dump.Documents
	for i := range docs {
		docs[i].Id = i
	}

	return docs, nil
}
