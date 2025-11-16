package accounthttp

import (
	accountcommands "elastic-logger-app/modules/account/usecase/commands"
	accountqueries "elastic-logger-app/modules/account/usecase/queries"

	"github.com/gin-gonic/gin"
)

type accountHttp struct {
	cmd   accountcommands.Commands
	query accountqueries.Queries
}

func NewAccountHTTP(cmd accountcommands.Commands, query accountqueries.Queries) *accountHttp {
	return &accountHttp{
		cmd:   cmd,
		query: query,
	}
}

func (s *accountHttp) Routes(g *gin.RouterGroup) {
	acc_route := g.Group("/accounts")
	{
		acc_route.POST("", s.handleCreateAccount())
	}
}
