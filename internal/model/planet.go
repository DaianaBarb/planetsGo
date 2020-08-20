package model

import (
	"projeto-star-wars-api-go/internal/provider/mongo/document"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Planet struct {
	ID      primitive.ObjectID
	Name    string
	Climate string
	Terrain string
}

type PlanetIn struct {
	Name    string `json:"name"`
	Climate string `json:"climate"`
	Terrain string `json:"terrain"`
}

type PlanetOut struct {
	ID                      primitive.ObjectID `json:"id"`
	Name                    string             `json:"name"`
	Climate                 string             `json:"climate"`
	Terrain                 string             `json:"terrain"`
	NumberOfFilmAppearances int                `json:"numberOfFilmAppearances"`
}

func (p *PlanetIn) ToPlanet() *Planet {
	return &Planet{
		Name:    p.Name,
		Climate: p.Climate,
		Terrain: p.Terrain,
	}
}

func (p *Planet) ToPlanetOut() *PlanetOut {
	return &PlanetOut{
		ID:      p.ID,
		Name:    p.Name,
		Climate: p.Climate,
		Terrain: p.Terrain,
	}
}

func (p *Planet) ToDocument(id primitive.ObjectID) *document.Planet {
	return &document.Planet{
		ID:      id,
		Name:    p.Name,
		Climate: p.Climate,
		Terrain: p.Terrain,
	}
}
