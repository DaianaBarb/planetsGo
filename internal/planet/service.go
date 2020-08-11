package planet

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	planets *mongo.Collection
}

func NewService(db *mongo.Database) *Service {
	return &Service{planets: db.Collection("planets")}
}

func (s *Service) Save(ctx context.Context, document *PlanetDocument)  error {

	one, err := s.planets.InsertOne(ctx, document)
	if err != nil {
		return errors.Wrap(err, "Erro ao salvar documento")
	}

	document.ID= one.InsertedID.(primitive.ObjectID)

	return nil
}


