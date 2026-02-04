package repositories

import "go.mongodb.org/mongo-driver/mongo"

type AccountRepository struct {
	colletion *mongo.Collection
}

func NewAccountRepository(col *mongo.Collection) *AccountRepository {
	return &AccountRepository{
		colletion: col,
	}
}
