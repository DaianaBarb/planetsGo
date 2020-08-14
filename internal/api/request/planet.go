package request

import "projeto-star-wars-api-go/internal/model"

type PlanetIn struct {
	Name    string `json:"name"`
	Climate string `json:"climate"`
	Terrain string `json:"terrain"`
}

func (p *PlanetIn) ToModel() *model.PlanetIn {
	return &model.PlanetIn{
		Name:    p.Name,
		Climate: p.Climate,
		Terrain: p.Terrain,
	}
}
