package tests

import (
	"my_website/internal/articles"
	"my_website/internal/database"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestArticleCreateOne(t *testing.T) {
	repo := articles.Repository(database.Test())
	article := &articles.Article{
		Title:   "Siema",
		Content: "tu jest tresc",
		Public:  true,
	}
	if err := repo.CreateOne(article); err != nil {
		t.Fatal(err)
	}
}

func TestArticleGetOneById(t *testing.T) {
	db := database.Test()
	a := &articles.Article{
		Id:      uuid.NewString(),
		Title:   "Siema",
		Content: "tu jest tresc",
		Public:  true,
	}
	_, err := db.Exec(`
	INSERT INTO articles (id, title, content, created_at, updated_at, public)
	VALUES ($1, $2, $3, $4, $5, $6)`, a.Id, a.Title, a.Content, time.Now(), time.Now(), a.Public)
	if err != nil {
		t.Fatal(err)
	}
	repo := articles.Repository(db)
	article, err := repo.GetOneById(a.Id)
	if err != nil {
		t.Fatal(err)
	}
	if article.Title != a.Title {
		t.Fail()
	}
}
