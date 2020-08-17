package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type PlanetIn struct {
	Name    string
	Climate string
	Terrain string
}

type PlanetOut struct {
	ID                      primitive.ObjectID
	Name                    string
	Climate                 string
	Terrain                 string
	NumberOfFilmAppearances int
}
