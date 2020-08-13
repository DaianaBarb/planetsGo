package planet

import "go.mongodb.org/mongo-driver/bson/primitive"

type PlanetIn struct {
	Name    string `json:"name"`
	Climate string `json:"climate"`
	Terrain string `json:"terrain"`
}
type PlanetDocument struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty"`
	Name                    string             `bson:"name"`
	Climate                 string             `bson:"climate"`
	Terrain                 string             `bson:"terrain"`
	NumberOfFilmAppearances int                `bson:"numberOfFilmAppearances"`
}
type PlanetOut struct {
	ID                      string `json:"id"`
	Name                    string `json:"name"`
	Climate                 string `json:"climate"`
	Terrain                 string `json:"terrain"`
	numberOfFilmAppearances int    `bson:"numberOfFilmAppearances"`
}

func (p *PlanetIn) ToDocument() *PlanetDocument {
	return &PlanetDocument{
		Name:                    p.Name,
		Climate:                 p.Climate,
		Terrain:                 p.Terrain,
		NumberOfFilmAppearances: 0,
	}
}
func (p *PlanetDocument) ToPlanetOut() *PlanetOut {
	return &PlanetOut{
		ID:                      primitive.ObjectID.String(p.ID),
		Name:                    p.Name,
		Climate:                 p.Climate,
		Terrain:                 p.Terrain,
		numberOfFilmAppearances: p.NumberOfFilmAppearances,
	}

}
