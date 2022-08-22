package fsstore

import (
	"kun-blog-golang/pkg/store/models"
	"reflect"
	"testing"
	"time"
)

func TestFSStore_DeleteBySlug(t *testing.T) {
	type fields struct {
		root string
	}
	type args struct {
		slug string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			F := FSStore{
				root: tt.fields.root,
			}
			if err := F.DeleteBySlug(tt.args.slug); (err != nil) != tt.wantErr {
				t.Errorf("DeleteBySlug() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFSStore_GetBySlug(t *testing.T) {
	type fields struct {
		root string
	}
	type args struct {
		slug string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			F := FSStore{
				root: tt.fields.root,
			}
			got, err := F.GetBySlug(tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBySlug() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBySlug() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSStore_List(t *testing.T) {
	type fields struct {
		root string
	}
	tests := []struct {
		name   string
		fields fields
		want   []*models.Post
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			F := FSStore{
				root: tt.fields.root,
			}
			if got := F.List(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSStore_Save(t *testing.T) {
	fs := NewFSStore("")
	err := fs.Save(&models.Post{
		PostConfig: &models.PostConfig{
			Slug:        "ABC123",
			Title:       "第一篇文章",
			CreatedTime: time.Now(),
			Author:      "Talon",
			Describe:    "文章描述",
		},
		Md:       "# 标题",
		FilePath: "./hello.md",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("Save success !")
}

func TestNewFSStore(t *testing.T) {
	type args struct {
		root string
	}
	tests := []struct {
		name string
		args args
		want *FSStore
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFSStore(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFSStore() = %v, want %v", got, tt.want)
			}
		})
	}
}
