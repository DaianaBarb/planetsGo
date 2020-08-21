package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestSWAPI_CountPlanetAppearancesOnMovies(t *testing.T) {
	swapi := NewSWAPI()
	type fields struct {
		APIURL string
	}
	type args struct {
		ctx        context.Context
		planetName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{name: "success",
			fields: fields{
				APIURL: swapi.APIURL,
			},
			args: args{
				ctx:        context.Background(),
				planetName: mock.Anything,
			},
			want:    0,
			wantErr: false,
		},
		{name: "success2",
			fields: fields{
				APIURL: swapi.APIURL,
			},
			args: args{
				ctx:        context.Background(),
				planetName: "Tatooine",
			},
			want:    5,
			wantErr: false,
		},
		{name: "error",
			fields: fields{
				APIURL: "bla bla bla",
			},
			args: args{
				ctx:        context.Background(),
				planetName: "Tatooine",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SWAPI{
				APIURL: tt.fields.APIURL,
			}
			got, err := s.CountPlanetAppearancesOnMovies(tt.args.ctx, tt.args.planetName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountPlanetAppearancesOnMovies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountPlanetAppearancesOnMovies() got = %v, want %v", got, tt.want)
			}

		})
	}
}
