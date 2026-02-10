package repositories

import (
	"context"
	"github/sgo-chat/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileRepository struct {
	collection *mongo.Collection
}

func NewProfileRepository(col *mongo.Collection) *ProfileRepository {
	return &ProfileRepository{
		collection: col,
	}
}

func (pr *ProfileRepository) CreateByAccountID(accountID primitive.ObjectID) error {
	_, err := pr.collection.InsertOne(context.Background(), models.Profile{
		AccountID: accountID,
		Username:  "",
		Avatar:    "",
		Fullname:  "",
		Bio:       "",
		Gender:    models.GenderUnspecified,
	})
	return err
}
