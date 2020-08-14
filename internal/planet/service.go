package planet

import (
	"context"
	"projeto-star-wars-api-go/internal/model"
	"projeto-star-wars-api-go/internal/provider/mongo/dao"
	"projeto-star-wars-api-go/internal/provider/swapi"
)

type Service interface {
	Save(parentContext context.Context, in *model.PlanetIn) (string, error)
	FindAll(ctx context.Context) ([]model.PlanetOut, error)
	DeleteById(ctx context.Context, id string) error
	UpdateById(ctx context.Context, p model.PlanetIn, id string) (*model.PlanetOut, error)
	FindById(ctx context.Context, id string) (*model.PlanetOut, error)
	FindByName(ctx context.Context, name string) ([]model.PlanetOut, error)
}

type serviceImpl struct {
	planets dao.Planets
}

func NewService(planets dao.Planets) Service {
	return &serviceImpl{planets: planets}
}

func (s *serviceImpl) Save(ctx context.Context, in *model.PlanetIn) (string, error) {
	HexID, err := s.planets.Save(ctx, in)
	if err != nil {
		return "", err
	}
	return HexID, nil
}
func (s *serviceImpl) FindAll(ctx context.Context) ([]model.PlanetOut, error) {

	models, err := s.planets.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var newListModel []model.PlanetOut
	for _, m := range models {
		var saw swapi.SWAPI
		var number int
		number, _ = saw.CountPlanetAppearancesOnMovies(ctx, m.Name)
		m.NumberOfFilmAppearances = number
		newListModel = append(newListModel, m)
	}

	return newListModel, nil
}

func (s *serviceImpl) DeleteById(ctx context.Context, id string) error {

	err := s.planets.DeleteById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceImpl) UpdateById(ctx context.Context, p model.PlanetIn, id string) (*model.PlanetOut, error) {

	planOut, err := s.planets.UpdateById(ctx, p, id)
	if err != nil {
		return nil, err
	}
	return planOut, nil

}
func (s *serviceImpl) FindById(ctx context.Context, id string) (*model.PlanetOut, error) {
	planOut, err := s.planets.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	var saw swapi.SWAPI
	var number int
	number, _ = saw.CountPlanetAppearancesOnMovies(ctx, planOut.Name)
	planOut.NumberOfFilmAppearances = number
	return planOut, nil
}
func (s *serviceImpl) FindByName(ctx context.Context, name string) ([]model.PlanetOut, error) {

	models, err := s.planets.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}

	var newListModel []model.PlanetOut
	for _, m := range models {
		var saw swapi.SWAPI
		var number int
		number, _ = saw.CountPlanetAppearancesOnMovies(ctx, m.Name)
		m.NumberOfFilmAppearances = number
		newListModel = append(newListModel, m)
	}

	return newListModel, nil
}
