package fsstore

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"kun-blog-golang/pkg/store"
	"kun-blog-golang/pkg/store/models"
	"os"
)

type FSStore struct {
	root string
}

func (F FSStore) List() []*models.Post {
	//TODO implement me
	panic("implement me")
}

func (F FSStore) DeleteBySlug(slug string) error {
	//TODO implement me
	panic("implement me")
}

func (F FSStore) GetBySlug(slug string) (*models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (F FSStore) Save(post *models.Post) error {
	var err error
	b, err := yaml.Marshal(post.PostConfig)
	if err != nil {
		return err
	}

	var mdBuffer bytes.Buffer
	mdBuffer.WriteString("---\n")
	mdBuffer.WriteString(string(b))
	mdBuffer.WriteString("---\n")
	mdBuffer.WriteString(post.Md)
	return os.WriteFile(post.FilePath, mdBuffer.Bytes(), 0600)
}

func NewFSStore(root string) *FSStore {
	return &FSStore{
		root: root,
	}
}

var _ store.StoreInterface = &FSStore{}
