package dao

import (
	"context"
	"projeto-star-wars-api-go/internal/model"
	"projeto-star-wars-api-go/internal/provider/mongo/document"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Planet interface {
	Save(parentContext context.Context, planet *model.Planet) (string, error)
	FindAll(ctx context.Context) ([]model.Planet, error)
	DeleteById(ctx context.Context, id string) error
	UpdateById(ctx context.Context, p *model.Planet, id string) error
	FindById(ctx context.Context, id string) (*model.Planet, error)
	FindByName(ctx context.Context, name string) ([]model.Planet, error)
}

type planet struct {
	collection *mongo.Collection
}

func (p *planet) GetDatabase() (*mongo.Database, error) {
	return p.GetDatabase()
}

func NewMongoPlanet(db *mongo.Database) Planet {
	return &planet{collection: db.Collection("planet")}
}

func (p *planet) Save(parentContext context.Context, planet *model.Planet) (string, error) {

	doc := document.Planet{
		Name:    planet.Name,
		Climate: planet.Climate,
		Terrain: planet.Terrain,
	}

	one, err := p.collection.InsertOne(parentContext, doc)
	if err != nil {
		return "", err
	}

	inserted := one.InsertedID.(primitive.ObjectID)
	return inserted.Hex(), nil
}

func (p *planet) FindAll(ctx context.Context) ([]model.Planet, error) {
	result, err := p.collection.Find(ctx, bson.M{})
	if err != nil { // se o erro nao for nulo
		return nil, err
	} // se o erro for igual a a nullo ele n da erro. se o erro for nulo ele da erro

	var documents []document.Planet // o erro tem que ser nullo para passar aqui e n retornar erro
	err = result.All(ctx, &documents)
	if err != nil {
		return nil, err
	}

	var planets []model.Planet
	for _, planet := range documents {
		planets = append(planets, model.Planet{
			ID:      planet.ID,
			Name:    planet.Name,
			Climate: planet.Climate,
			Terrain: planet.Terrain,
		})
	}
	return planets, err
}

func (p *planet) DeleteById(ctx context.Context, id string) error {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = p.collection.DeleteOne(ctx, bson.M{"_id": oID})
	if err != nil {
		return err
	}

	return nil
}

func (p *planet) UpdateById(ctx context.Context, planet *model.Planet, id string) error {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	document := planet.ToDocument(oID)

	opts := options.Update().SetUpsert(true)

	_, err = p.collection.UpdateOne(ctx, bson.M{"_id": document.ID}, bson.D{{"$set", document}}, opts)

	if err != nil {
		return err
	}

	return nil
}

func (p *planet) FindById(ctx context.Context, id string) (*model.Planet, error) {

	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := p.collection.FindOne(ctx, bson.M{"_id": oID})

	var doc document.Planet
	err = result.Decode(&doc)
	if err != nil {
		return nil, err
	}

	return &model.Planet{
		ID:      doc.ID,
		Name:    doc.Name,
		Climate: doc.Climate,
		Terrain: doc.Terrain,
	}, nil
}

func (p *planet) FindByName(ctx context.Context, name string) ([]model.Planet, error) {
	result, err := p.collection.Find(ctx, bson.M{"name": name})
	if err != nil { // se o erro nao for nulo
		return nil, err
	}

	var documents []document.Planet
	err = result.All(ctx, &documents)
	if err != nil {
		return nil, err
	}

	var planets []model.Planet

	for _, doc := range documents {
		planet := model.Planet{
			ID:      doc.ID,
			Name:    doc.Name,
			Climate: doc.Climate,
			Terrain: doc.Terrain,
		}

		planets = append(planets, planet)
	}

	return planets, nil
}
