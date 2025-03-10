package repository

import (
	"context"
	"my_website/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresUserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(c *pgxpool.Pool) domain.UserRepository {
	return &postgresUserRepository{db: c}
}

func (p *postgresUserRepository) CreateOne(user *domain.User) error {
	sql := `INSERT INTO users (email, password_hash, is_superuser) VALUES ($1, $2, $3)`
	_, err := p.db.Exec(context.Background(), sql, user.Email, user.PasswordHash, user.IsSuperuser)
	return err
}

func (p *postgresUserRepository) UpdateOneById(id string, user *domain.User) error {
	return nil
}

func (p *postgresUserRepository) DeleteOneById(id string) error {
	return nil
}

func (p *postgresUserRepository) GetOneByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := p.db.QueryRow(context.Background(), `SELECT * FROM users WHERE email=$1`, email).Scan(
		&user.Id,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsSuperuser,
	)
	return &user, err
}

func (p *postgresUserRepository) GetOneById(id string) (*domain.User, error) {
	return &domain.User{}, nil
}

func (p *postgresUserRepository) GetAll() ([]domain.User, error) {
	return []domain.User{}, nil
}
