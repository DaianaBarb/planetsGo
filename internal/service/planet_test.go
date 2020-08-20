package service

import (
	"context"
	"projeto-star-wars-api-go/internal/model"
	"projeto-star-wars-api-go/internal/provider/mongo/dao/mocks"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/stretchr/testify/mock"
)

func Test_planet_Save(t *testing.T) {
	type fields struct {
		dao   *mocks.Planet
		swapi SWAPI
	}
	type args struct {
		ctx context.Context
		in  *model.PlanetIn
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
		mock    func(repository *mocks.Planet)
	}{
		{
			name: "save sucess",
			fields: fields{
				dao:   new(mocks.Planet),
				swapi: NewSWAPI(),
			},
			args: args{
				ctx: context.Background(),
				in: &model.PlanetIn{
					Name:    mock.Anything,
					Climate: mock.Anything,
					Terrain: mock.Anything,
				},
			},
			want:    mock.Anything,
			wantErr: false,
			mock: func(repository *mocks.Planet) {
				repository.On("Save", mock.Anything, mock.Anything).Return(mock.Anything, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.dao)
			s := &planet{
				dao:   tt.fields.dao,
				swapi: tt.fields.swapi,
			}
			got, err := s.Save(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Save() got = %v, want %v", got, tt.want)
			}
			tt.fields.dao.AssertExpectations(t)
		})
	}
}

func Test_planet_DeleteById(t *testing.T) {
	type fields struct {
		dao   *mocks.Planet
		swapi SWAPI
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    error
		wantErr bool
		mock    func(repository *mocks.Planet)
	}{
		{
			name: "delete sucess",
			fields: fields{
				dao:   new(mocks.Planet),
				swapi: NewSWAPI(),
			},
			want:    nil,
			wantErr: false,
			mock: func(repository *mocks.Planet) {
				repository.On("DeleteById", mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.dao)
			s := &planet{
				dao:   tt.fields.dao,
				swapi: tt.fields.swapi,
			}
			if err := s.DeleteById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.fields.dao.AssertExpectations(t)
		})
	}
}

func Test_planet_FindById(t *testing.T) {
	idd := primitive.NewObjectID()
	type fields struct {
		dao   *mocks.Planet
		swapi SWAPI
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.PlanetOut
		wantErr bool
		mock    func(repository *mocks.Planet)
	}{
		{name: "success",
			fields: fields{
				dao:   new(mocks.Planet),
				swapi: NewSWAPI(),
			},
			args: args{
				ctx: context.Background(),
				id:  mock.Anything,
			},
			want: &model.PlanetOut{
				ID:                      idd,
				Name:                    mock.Anything,
				Climate:                 mock.Anything,
				Terrain:                 mock.Anything,
				NumberOfFilmAppearances: 0,
			},
			wantErr: false,
			mock: func(repository *mocks.Planet) {
				repository.On("FindById", mock.Anything, mock.Anything).Return(&model.Planet{
					ID:      idd,
					Name:    mock.Anything,
					Climate: mock.Anything,
					Terrain: mock.Anything,
				}, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.dao)
			s := &planet{
				dao:   tt.fields.dao,
				swapi: tt.fields.swapi,
			}
			got, err := s.FindById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindById() got = %v, want %v", got, tt.want)
			}
			tt.fields.dao.AssertExpectations(t)
		})
	}
}
