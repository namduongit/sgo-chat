package repositories

import (
	"context"
	"github/sgo-chat/internals/configs/errors"
	"github/sgo-chat/models"
	"github/sgo-chat/utils"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountRepository struct {
	bcrypt *utils.Bcrypt

	profileRepository *ProfileRepository
	profileQueue      chan primitive.ObjectID

	collection *mongo.Collection
}

func NewAccountRepository(col *mongo.Collection, bcrypt *utils.Bcrypt, profileRepository *ProfileRepository) *AccountRepository {
	ar := &AccountRepository{
		collection:        col,
		bcrypt:            bcrypt,
		profileRepository: profileRepository,
		profileQueue:      make(chan primitive.ObjectID, 100), // Initialize the channel with a buffer size, 100 is arbitrary
	}
	go ar.processProfileQueue()
	return ar
}

func (ar *AccountRepository) FindByEmail(ctx context.Context, email string) (*models.Account, error) {
	result := ar.collection.FindOne(ctx, bson.M{"email": email})
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, &errors.MongoError{Err: result.Err()}
	}
	data := &models.Account{}
	err := result.Decode(data)
	return data, err
}

func (ar *AccountRepository) Create(ctx context.Context, email string, password string) (*models.Account, error) {
	hashedPassword, err := ar.bcrypt.HashPassword(password)
	if err != nil {
		return nil, &errors.ServerError{Message: err.Error()}
	}
	data := &models.Account{
		Email:    email,
		Password: string(hashedPassword),
		CreateAt: time.Now(),
	}

	result, err := ar.collection.InsertOne(ctx, data)
	if err != nil {
		return nil, &errors.MongoError{Err: err}
	}

	// Set the ID from the insert result
	data.ID = result.InsertedID.(primitive.ObjectID)

	ar.profileQueue <- data.ID
	return data, nil
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
