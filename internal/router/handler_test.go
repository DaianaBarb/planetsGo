package router

import (
	"net/http"
	"net/http/httptest"
	"projeto-star-wars-api-go/internal/service/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestPlanetHandler_SavePlanet(t *testing.T) {
	//var rr *http.Request

	//planetJSON, _ := json.Marshal(planetToCreate)
	//r = planetJSON

	type fields struct {
		service *mocks.Planet
	}
	type args struct {
		r *http.Request
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(fs *mocks.Planet)
	}{
		{name: "sucesss",
			fields: fields{
				service: new(mocks.Planet),
			},
			args: args{
				r: &http.Request{},
			},
			wantErr: false,
			mock: func(fs *mocks.Planet) {
				fs.On("SavePlanet", mock.Anything, mock.Anything).
					Return(http.StatusCreated)
			},
		},
	}
	for _, tt := range tests {
		tt.mock(tt.fields.service)
		t.Run(tt.name, func(t *testing.T) {
			p := &PlanetHandler{
				service: tt.fields.service,
			}
			w := httptest.NewRecorder()
			p.SavePlanet(w, tt.args.r)

			tt.fields.service.AssertExpectations(t)

		})
	}
}
