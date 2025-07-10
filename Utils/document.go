package utils

import (
	"archive/zip"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type MatchResult struct {
	FileName string
	RowNum   int
	RowData  []string
}

func SearchCSVFile(filePath string, keywords []string) ([]MatchResult, error) {
	var results []MatchResult

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil || len(records) < 2 {
		return nil, err
	}

	headers := records[0]

	for i, row := range records[1:] {
		text := strings.ToLower(strings.Join(row, " "))
		tokens := Filter(Tokenize(text))
		if MatchesAny(tokens, keywords) {
			var named []string
			for j := 0; j < len(row) && j < len(headers); j++ {
				named = append(named, fmt.Sprintf("%-12s: %s", headers[j], row[j]))
			}
			results = append(results, MatchResult{
				FileName: filepath.Base(filePath),
				RowNum:   i + 2,
				RowData:  named,
			})
		}
	}
	return results, nil
}

func SearchTXTFile(filePath string, keywords []string) ([]MatchResult, error) {
	var results []MatchResult

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNum := 1
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		tokens := Filter(Tokenize(line))
		if MatchesAny(tokens, keywords) {
			results = append(results, MatchResult{
				FileName: filepath.Base(filePath),
				RowNum:   lineNum,
				RowData:  []string{scanner.Text()},
			})
		}
		lineNum++
	}
	return results, nil
}

func SearchZipFile(zipPath string, keywords []string) ([]MatchResult, error) {
	var results []MatchResult

	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	for _, f := range r.File {
		if strings.HasSuffix(f.Name, ".txt") || strings.HasSuffix(f.Name, ".csv") {
			rc, err := f.Open()
			if err != nil {
				continue
			}
			content, err := io.ReadAll(rc)
			rc.Close()
			if err != nil {
				continue
			}

			tmpFile := filepath.Join(os.TempDir(), f.Name)
			os.WriteFile(tmpFile, content, 0644)

			var partialResults []MatchResult
			if strings.HasSuffix(f.Name, ".txt") {
				partialResults, _ = SearchTXTFile(tmpFile, keywords)
			} else if strings.HasSuffix(f.Name, ".csv") {
				partialResults, _ = SearchCSVFile(tmpFile, keywords)
			}
			for i := range partialResults {
				partialResults[i].FileName = f.Name
			}
			results = append(results, partialResults...)
		}
	}
	return results, nil
}
