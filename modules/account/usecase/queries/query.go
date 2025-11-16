package accountqueries

type Queries struct {
}

type Builder interface {
	BuildAccountQueryRepo() AccountQueryRepo
}

func NewAccountQueryWithBuilder(b Builder) Queries {
	return Queries{}
}

type AccountQueryRepo interface {
}
