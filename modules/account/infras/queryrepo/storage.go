package accountqueryrepo

import "go.mongodb.org/mongo-driver/mongo"

type AccountQueryRepo struct {
	mongo *mongo.Client
}

func NewAccountQueryRepo(mongo *mongo.Client) *AccountQueryRepo {
	return &AccountQueryRepo{
		mongo: mongo,
	}
}
