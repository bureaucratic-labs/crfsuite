# CRFSuite [![Build Status](https://travis-ci.org/bureaucratic-labs/crfsuite.svg?branch=master)](https://travis-ci.org/bureaucratic-labs/crfsuite)
Go bindings for CRFSuite

Things to be done:
* [ ] Training support
* [ ] Evaluation support (?)
* [x] Tagging support

# Tagging example

```go
package main

import (
	"fmt"
	"strings"
	"github.com/bureaucratic-labs/crfsuite"
)

// User-defined function, that returns features for each item in input sequence
// Interface is very similar (and based on) to python-crfsuite
func getFeatures(items []string, position int) []crfsuite.Feature {
	result := make([]Feature, 0)
	// Include lowercased value of item (actually, just char) as feature
	result = append(result, Feature{
		Key:   fmt.Sprintf("lower=%v", strings.ToLower(items[position])),
		Value: 1.0,
	})
	// There also can be more features, depending on your task
	return result
}

func main() {
	// Load pre-trained model for tokenization (see b-labs/models repo)
	model := NewModelFromFile("test_data/tokenization-model.crfsuite")
	tagger := model.GetTagger()
	// Input data must be an array of strings, but that can be changed in future
	input := []string{"т", "е", "с", "т", "."}
	ids := tagger.Tag(input, getFeatures)
	labels := tagger.IDsToLabels(ids)
	fmt.Println(labels) // will output some BIO labels
}
```
