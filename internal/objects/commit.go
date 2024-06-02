package objects

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Prayag2003/tig/internal/utils"
)

type Commit struct {
	Hash    string `json:"hash"`
	Message string `json:"message"`
	Tree    string `json:"tree"`
	Parent  string `json:"parent"`
}

func NewCommit(message, tree, parent string) *Commit {
	commit := &Commit{Message: message, Tree: tree, Parent: parent}
	commit.Hash = utils.Hash([]byte(commit.String()))
	return commit
}

func (c *Commit) Save() error {
	commitPath := filepath.Join(".tig", "objects", c.Hash[:2], c.Hash[2:])
	err := os.MkdirAll(filepath.Dir(commitPath), 0755)
	if err != nil {
		return err
	}
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(commitPath, data, 0644)
}

func (c *Commit) Load(data []byte) error {
	return json.Unmarshal(data, c)
}

func (c *Commit) String() string {
	return fmt.Sprintf("tree: %s\nparent: %s\nmessage: %s", c.Tree, c.Parent, c.Message)
}
