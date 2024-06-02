package objects

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Prayag2003/tig/internal/utils"
)

type Blob struct {
	Hash string
	Data []byte
}

func NewBlob(data []byte) *Blob {
	hash := utils.Hash(data)
	return &Blob{Hash: hash, Data: data}
}

func (b *Blob) Save() error {
	objectPath := filepath.Join(".tig", "objects", b.Hash[:5], b.Hash[5:])
	err := os.MkdirAll(filepath.Dir(objectPath), 0755)
	if err != nil {
		return err
	}
	err = os.WriteFile(objectPath, b.Data, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Saved blob:", b.Hash)
	return nil
}
