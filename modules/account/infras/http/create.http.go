package accounthttp

import (
	"elastic-logger-app/common"
	accountcommands "elastic-logger-app/modules/account/usecase/commands"

	"github.com/gin-gonic/gin"
)

func (s *accountHttp) handleCreateAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto accountcommands.CreateAccountCmdDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		// if err := s.query.VerifyEmail.Handle(ctx, dto.Email); err != nil {
		// 	common.ResponseError(ctx, err)
		// 	return
		// }

		resp, err := s.cmd.CreateAccount.Handle(ctx, &dto)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, resp)
	}
}
