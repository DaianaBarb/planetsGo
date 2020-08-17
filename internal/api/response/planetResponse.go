package response

import (
	"projeto-star-wars-api-go/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlanetOut struct {
	ID                      primitive.ObjectID `json:"id"`
	Name                    string             `json:"name"`
	Climate                 string             `json:"climate"`
	Terrain                 string             `json:"terrain"`
	NumberOfFilmAppearances int                `json:"numberOfFilmAppearances"`
}

func (p *PlanetOut) FromModel(out model.PlanetOut) *PlanetOut {
	p.ID = out.ID
	p.Name = out.Name
	p.Climate = out.Climate
	p.Terrain = out.Terrain
	p.NumberOfFilmAppearances = out.NumberOfFilmAppearances
	return p
}
