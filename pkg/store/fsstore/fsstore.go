package fsstore

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"io/fs"
	"kun-blog-golang/pkg/store"
	"kun-blog-golang/pkg/store/models"
	"kun-blog-golang/pkg/utils"
	"log"
	"os"
	"path/filepath"
)

type FSStore struct {
	root string
}

func (this *FSStore) List() (posts []*models.Post, err error) {
	posts = make([]*models.Post, 0)
	err = filepath.Walk(this.root, func(path string, info fs.FileInfo, err error) error {
		log.Println("path:", path, err)
		var f *os.File
		var postM *models.Post
		if fileStat, err := os.Stat(path); err != nil {
			return err
		} else if fileStat.IsDir() {
			return nil
		}
		if f, err = os.Open(path); err != nil {
			return err
		}
		defer f.Close()
		if postM, err = utils.ParseFileToPost(f); err != nil {
			return err
		}
		posts = append(posts, postM)
		return nil
	})
	return
}

func (F FSStore) DeleteBySlug(slug string) error {
	//TODO implement me
	panic("implement me")
}

func (F FSStore) GetBySlug(slug string) (*models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (this *FSStore) Save(post *models.Post) error {
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

	filePath := filepath.Join(this.root, post.FileName)
	return os.WriteFile(filePath, mdBuffer.Bytes(), 0600)
}

func NewFSStore(root string) *FSStore {
	return &FSStore{
		root: root,
	}
}

var _ store.StoreInterface = &FSStore{}
