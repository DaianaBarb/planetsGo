package planet

import (
	"context"
	"projeto-star-wars-api-go/internal/api/response"
	"projeto-star-wars-api-go/internal/model"
	"projeto-star-wars-api-go/internal/provider/mongo/dao"
	"projeto-star-wars-api-go/internal/provider/swapi"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	Save(parentContext context.Context, in *model.PlanetIn) (string, error)
	FindAll(ctx context.Context) (*[]response.PlanetOut, error)
	DeleteById(ctx context.Context, id string) error
	UpdateById(ctx context.Context, p model.PlanetIn, id string) (*response.PlanetOut, error)
	FindById(ctx context.Context, id string) (*response.PlanetOut, error)
	FindByName(ctx context.Context, name string) (*[]response.PlanetOut, error)
	Healthcheck() (*mongo.Database, error)
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
func (s *serviceImpl) FindAll(ctx context.Context) (*[]response.PlanetOut, error) {

	models, err := s.planets.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var newListModel []response.PlanetOut
	var planOut2 response.PlanetOut
	for _, m := range models {
		var saw swapi.SWAPI
		var number int
		number, _ = saw.CountPlanetAppearancesOnMovies(ctx, m.Name)
		m.NumberOfFilmAppearances = number
		planOut2.FromModel(m)
		newListModel = append(newListModel, planOut2)
	}

	return &newListModel, nil
}

func (s *serviceImpl) DeleteById(ctx context.Context, id string) error {

	err := s.planets.DeleteById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceImpl) UpdateById(ctx context.Context, p model.PlanetIn, id string) (*response.PlanetOut, error) {

	planOut, err := s.planets.UpdateById(ctx, p, id)
	if err != nil {
		return nil, err
	}
	var planOut2 response.PlanetOut

	planOut2.FromModel(*planOut)
	return &planOut2, nil

}
func (s *serviceImpl) FindById(ctx context.Context, id string) (*response.PlanetOut, error) {
	planOut, err := s.planets.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	var planOut2 response.PlanetOut

	planOut2.FromModel(*planOut)

	var saw swapi.SWAPI
	var number int
	number, _ = saw.CountPlanetAppearancesOnMovies(ctx, planOut2.Name)
	planOut.NumberOfFilmAppearances = number
	return &planOut2, nil
}
func (s *serviceImpl) FindByName(ctx context.Context, name string) (*[]response.PlanetOut, error) {

	models, err := s.planets.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}

	var newListModel []response.PlanetOut
	var planOut2 response.PlanetOut
	for _, m := range models {
		var saw swapi.SWAPI
		var number int
		number, _ = saw.CountPlanetAppearancesOnMovies(ctx, m.Name)
		m.NumberOfFilmAppearances = number
		planOut2.FromModel(m)
		newListModel = append(newListModel, planOut2)
	}

	return &newListModel, nil
}

func (s *serviceImpl) Healthcheck() (*mongo.Database, error) {

	cliente, err := s.planets.GetDatabase()
	if err != nil {
		return nil, err
	}

	return cliente, nil

}
