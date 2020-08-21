package router

import (
	"context"
	"net/http"
	"projeto-star-wars-api-go/internal/service/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestPlanetHandler_SavePlanet(t *testing.T) {
	type fields struct {
		service *mocks.Planet
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantErr       bool
		mock          func(fs *mocks.Planet)
	}{
		{name: "sucesss",
			fields: fields{
				service: new(mocks.Planet),
			},
			args: args{r:,
				w: },
				wantErr: false,
				mock: func(fs *mocks.Planet) {
				fs.On("SavePlanet", mock.Anything, mock.Anything).
					Return(nil)
			},

		},
	}
	for _, tt := range tests {
		tt.mock(tt.fields.service)
		t.Run(tt.name, func(t *testing.T) {
			p := &PlanetHandler{
				service: tt.fields.service,
			}
			//err := p.SavePlanet(tt.args.w, tt.args.r)


			tt.fields.service.AssertExpectations(t)

		})
	}
}
