package planet

import (
	"context"
	"projeto-star-wars-api-go/swapi"

	"go.mongodb.org/mongo-driver/mongo/options"

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

	var saw swapi.SWAPI
	var number int
	number, err := saw.CountPlanetAppearancesOnMovies(context.Background(), document.Name)
	document.NumberOfFilmAppearances = number
	one, err := s.planets.InsertOne(ctx, document)
	if err != nil {
		return err
	}

	document.ID = one.InsertedID.(primitive.ObjectID)

	return nil
}
func (s *Service) FindAll(ctx context.Context) ([]PlanetOut, error) {

	result, err := s.planets.Find(ctx, bson.M{})
	if err != nil { // se o erro nao for nulo
		return nil, err
	} // se o erro for igual a a nullo ele n da erro. se o erro for nulo ele da erro
	var models []PlanetOut // o erro tem que ser nullo para passar aqui e n retornar erro
	err = result.All(ctx, &models)
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (s *Service) DeleteById(ctx context.Context, id string) error {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = s.planets.DeleteOne(ctx, bson.M{"_id": oID})
	return err
}

func (s *Service) UpdateById(ctx context.Context, p PlanetIn, id string) (*PlanetDocument, error) {
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
func (s *Service) FindById(ctx context.Context, id string) (*PlanetOut, error) {
	//**
	oID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	result := s.planets.FindOne(ctx, bson.M{"_id": oID})

	var model PlanetOut
	err = result.Decode(&model)
	if err != nil {
		return nil, err
	}
	var saw swapi.SWAPI
	var number int
	number, _ = saw.CountPlanetAppearancesOnMovies(ctx, model.Name)
	model.NumberOfFilmAppearances = number
	//opts := options.Update().SetUpsert(true)
	_, err = s.planets.UpdateOne(ctx, bson.M{"_id": model.ID}, bson.D{{"$set", model}})
	if err != nil {
		return nil, err
	}
	return &model, nil
}
func (s *Service) FindByName(ctx context.Context, name string) (*PlanetDocument, error) {

	result := s.planets.FindOne(ctx, bson.M{"name": name})

	var model PlanetDocument
	err := result.Decode(&model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}
