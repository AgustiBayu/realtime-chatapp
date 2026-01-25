package repository

import (
	"context"
	"database/sql"
	"realtime-chatapp/internal/domain"
)

type AuthRepositoyImpl struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) domain.AuthRepository {
	return &AuthRepositoyImpl{
		DB: db,
	}
}

func (r *AuthRepositoyImpl) Save(ctx context.Context, user domain.User) error {
	//this transaction
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	SQL := "INSERT INTO USERS (name, email, password) VALUES($1,$2,$3)"
	_, err = tx.ExecContext(ctx, SQL, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *AuthRepositoyImpl) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	SQL := "SELECT id, name, email, password FROM users WHERE email = $1"
	var user domain.User
	err := r.DB.QueryRowContext(ctx, SQL, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *AuthRepositoyImpl) FindById(ctx context.Context, id int) (domain.User, error) {
	SQL := "SELECT id, name, email, password FROM users WHERE id = $1"
	var user domain.User
	err := r.DB.QueryRowContext(ctx, SQL, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
