package main

import (
	"flag"
	"log"
	"time"

	"github.com/SakshipawarTechsierra/textsearchengine/Utils/index"
	"github.com/SakshipawarTechsierra/textsearchengine/Utils"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "NewCsv.zip", "wiki abstract dump path or zipped CSV")
	flag.StringVar(&query, "q", "small wild cat", "search query")
	flag.Parse()

	log.Println("Full text search is in progress...")

	start := time.Now()
	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	idx := index.New()
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	matchIDs := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchIDs), time.Since(start))

	for _, id := range matchIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}
