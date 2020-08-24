package service

import (
	"context"
	"projeto-star-wars-api-go/internal/model"
	"projeto-star-wars-api-go/internal/provider/mongo/dao"
)

type Planet interface {
	Save(parentContext context.Context, planet *model.PlanetIn) (string, error)
	FindAll(ctx context.Context) ([]model.PlanetOut, error)
	DeleteById(ctx context.Context, id string) error
	Update(ctx context.Context, p model.PlanetIn, id string) error
	FindById(ctx context.Context, id string) (*model.PlanetOut, error)
	FindByName(ctx context.Context, name string) ([]model.PlanetOut, error)
}

type planet struct {
	dao   dao.Planet
	swapi SWAPI
}

func NewPlanet(planets dao.Planet, swapi SWAPI) Planet {
	return &planet{dao: planets, swapi: swapi}
}

func (s *planet) Save(ctx context.Context, in *model.PlanetIn) (string, error) {
	planetModel := in.ToPlanet()
	HexID, err := s.dao.Save(ctx, planetModel)

	if err != nil {
		return "", err
	}

	return HexID, nil
}

func (s *planet) FindAll(ctx context.Context) ([]model.PlanetOut, error) {

	planets, err := s.dao.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var planetOuts []model.PlanetOut

	for _, planet := range planets {
		planetOut := planet.ToPlanetOut()
		appearances, err := s.swapi.CountPlanetAppearancesOnMovies(ctx, planetOut.Name)

		if err != nil {
			return nil, err
		}

		planetOut.NumberOfFilmAppearances = appearances
		planetOuts = append(planetOuts, *planetOut)
	}
	return planetOuts, nil
}

func (s *planet) DeleteById(ctx context.Context, id string) error {

	err := s.dao.DeleteById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *planet) Update(ctx context.Context, p model.PlanetIn, id string) error {
	planet := p.ToPlanet()

	err := s.dao.Update(ctx, planet, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *planet) FindById(ctx context.Context, id string) (*model.PlanetOut, error) {
	planet, err := s.dao.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	planetOut := planet.ToPlanetOut()

	appearances, err := s.swapi.CountPlanetAppearancesOnMovies(ctx, planetOut.Name)
	if err != nil {
		return nil, err
	}

	planetOut.NumberOfFilmAppearances = appearances

	return planetOut, nil
}

func (s *planet) FindByName(ctx context.Context, name string) ([]model.PlanetOut, error) {

	planets, err := s.dao.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}

	var planetOuts []model.PlanetOut

	for _, planet := range planets {
		planetOut := planet.ToPlanetOut()
		appearances, err := s.swapi.CountPlanetAppearancesOnMovies(ctx, planetOut.Name)

		if err != nil {
			return nil, err
		}

		planetOut.NumberOfFilmAppearances = appearances
		planetOuts = append(planetOuts, *planetOut)
	}
	return planetOuts, nil
}
