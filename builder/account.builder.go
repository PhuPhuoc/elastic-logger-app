package builder

import (
	"database/sql"
	accountcommandrepo "elastic-logger-app/modules/account/infras/commandrepo"
	accountqueryrepo "elastic-logger-app/modules/account/infras/queryrepo"
	accountcommands "elastic-logger-app/modules/account/usecase/commands"
	accountqueries "elastic-logger-app/modules/account/usecase/queries"

	"go.mongodb.org/mongo-driver/mongo"
)

type accountBuilder struct {
	db    *sql.DB
	mongo *mongo.Client
}

func NewAccountBuilder(db *sql.DB, mongo *mongo.Client) accountBuilder {
	return accountBuilder{db: db, mongo: mongo}
}

func (s accountBuilder) BuildAccountCommandRepo() accountcommands.AccountCommandRepo {
	return accountcommandrepo.NewAccountCommandRepo(s.db)
}

func (s accountBuilder) BuildAccountQueryRepo() accountqueries.AccountQueryRepo {
	return accountqueryrepo.NewAccountQueryRepo(s.mongo)
}
