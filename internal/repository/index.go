package repository

import (
	"encoding/json"
	"fmt"
)

type Index struct {
	Entries map[string]string `json:"entries"`
}

func NewIndex() *Index {
	return &Index{Entries: make(map[string]string)}
}

func (i *Index) Add(file, hash string) {
	i.Entries[file] = hash
}

func (i *Index) Load(data []byte) error {
	return json.Unmarshal(data, i)
}

func (i *Index) Save() ([]byte, error) {
	fmt.Println("Saving index")
	return json.Marshal(i)
}
