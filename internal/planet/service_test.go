package planet

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"projeto-star-wars-api-go/internal/model"
	"projeto-star-wars-api-go/internal/provider/mongo/dao/mocks"
	"testing"
)

func Test_serviceImpl_Save(t *testing.T) {
	type fields struct {
		planets *mocks.Planets
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
		mock    func(repository *mocks.Planets)
	}{
		{
			name: "success",
			fields: fields{
				planets: new(mocks.Planets), //&mocks.Planets{}
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
			mock: func(repository *mocks.Planets) {
				repository.On("Save", mock.Anything, mock.Anything).Return(mock.Anything, nil).Once()
			},
		},
		{
			name: "fail when try to save planet",
			fields: fields{
				planets: new(mocks.Planets),
			},
			args: args{
				ctx: context.Background(),
				in: &model.PlanetIn{
					Name:    mock.Anything,
					Climate: mock.Anything,
					Terrain: mock.Anything,
				},
			},
			want:    "",
			wantErr: true,
			mock: func(repository *mocks.Planets) {
				repository.On("Save", mock.Anything, &model.PlanetIn{
					Name:    mock.Anything,
					Climate: mock.Anything,
					Terrain: mock.Anything,
				}).Return("", errors.New("error to save in database"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mock(tt.fields.planets)

			s := &serviceImpl{
				planets: tt.fields.planets,
			}
			got, err := s.Save(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Save() got = %v, want %v", got, tt.want)
			}

			tt.fields.planets.AssertExpectations(t)
		})
	}
}
