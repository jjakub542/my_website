package tests

import (
	"my_website/internal/domain"
	"my_website/internal/repository"
	"testing"
)

func TestArticleCreateOne(t *testing.T) {
	var err error

	article := &domain.Article{
		Title:       "tytuł artykułu",
		Description: "streszczenie/opis",
		Content:     "Writing a custom session middleware in Go Echo involves intercepting requests to handle session creation, validation, and management. Below is a complete guide to creating a basic custom session middleware.",
		Public:      true,
	}

	repo := repository.New(TestDB)
	err = repo.Article.CreateOne(article)
	if err != nil {
		t.Fatal(err)
	}
}

func TestArticleGetAll(t *testing.T) {
	repo := repository.New(TestDB)
	articles, err := repo.Article.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	if articles[0].Title != "tytuł artykułu" {
		t.Fail()
	}
	article2 := &domain.Article{
		Title:       "tytuł artykułu po aktualizacji",
		Description: "streszczenie/opis",
		Content:     "Writing a custom session middleware in Go Echo involves intercepting requests to handle session creation, validation, and management. Below is a complete guide to creating a basic custom session middleware.",
		Public:      true,
	}
	err = repo.Article.UpdateOneById(article2, articles[0].Id)
	if err != nil {
		t.Fatal(err)
	}
	articleFinal, err := repo.Article.GetOneById(articles[0].Id)
	if err != nil {
		t.Fatal(err)
	}
	if articleFinal.Title != "tytuł artykułu po aktualizacji" {
		t.Fail()
	}
}
