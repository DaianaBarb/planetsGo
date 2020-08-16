package dao

import (
	"context"
	_ "projeto-star-wars-api-go/internal/api/response"
	"projeto-star-wars-api-go/internal/model"
	"projeto-star-wars-api-go/internal/provider/mongo/document"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Planets interface {
	Save(parentContext context.Context, planet *model.PlanetIn) (string, error)
	FindAll(ctx context.Context) ([]model.PlanetOut, error)
	DeleteById(ctx context.Context, id string) error
	UpdateById(ctx context.Context, p model.PlanetIn, id string) (*model.PlanetOut, error)
	FindById(ctx context.Context, id string) (*model.PlanetOut, error)
	FindByName(ctx context.Context, name string) ([]model.PlanetOut, error)
}

type planets struct {
	planets *mongo.Collection
}

func NewMongoPlanet(db *mongo.Database) Planets {
	return &planets{planets: db.Collection("planets")}
}

func (p planets) Save(parentContext context.Context, planet *model.PlanetIn) (string, error) {

	doc := new(document.PlanetDocument).FromModel(planet)

	one, err := p.planets.InsertOne(parentContext, doc)
	if err != nil {
		return "", err
	}

	inserted := one.InsertedID.(primitive.ObjectID)
	return inserted.Hex(), nil
}

func (p planets) FindAll(ctx context.Context) ([]model.PlanetOut, error) {
	result, err := p.planets.Find(ctx, bson.M{})
	if err != nil { // se o erro nao for nulo
		return nil, err
	} // se o erro for igual a a nullo ele n da erro. se o erro for nulo ele da erro
	var planets []document.PlanetDocument // o erro tem que ser nullo para passar aqui e n retornar erro
	err = result.All(ctx, &planets)
	if err != nil {
		return nil, err
	}

	//Transformar para model.Planeout
	var planetOut []model.PlanetOut
	for _, planet := range planets {
		// res := *planet.ToPlanetOut()
		planetOut = append(planetOut, *planet.ToPlanetOut())
	}

	return planetOut, err
}

func (p planets) DeleteById(ctx context.Context, id string) error {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = p.planets.DeleteOne(ctx, bson.M{"_id": oID})
	return nil
}

func (p planets) UpdateById(ctx context.Context, planetIn model.PlanetIn, id string) (*model.PlanetOut, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	m := document.PlanetDocument{
		ID:      oID,
		Name:    planetIn.Name,
		Climate: planetIn.Climate,
		Terrain: planetIn.Terrain,
	}
	opts := options.Update().SetUpsert(true)
	_, err = p.planets.UpdateOne(ctx, bson.M{"_id": m.ID}, bson.D{{"$set", m}}, opts)
	if err != nil {
		return nil, err
	}

	return m.ToPlanetOut(), nil
}

func (p planets) FindById(ctx context.Context, id string) (*model.PlanetOut, error) {

	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := p.planets.FindOne(ctx, bson.M{"_id": oID})

	var doc document.PlanetDocument
	err = result.Decode(&doc)
	if err != nil {
		return nil, err
	}

	return doc.ToPlanetOut(), nil
}

func (p planets) FindByName(ctx context.Context, name string) ([]model.PlanetOut, error) {
	result, err := p.planets.Find(ctx, bson.M{"name": name})
	if err != nil { // se o erro nao for nulo
		return nil, err
	}

	var models []document.PlanetDocument
	err = result.All(ctx, &models)
	if err != nil {
		return nil, err
	}

	//Transformar para model.Planeout
	var planetOut []model.PlanetOut
	for _, planet := range models {
		planetOut = append(planetOut, *planet.ToPlanetOut())
	}

	return planetOut, nil
}
