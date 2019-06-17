package main

import (
	"encoding/json"
	"fmt"
)

type FilterContainer struct {
	AdFilter     *AdFilterResult
	ObjectFilter *ObjectFilterResult
}

type AdFilterResult struct {
	Ad string
}

type ObjectFilterResult struct {
	Objects []string
}

func main() {
	c := &FilterContainer{
		AdFilter:     &AdFilterResult{Ad: "normal"},
		ObjectFilter: &ObjectFilterResult{[]string{"shoe", "gun"}},
	}
	fmt.Println(c)

}

func insertMongoDB(collectionName string, doc interface{}) {
	data, err := json.Marshal(doc)
	if err != nil {
		panic(err)
	}
	fmt.Println(collectionName, string(data))
}
