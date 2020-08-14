package planet

//type PlanetIn struct {
//	Name    string `json:"name"`
//	Climate string `json:"climate"`
//	Terrain string `json:"terrain"`
//}
////type PlanetDocument struct {
////	ID      primitive.ObjectID `bson:"_id,omitempty"`
////	Name    string             `bson:"name"`
////	Climate string             `bson:"climate"`
////	Terrain string             `bson:"terrain"`
////}
//
//type PlanetOut struct {
//	ID                      primitive.ObjectID `json:"id"`
//	Name                    string             `json:"name"`
//	Climate                 string             `json:"climate"`
//	Terrain                 string             `json:"terrain"`
//	NumberOfFilmAppearances int                `json:"numberOfFilmAppearances"`
//}
//
//func (p *PlanetIn) ToDocument() *PlanetDocument {
//	return &PlanetDocument{
//		Name:    p.Name,
//		Climate: p.Climate,
//		Terrain: p.Terrain,
//	}
//}
//func (p *PlanetDocument) ToPlanetOut() *PlanetOut {
//	return &PlanetOut{
//		ID:                      p.ID,
//		Name:                    p.Name,
//		Climate:                 p.Climate,
//		Terrain:                 p.Terrain,
//		NumberOfFilmAppearances: 0,
//	}
//
//}
