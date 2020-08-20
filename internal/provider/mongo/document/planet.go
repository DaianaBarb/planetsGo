package document

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Planet struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Climate string             `bson:"climate"`
	Terrain string             `bson:"terrain"`
}
