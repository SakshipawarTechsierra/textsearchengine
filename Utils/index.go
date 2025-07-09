package utils

// Index maps tokens to document IDs
type Index map[string][]int

// Add processes and adds all documents to the index
func (idx Index) Add(docs []Document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			idx[token] = append(idx[token], doc.Id)
		}
	}
}
