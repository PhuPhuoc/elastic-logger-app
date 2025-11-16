package accountcommands

import (
	"context"
	accountdomain "elastic-logger-app/modules/account/domain"
)

type Commands struct {
	CreateAccount *createAccountHandler
}

type Builder interface {
	BuildAccountCommandRepo() AccountCommandRepo
}

func NewAccountCmdWithBuilder(b Builder) Commands {
	return Commands{
		CreateAccount: NewCreateAccountHandler(
			b.BuildAccountCommandRepo(),
		),
	}
}

type AccountCommandRepo interface {
	Create(ctx context.Context, entity *accountdomain.Account) error
}
