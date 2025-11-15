package accountqueryrepo

import "go.mongodb.org/mongo-driver/mongo"

type accountCommandRepo struct {
	mongo *mongo.Client
}

func NewAccountRepo(mongo *mongo.Client) *accountCommandRepo {
	return &accountCommandRepo{
		mongo: mongo,
	}
}
