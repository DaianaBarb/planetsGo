package planet

import (
	"context"
	"errors"
	"projeto-star-wars-api-go/internal/model"
	"projeto-star-wars-api-go/internal/provider/mongo/dao/mocks"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/stretchr/testify/mock"
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
func Test_serviceImpl_FindById(t *testing.T) {
	idd := primitive.NewObjectID()

	type fields struct {
		planets *mocks.Planets //Faltou mudar o tipo de dao.Planets para *mocks.Planets. Temos que mudar para passar o mock e e podermos definir o retorno necessário
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
		mock    func(repository *mocks.Planets)
	}{

		{
			name: "success",
			fields: fields{
				planets: new(mocks.Planets), //&mocks.Planets{}
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
			mock: func(repository *mocks.Planets) {
				repository.On("FindById", mock.Anything, mock.Anything).Return(&model.PlanetOut{
					ID:                      idd,
					Name:                    mock.Anything,
					Climate:                 mock.Anything,
					Terrain:                 mock.Anything,
					NumberOfFilmAppearances: 0,
				}, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mock(tt.fields.planets) //Faltou adicionar essa chamada do método para popular a variável repository no field 	mock: func(repository *mocks.Planets)

			s := &serviceImpl{
				planets: tt.fields.planets,
			}
			got, err := s.FindById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindById() got = %v, want %v", got, tt.want)
			}

			tt.fields.planets.AssertExpectations(t) // Faltou adicionar esse assert para do mock para compararmos se a resposta está de acordo com o que queremos
		})
	}
}

func Test_serviceImpl_DeleteById(t *testing.T) {
	type fields struct {
		planets *mocks.Planets
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
		mock    func(repository *mocks.Planets)
	}{
		{
			name: "success",
			fields: fields{
				planets: new(mocks.Planets), //&mocks.Planets{}
			},
			args: args{
				ctx: context.Background(),
				id:  mock.Anything,
			},
			want:    nil,
			wantErr: false,
			mock: func(repository *mocks.Planets) {
				repository.On("DeleteById", mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.planets)
			s := &serviceImpl{
				planets: tt.fields.planets,
			}
			if err := s.DeleteById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.fields.planets.AssertExpectations(t)
		})
	}
}

func Test_serviceImpl_UpdateById(t *testing.T) {
	id:= primitive.NewObjectID().Hex()
	idd,_:= primitive.ObjectIDFromHex(id)
	type fields struct {
		planets *mocks.Planets
	}
	type args struct {
		ctx context.Context
		p   model.PlanetIn
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.PlanetOut
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
				p: model.PlanetIn{
					Name: mock.Anything,
	                Climate : mock.Anything,
	                Terrain :mock.Anything,
				},
				id:  id,
			},
			want:    &model.PlanetOut{
				ID                  :       idd,
	            Name                :    mock.Anything,
	            Climate             :    mock.Anything,
	            Terrain             :    mock.Anything,
	            NumberOfFilmAppearances :0,
			},
			wantErr: false,
			mock: func(repository *mocks.Planets) {
				repository.On("UpdateById", mock.Anything, mock.Anything,mock.Anything).Return(&model.PlanetOut{
				ID                  :       idd,
	            Name                :    mock.Anything,
	            Climate             :    mock.Anything,
	            Terrain             :    mock.Anything,
	            NumberOfFilmAppearances :0,
			},nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.planets)
			s := &serviceImpl{
				planets: tt.fields.planets,
			}
			got, err := s.UpdateById(tt.args.ctx, tt.args.p, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateById() got = %v, want %v", got, tt.want)
			}
			tt.fields.planets.AssertExpectations(t)
		})
	}
}