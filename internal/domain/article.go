package domain

import (
	"io"
	"mime/multipart"
	"os"
	"time"
)

type Image struct {
	Id         string    `json:"id"`
	Filename   string    `json:"filename"`
	UploadedAt time.Time `json:"uploaded_at"`
	ArticleId  string    `json:"article_id"`
}

type Article struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Public      bool      `json:"public"`
	Images      []Image   `json:"images"`
}

func (i *Image) Save(src multipart.File) error {
	dst, err := os.Create("web/static/uploads/" + i.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

func (i *Image) Remove() error {
	filepath := "web/static/uploads/" + i.Filename
	err := os.Remove(filepath)
	if err != nil {
		return err
	}
	return nil
}

func (i *Image) GetUrl() string {
	return "/static/uploads/" + i.Filename
}

type ArticleRepository interface {
	GetAll() ([]Article, error)
	GetAllPublic() ([]Article, error)
	GetAllPublicBetween(int, int) ([]Article, error)
	GetCount() (int, error)
	GetOneById(string) (*Article, error)
	CreateOne(*Article) error
	UpdateOneById(*Article, string) error
	DeleteOneById(string) error
	AttachImage(*Image, string) error
	RemoveImage(string) error
	GetArticleImages(string) ([]Image, error)
}
