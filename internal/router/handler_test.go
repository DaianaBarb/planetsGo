package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"projeto-star-wars-api-go/internal/model"
	"projeto-star-wars-api-go/internal/service/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestPlanetHandler_SavePlanet(t *testing.T) {

	//var rr *http.Request planetToCreate := model.Planet{
	//Name:    "Tatooine",
	//	Climate: "arid",
	//	Terrain: "desert"}
	//planetJSON, _ := json.Marshal(planetToCreate)

	//b := bytes.NewBuffer([]byte(planetJSON))
	m := &model.PlanetIn{Name: "Test", Climate: "Test", Terrain: "Test"}

	b, _ := json.Marshal(m)
	r := bytes.NewBuffer(b)

	request, err := http.NewRequest("GET", "/http://localhost:8080/planets", r)
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		service *mocks.Planet
	}
	type args struct {
		r *http.Request
		w http.Response
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
				r: request,
				w: http.Response{},
			},
			wantErr: false,
			mock: func(fs *mocks.Planet) {
				fs.On("Save", mock.Anything, mock.Anything).
					Return(http.StatusCreated).Once()
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
