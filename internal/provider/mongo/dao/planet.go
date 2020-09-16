package dao

import (
	"context"
	"projeto-star-wars-api-go/internal/model"
	"projeto-star-wars-api-go/internal/provider/mongo/document"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Planet struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoPlanet(client *mongo.Client, db *mongo.Database) *Planet { //retorna a interface
	return &Planet{client: client, collection: db.Collection("planet")}
}

func (p *Planet) Save(parentContext context.Context, planet *model.Planet) (string, error) {

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

func (p *Planet) DeleteById(ctx context.Context, id string) error {
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

func (p *Planet) Update(ctx context.Context, planet *model.Planet, id string) error {
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

func (p *Planet) FindById(ctx context.Context, id string) (*model.Planet, error) {

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

	return ToPlanet(doc), nil
}

func (p *Planet) FindByParam(ctx context.Context, param *model.PlanetIn) ([]model.Planet, error) {

	filter := p.getFilter(param)

	result, err := p.collection.Find(ctx, filter)
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
		planet := ToPlanet(doc)

		planets = append(planets, *planet)
	}

	return planets, nil
}
func (p *Planet) getFilter(params *model.PlanetIn) bson.D {
	filter := bson.D{}

	filter = p.appendFilter("name", params.Name, &filter)
	filter = p.appendFilter("climate", params.Climate, &filter)
	filter = p.appendFilter("terrain", params.Terrain, &filter)

	return filter

}
func (p *Planet) appendFilter(field, value string, filter *bson.D) bson.D {
	if len(value) > 0 {
		*filter = append(*filter, bson.E{Key: field, Value: value})
	}

	return *filter
}
func (p *Planet) Check(ctx context.Context) error {
	ctx, _ = context.WithTimeout(ctx, time.Second)
	return p.client.Ping(ctx, nil)
}

func ToPlanet(p document.Planet) *model.Planet {

	return &model.Planet{
		ID:      p.ID,
		Name:    p.Name,
		Climate: p.Climate,
		Terrain: p.Terrain,
	}
}
