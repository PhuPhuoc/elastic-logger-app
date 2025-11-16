package accountcommands

import (
	"context"
	"elastic-logger-app/common"
	accountdomain "elastic-logger-app/modules/account/domain"
)

type CreateAccountCmdDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createAccountHandler struct {
	commandrepo AccountCommandRepo
}

func NewCreateAccountHandler(cmdRepo AccountCommandRepo) *createAccountHandler {
	return &createAccountHandler{
		commandrepo: cmdRepo,
	}
}

type ResponseCreateAccountDTO struct {
	Id string `json:"id"`
}

func (h *createAccountHandler) Handle(ctx context.Context, dto *CreateAccountCmdDTO) (*ResponseCreateAccountDTO, error) {
	accid := common.GenUUID()
	entity, _ := accountdomain.NewAccount(
		accid.String(),
		dto.Name,
		dto.Email,
		dto.Password,
		accountdomain.StatusActivated,
		nil,
	)

	if err := h.commandrepo.Create(ctx, entity); err != nil {
		return nil, common.NewInternalServerError("cannot create new account", "cannot get insert account into db")
	}

	response := &ResponseCreateAccountDTO{
		Id: accid.String(),
	}

	return response, nil
}
