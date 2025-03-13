package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitTables(db *pgxpool.Pool, path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	sql := string(content)
	_, err = db.Exec(context.Background(), sql)
	return err
}

func DropTables(db *pgxpool.Pool) error {
	var err error
	_, err = db.Exec(context.Background(), `
		DROP TABLE images;
		DROP TABLE articles;
		DROP TABLE users;
	`)

	if err != nil {
		return err
	}

	return nil
}

func InsertExampleArticles(db *pgxpool.Pool) error {
	_, err := db.Exec(context.Background(), `
INSERT INTO articles (title, description, content, public) VALUES
('Article 1', 'Description for Article 1', '# Article 1\nThis is some markdown content for Article 1.', TRUE),
('Article 2', 'Description for Article 2', '# Article 2\nThis is some markdown content for Article 2.', TRUE),
('Article 3', 'Description for Article 3', '# Article 3\nThis is some markdown content for Article 3.', TRUE),
('Article 4', 'Description for Article 4', '# Article 4\nThis is some markdown content for Article 4.', TRUE),
('Article 5', 'Description for Article 5', '# Article 5\nThis is some markdown content for Article 5.', TRUE),
('Article 6', 'Description for Article 6', '# Article 6\nThis is some markdown content for Article 6.', TRUE),
('Article 7', 'Description for Article 7', '# Article 7\nThis is some markdown content for Article 7.', TRUE),
('Article 8', 'Description for Article 8', '# Article 8\nThis is some markdown content for Article 8.', TRUE),
('Article 9', 'Description for Article 9', '# Article 9\nThis is some markdown content for Article 9.', TRUE),
('Article 10', 'Description for Article 10', '# Article 10\nThis is some markdown content for Article 10.', TRUE),
('Article 11', 'Description for Article 11', '# Article 11\nThis is some markdown content for Article 11.', TRUE),
('Article 12', 'Description for Article 12', '# Article 12\nThis is some markdown content for Article 12.', TRUE),
('Article 13', 'Description for Article 13', '# Article 13\nThis is some markdown content for Article 13.', TRUE),
('Article 14', 'Description for Article 14', '# Article 14\nThis is some markdown content for Article 14.', TRUE),
('Article 15', 'Description for Article 15', '# Article 15\nThis is some markdown content for Article 15.', TRUE),
('Article 16', 'Description for Article 16', '# Article 16\nThis is some markdown content for Article 16.', TRUE),
('Article 17', 'Description for Article 17', '# Article 17\nThis is some markdown content for Article 17.', TRUE),
('Article 18', 'Description for Article 18', '# Article 18\nThis is some markdown content for Article 18.', TRUE),
('Article 19', 'Description for Article 19', '# Article 19\nThis is some markdown content for Article 19.', TRUE),
('Article 20', 'Description for Article 20', '# Article 20\nThis is some markdown content for Article 20.', TRUE),
('Article 21', 'Description for Article 21', '# Article 21\nThis is some markdown content for Article 21.', TRUE),
('Article 22', 'Description for Article 22', '# Article 22\nThis is some markdown content for Article 22.', TRUE),
('Article 23', 'Description for Article 23', '# Article 23\nThis is some markdown content for Article 23.', TRUE),
('Article 24', 'Description for Article 24', '# Article 24\nThis is some markdown content for Article 24.', TRUE),
('Article 25', 'Description for Article 25', '# Article 25\nThis is some markdown content for Article 25.', TRUE),
('Article 26', 'Description for Article 26', '# Article 26\nThis is some markdown content for Article 26.', TRUE),
('Article 27', 'Description for Article 27', '# Article 27\nThis is some markdown content for Article 27.', TRUE);
	`)
	if err != nil {
		return err
	}
	return nil
}
