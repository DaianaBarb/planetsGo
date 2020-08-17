package document

import (
	"projeto-star-wars-api-go/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlanetDocument struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Climate string             `bson:"climate"`
	Terrain string             `bson:"terrain"`
}

func (p *PlanetDocument) ToPlanetOut() *model.PlanetOut {
	return &model.PlanetOut{
		ID:                      p.ID,
		Name:                    p.Name,
		Climate:                 p.Climate,
		Terrain:                 p.Terrain,
		NumberOfFilmAppearances: 0,
	}
}

func (p *PlanetDocument) FromModel(in *model.PlanetIn) *PlanetDocument {
	return &PlanetDocument{
		Name:    in.Name,
		Climate: in.Climate,
		Terrain: in.Terrain,
	}
}
