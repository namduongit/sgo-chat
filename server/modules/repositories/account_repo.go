package repositories

import (
	"context"
	"github/sgo-chat/internals/configs/errors"
	"github/sgo-chat/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountRepository struct {
	profileRepository *ProfileRepository
	profileQueue      chan primitive.ObjectID

	collection *mongo.Collection
}

func NewAccountRepository(col *mongo.Collection, profileRepository *ProfileRepository) *AccountRepository {
	ar := &AccountRepository{
		collection:        col,
		profileRepository: profileRepository,
		// Initialize the channel with a buffer size, 100 is arbitrary
		profileQueue: make(chan primitive.ObjectID, 100),
	}
	go ar.processProfileQueue()
	return ar
}

func (ar *AccountRepository) FindByEmail(ctx context.Context, email string) (*models.Account, error) {
	result := ar.collection.FindOne(ctx, bson.M{"email": email})
	if result.Err() != nil {
		return nil, errors.BadRequestError("Not found email")
	}
	data := &models.Account{}
	err := result.Decode(data)
	return data, err
}

func (ar *AccountRepository) Create(ctx context.Context, dto *models.Account) (*models.Account, error) {
	result, err := ar.collection.InsertOne(ctx, dto)
	if err != nil {
		return nil, &errors.MongoError{Err: err}
	}

	// Set the ID from the insert result
	dto.ID = result.InsertedID.(primitive.ObjectID)

	ar.profileQueue <- dto.ID
	return dto, nil
}

func (ar *AccountRepository) processProfileQueue() {
	for accountID := range ar.profileQueue {
		err := ar.profileRepository.CreateByAccountID(accountID)
		if err != nil {
			log.Printf("ERROR: Failed to create profile for Account ID %s: %v\n", accountID.Hex(), err)
		} else {
			log.Printf("DEBUG: Profile created successfully for Account ID %s\n", accountID.Hex())
		}
	}
}
