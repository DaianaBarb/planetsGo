package service

import (
	"context"
	"projeto-star-wars-api-go/internal/model"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Planet interface {
	Save(parentContext context.Context, planet *model.PlanetIn) (string, error)
	DeleteById(ctx context.Context, id string) error
	Update(ctx context.Context, p model.PlanetIn, id string) error
	FindById(ctx context.Context, id string) (*model.PlanetOut, error)
	FindByParam(ctx context.Context, param *model.PlanetIn) ([]model.PlanetOut, error)
}

type PlanetRepository interface {
	Save(parentContext context.Context, planet *model.Planet) (string, error)
	DeleteById(ctx context.Context, id string) error
	Update(ctx context.Context, p *model.Planet, id string) error
	FindById(ctx context.Context, id string) (*model.Planet, error)
	FindByParam(ctx context.Context, param *model.PlanetIn) ([]model.Planet, error)
}

type planet struct {
	repository PlanetRepository
	swapi      SwapiInterface
}

func NewPlanet(planets PlanetRepository, swapi SwapiInterface) Planet {
	return &planet{repository: planets, swapi: swapi}
}

func (s *planet) Save(ctx context.Context, in *model.PlanetIn) (string, error) {
	planetModel := in.ToPlanet()
	HexID, err := s.repository.Save(ctx, planetModel)

	if err != nil {
		return "", err
	}

	return HexID, nil
}

func (s *planet) DeleteById(ctx context.Context, id string) error {

	err := s.repository.DeleteById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *planet) Update(ctx context.Context, p model.PlanetIn, id string) error {
	planet := p.ToPlanet()

	err := s.repository.Update(ctx, planet, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *planet) FindById(ctx context.Context, id string) (*model.PlanetOut, error) {
	planet, err := s.repository.FindById(ctx, id)
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

func (s *planet) FindByParam_exemplo1(ctx context.Context, param *model.PlanetIn) ([]model.PlanetOut, error) {

	planets, err := s.repository.FindByParam(ctx, param)

	if err != nil {
		return nil, err
	}
	var planetOuts []model.PlanetOut
	ch := make(chan model.PlanetOut, len(planets))
	//O waitgroup serve para fazer a thread principal esperar todas as goroutines terminarem
	var w sync.WaitGroup
	for _, planet := range planets {
		// significa que a concorrencia sera feia de um em um ele diz quantos devem esperar
		w.Add(1)
		planetOut := planet.ToPlanetOut()
		// uma goroutines
		go s.getCountPlanet(ctx, *planetOut, ch, &w)
	}
	// fechando a goroutines
	go waitTOClose(&w, ch)

	for planetOut := range ch {
		planetOuts = append(planetOuts, planetOut)
	}

	return planetOuts, nil
}
func waitTOClose(w *sync.WaitGroup, ch chan model.PlanetOut) {
	w.Wait()
	close(ch)
}
func (s *planet) getCountPlanet(ctx context.Context, p model.PlanetOut, ch chan model.PlanetOut, w *sync.WaitGroup) {
	// e executado no retorno do metodo
	defer w.Done()
	appearances, err := s.swapi.CountPlanetAppearancesOnMovies(ctx, p.Name)

	if err != nil {
		return
	}

	p.NumberOfFilmAppearances = appearances
	ch <- p
}
func (s *planet) FindByParam(ctx context.Context, param *model.PlanetIn) ([]model.PlanetOut, error) {
	planets, err := s.repository.FindByParam(ctx, param)
	if err != nil {
		return nil, err
	}
	var planetOuts []model.PlanetOut
	//ch := make(chan model.PlanetOut, len(planets))
	var wg sync.WaitGroup
	wg.Add(len(planets))
	for _, planet := range planets {
		planetOut := planet.ToPlanetOut()
		go func(planet model.PlanetOut) error {
			defer wg.Done()
			appearances, err := s.swapi.CountPlanetAppearancesOnMovies(ctx, planet.Name)
			if err != nil {
				return err
			}
			planet.NumberOfFilmAppearances = appearances
			return nil
		}(*planetOut)
		planetOuts = append(planetOuts, *planetOut)
	}
	wg.Wait()
	// n usei channel
	return planetOuts, nil
}

func (s *planet) FindByParam_exemplo2(parentCtx context.Context, param *model.PlanetIn) ([]model.PlanetOut, error) {

	planets, err := s.repository.FindByParam(parentCtx, param)
	if err != nil {
		return nil, err
	}
 // group
	g, ctx := errgroup.WithContext(parentCtx)
	defer ctx.Done()
 // channel
	ch := make(chan *model.PlanetOut, len(planets))

	g.Go(func() error {

		defer close(ch)
		childG, _ := errgroup.WithContext(ctx)

		for _, planet := range planets {
			planet := planet

			childG.Go(func() error {
				p := planet.ToPlanetOut()

				appearances, err := s.swapi.CountPlanetAppearancesOnMovies(ctx, p.Name)
				if err != nil {
					return err
				}

				p.NumberOfFilmAppearances = appearances

				ch <- p

				return nil
			})
		}

		return childG.Wait()
	})

	var planetOuts []model.PlanetOut
	g.Go(func() error {
		for p := range ch {
			planetOuts = append(planetOuts, *p)
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return planetOuts, nil
}
