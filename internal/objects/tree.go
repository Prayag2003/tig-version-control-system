package objects

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Prayag2003/tig/internal/utils"
)

type Tree struct {
	Hash    string      `json:"hash"`
	Entries []TreeEntry `json:"entries"`
}

type TreeEntry struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}

func NewTree(entries []TreeEntry) *Tree {
	tree := &Tree{Entries: entries}
	tree.Hash = utils.Hash([]byte(tree.String()))
	return tree
}

func (t *Tree) Save() error {
	treePath := filepath.Join(".tig", "objects", t.Hash[:2], t.Hash[2:])
	err := os.MkdirAll(filepath.Dir(treePath), 0755)
	if err != nil {
		return err
	}
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(treePath, data, 0644)
}

func (t *Tree) String() string {
	return fmt.Sprintf("entries: %v", t.Entries)
}
