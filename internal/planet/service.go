package planet

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	planets *mongo.Collection
}

func NewService(db *mongo.Database) *Service {
	return &Service{planets: db.Collection("planets")}
}

func (s *Service) Save(ctx context.Context, document *PlanetDocument) error {

	one, err := s.planets.InsertOne(ctx, document)
	if err != nil {
		return errors.Wrap(err, "Erro ao salvar documento")
	}

	document.ID = one.InsertedID.(primitive.ObjectID)

	return nil
}
func (s *Service) FindAll(ctx context.Context) ([]PlanetDocument,error) {

	result, err := s.planets.Find(ctx, bson.M{})
	if err != nil { // se o erro nao for nulo
		return nil, err
	} // se o erro for igual a a nullo ele n da erro. se o erro for nulo ele da erro
	var models []PlanetDocument // o erro tem que ser nullo para passar aqui e n retornar erro
	err = result.All(ctx, &models)
	if err != nil {
		return nil, err
	}

	return models, nil
}


func (s *Service) DeleteById(ctx context.Context,id string) error {
 oID, err := primitive.ObjectIDFromHex(id)
 if err != nil{
 	return err
 }
	_,err = s.planets.DeleteOne(ctx,bson.M{"_id":oID})
	return err
}

func (s *Service) UpdateById(ctx context.Context,p PlanetIn,id string) (*PlanetDocument, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	model := PlanetDocument{
		ID:      oID,
		Name:    p.Name,
		Climate: p.Climate,
		Terrain: p.Terrain,
	}
	opts := options.Update().SetUpsert(true)
	_, err = s.planets.UpdateOne(ctx, bson.M{"_id": model.ID}, bson.D{{"$set", model}}, opts)
	if err != nil {
		return nil, err
	}
	return &model, nil

}
func (s *Service) FindById(ctx context.Context,id string) (*PlanetDocument, error) {
	//**
	oID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	result := s.planets.FindOne(ctx, bson.M{"_id": oID})

	var model PlanetDocument
	err = result.Decode(&model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}
func (s *Service) FindByName(ctx context.Context,name string) (*PlanetDocument, error) {

	result := s.planets.FindOne(ctx, bson.M{"name": name})

	var model PlanetDocument
	err := result.Decode(&model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}