package errors

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoError struct {
	Err error
}

func (e *MongoError) Error() string {
	if e.Err != nil {
		if mongo.IsDuplicateKeyError(e.Err) {
			return "Resource already exists"
		}
		// Can't connect to server
		if mongo.IsNetworkError(e.Err) {
			return "Database connection error"
		}
		// Operation exceeded time limit. Example: time to connect is more than server selection timeout
		if mongo.IsTimeout(e.Err) {
			return "Database timeout error"
		}
	}

	return "Error in database"
}
