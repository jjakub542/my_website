package repository

import (
	"context"
	"my_website/internal/domain"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresArticleRepository struct {
	db *pgxpool.Pool
}

func NewArticleRepository(c *pgxpool.Pool) domain.ArticleRepository {
	return &postgresArticleRepository{db: c}
}

func (p *postgresArticleRepository) GetAll() ([]domain.Article, error) {
	var articles []domain.Article
	sql := `SELECT * FROM articles ORDER BY created_at DESC;`
	rows, err := p.db.Query(context.Background(), sql)
	if err != nil {
		return articles, err
	}
	for rows.Next() {
		var article domain.Article
		if err := rows.Scan(
			&article.Id,
			&article.Title,
			&article.Description,
			&article.Content,
			&article.CreatedAt,
			&article.UpdatedAt,
			&article.Public,
		); err != nil {
			return articles, err
		}
		articles = append(articles, article)
	}
	if err = rows.Err(); err != nil {
		return articles, err
	}
	return articles, nil
}

func (p *postgresArticleRepository) GetCount() (int, error) {
	var count int
	err := p.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM articles").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (p *postgresArticleRepository) GetAllPublicBetween(limit int, offset int) ([]domain.Article, error) {
	var articles []domain.Article
	sql := `SELECT * FROM articles WHERE public=true ORDER BY created_at DESC LIMIT $1 OFFSET $2;`
	rows, err := p.db.Query(context.Background(), sql, limit, offset)
	if err != nil {
		return articles, err
	}
	for rows.Next() {
		var article domain.Article
		if err := rows.Scan(
			&article.Id,
			&article.Title,
			&article.Description,
			&article.Content,
			&article.CreatedAt,
			&article.UpdatedAt,
			&article.Public,
		); err != nil {
			return articles, err
		}
		articles = append(articles, article)
	}
	if err = rows.Err(); err != nil {
		return articles, err
	}
	return articles, nil
}

func (p *postgresArticleRepository) GetAllPublic() ([]domain.Article, error) {
	var articles []domain.Article
	sql := `SELECT * FROM articles WHERE public=true ORDER BY created_at DESC;`
	rows, err := p.db.Query(context.Background(), sql)
	if err != nil {
		return articles, err
	}
	for rows.Next() {
		var article domain.Article
		if err := rows.Scan(
			&article.Id,
			&article.Title,
			&article.Description,
			&article.Content,
			&article.CreatedAt,
			&article.UpdatedAt,
			&article.Public,
		); err != nil {
			return articles, err
		}
		articles = append(articles, article)
	}
	if err = rows.Err(); err != nil {
		return articles, err
	}
	return articles, nil
}

func (p *postgresArticleRepository) GetOneById(id string) (*domain.Article, error) {
	var article domain.Article
	sql := `SELECT * FROM articles WHERE id=$1;`
	err := p.db.QueryRow(context.Background(), sql, id).Scan(
		&article.Id,
		&article.Title,
		&article.Description,
		&article.Content,
		&article.CreatedAt,
		&article.UpdatedAt,
		&article.Public,
	)
	return &article, err
}

func (p *postgresArticleRepository) CreateOne(a *domain.Article) error {
	sql := `INSERT INTO articles (title, description, content, public)
	VALUES ($1, $2, $3, $4)`
	_, err := p.db.Exec(context.Background(), sql, a.Title, a.Description, a.Content, a.Public)
	return err
}

func (p *postgresArticleRepository) UpdateOneById(a *domain.Article, id string) error {
	sql := `UPDATE articles SET
	title=$1, description=$2, content=$3, updated_at=$4, public=$5
	WHERE id=$6`
	_, err := p.db.Exec(context.Background(), sql, a.Title, a.Description, a.Content, time.Now(), a.Public, id)
	return err
}

func (p *postgresArticleRepository) DeleteOneById(id string) error {
	sql := `DELETE FROM articles WHERE id=$1`
	_, err := p.db.Exec(context.Background(), sql, id)
	return err
}

func (p *postgresArticleRepository) AttachImage(image *domain.Image, article_id string) error {
	sql := `INSERT INTO images (filename, article_id) VALUES ($1, $2)`
	_, err := p.db.Exec(context.Background(), sql, image.Filename, article_id)
	return err
}

func (p *postgresArticleRepository) RemoveImage(filename string) error {
	// filename should be indexed !!!
	sql := `DELETE FROM images WHERE filename=$1`
	_, err := p.db.Exec(context.Background(), sql, filename)
	return err
}

func (p *postgresArticleRepository) GetArticleImages(article_id string) ([]domain.Image, error) {
	var images []domain.Image
	sql := `SELECT * FROM images WHERE article_id=$1`
	rows, err := p.db.Query(context.Background(), sql, article_id)
	if err != nil {
		return images, err
	}
	for rows.Next() {
		var image domain.Image
		if err := rows.Scan(
			&image.Id,
			&image.Filename,
			&image.UploadedAt,
			&image.ArticleId,
		); err != nil {
			return images, err
		}
		images = append(images, image)
	}
	if err = rows.Err(); err != nil {
		return images, err
	}
	return images, nil
}
