package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/sqweek/dialog"
	"github.com/SakshipawarTechsierra/textsearchengine/Utils"
)

func main() {
	filePath, err := dialog.File().Title("Select a .csv or .txt or .zip file").Load()
	if err != nil {
		log.Fatal("❌ File selection cancelled or failed:", err)
	}
	fmt.Println("📂 Selected file:", filePath)

	fmt.Print("🔍 Enter keyword to search: ")
	var keyword string
	fmt.Scanln(&keyword)
	query := strings.ToLower(keyword)
	keywords := utils.Filter(utils.Tokenize(query))

	var results []utils.MatchResult
	if strings.HasSuffix(filePath, ".csv") {
		results, _ = utils.SearchCSVFile(filePath, keywords)
	} else if strings.HasSuffix(filePath, ".txt") {
		results, _ = utils.SearchTXTFile(filePath, keywords)
	} else if strings.HasSuffix(filePath, ".zip") {
		results, _ = utils.SearchZipFile(filePath, keywords)
	} else {
		log.Fatal("❌ Unsupported file type. Only .csv, .txt and .zip allowed.")
	}

	if len(results) == 0 {
		fmt.Println("❌ No matches found.")
	} else {
		fmt.Printf("✅ Found %d matches:\n", len(results))
		for _, r := range results {
			fmt.Printf("📄 %s (Line %d):\n", r.FileName, r.RowNum)
			for _, field := range r.RowData {
				fmt.Printf("  %s\n", field)
			}
		}
	}
}
