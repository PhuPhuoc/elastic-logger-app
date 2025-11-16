package accountcommandrepo

import (
	"context"
	"database/sql"
	accountdomain "elastic-logger-app/modules/account/domain"
	"elastic-logger-app/modules/account/infras/commandrepo/sqlc"

	_ "github.com/go-sql-driver/mysql"
)

type accountCommandRepo struct {
	db    *sql.DB
	store *sqlc.Queries
}

func NewAccountCommandRepo(db *sql.DB) *accountCommandRepo {
	return &accountCommandRepo{
		db:    db,
		store: sqlc.New(db),
	}
}

func (r *accountCommandRepo) Create(ctx context.Context, entity *accountdomain.Account) error {
	_, err := r.store.CreateAccount(ctx, sqlc.CreateAccountParams{
		ID:       entity.GetID(),
		Name:     entity.GetName(),
		Email:    entity.GetEmail(),
		Password: entity.GetPassword(),
		Status:   int(entity.GetStatus()),
	})
	return err
}
