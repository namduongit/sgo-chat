package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
- Set time to connect to MongoDB is 10 seconds
- Set server selection timeout is 1 second
- If cannot connect to server in 1 second, return error

*/

func ConnectMongo(url string) (*mongo.Client, error) {
	// Set a timeout to connect to MongoDB. Max 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set client options with a shorter server selection timeout
	// Ping to server. if failed in 1 second, return error
	clientOpts := options.Client().
		ApplyURI(url).
		SetServerSelectionTimeout(1 * time.Second)

	// Connect to MongoDB
	// If timeout occurs, it will return an error
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	return client, nil

	/*
		End of function
		- cancel() will be called here due to defer
		- ctx will die here
	*/
}

/*
- collections: map[string][]string
  - key: collection name
  - value: list of fields to create unique index
  - Example:
    collections := map[string][]string{
    "accounts": {"email"},
    "users":    {"username", "phone"},
    }
*/
func SetupUniqueIndex(client *mongo.Client, name string, collections map[string][]string) error {
	database := client.Database(name)

	for collection, fields := range collections {
		collection := database.Collection(collection)

		for _, field := range fields {
			indexModel := mongo.IndexModel{
				// Example: map[string]int{"email": 1}. If field is "email" (Create object for index)
				Keys:    map[string]int{field: 1},
				Options: options.Index().SetUnique(true),
			}

			_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
