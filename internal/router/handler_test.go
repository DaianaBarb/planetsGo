package router

import (
	"io"
	"net/http"
	"net/http/httptest"
	"projeto-star-wars-api-go/internal/model"
	"projeto-star-wars-api-go/internal/service/mocks"
	"strings"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

func TestPlanetHandler_SavePlanet(t *testing.T) {

	type fields struct {
		service *mocks.Planet
	}
	type args struct {
		body io.Reader
	}

	tests := []struct {
		name               string
		fields             fields
		args               args
		wantHttpStatusCode int
		mock               func(fs *mocks.Planet)
	}{
		{
			name: "sucesss",
			fields: fields{
				service: new(mocks.Planet),
			},
			args: args{
				body: strings.NewReader(`{"name":"mock", "climate":"mock", "terrain":"mock"}`),
			},
			wantHttpStatusCode: http.StatusCreated,
			mock: func(fs *mocks.Planet) {
				fs.On("Save", mock.Anything, mock.Anything).Return(mock.Anything, nil).Once()
			},
		},
		{
			name: "return 422 when don't send the body",
			fields: fields{
				service: new(mocks.Planet),
			},
			args: args{
				body: strings.NewReader(``),
			},
			wantHttpStatusCode: http.StatusUnprocessableEntity,
			mock: func(fs *mocks.Planet) {
				fs.On("Save", mock.Anything, mock.Anything).Maybe().Times(0)
			},
		},
	}
	for _, tt := range tests {
		tt.mock(tt.fields.service)
		t.Run(tt.name, func(t *testing.T) {
			p := &PlanetHandler{
				service: tt.fields.service,
			}

			request := httptest.NewRequest(http.MethodPost, "/planets", tt.args.body)
			recorder := httptest.NewRecorder()

			p.SavePlanet(recorder, request)

			assert.Equal(t, tt.wantHttpStatusCode, recorder.Code)

			tt.fields.service.AssertExpectations(t)

		})
	}
}

func TestPlanetHandler_DeleteById(t *testing.T) {
	id := primitive.NewObjectID()
	type fields struct {
		service *mocks.Planet
	}
	type args struct {
		id string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantHttpStatusCode int
		mock               func(fs *mocks.Planet)
	}{
		{
			name: "sucesss",
			fields: fields{
				service: new(mocks.Planet),
			},
			args: args{

				id: (id.Hex()),
			},
			wantHttpStatusCode: http.StatusOK,
			mock: func(fs *mocks.Planet) {
				fs.On("DeleteById", mock.Anything, mock.Anything).Return(nil).Once()
			}},
	}
	for _, tt := range tests {
		tt.mock(tt.fields.service)
		t.Run(tt.name, func(t *testing.T) {
			p := &PlanetHandler{
				service: tt.fields.service,
			}
			request := httptest.NewRequest(http.MethodDelete, "/planets/"+id.Hex(), nil)
			recorder := httptest.NewRecorder()

			p.DeleteById(recorder, request)

			assert.Equal(t, tt.wantHttpStatusCode, recorder.Code)

			tt.fields.service.AssertExpectations(t)
		})
	}
}

func TestPlanetHandler_GetAll(t *testing.T) {
	var planets []model.PlanetOut
	type fields struct {
		service *mocks.Planet
	}

	tests := []struct {
		name               string
		fields             fields
		wantHttpStatusCode int
		mock               func(fs *mocks.Planet)
	}{
		{
			name: "sucesss",
			fields: fields{
				service: new(mocks.Planet),
			},
			wantHttpStatusCode: http.StatusOK,
			mock: func(fs *mocks.Planet) {
				fs.On("FindAll", mock.Anything, mock.Anything).Return(planets, nil).Once()
			}},
	}
	for _, tt := range tests {
		tt.mock(tt.fields.service)
		t.Run(tt.name, func(t *testing.T) {
			p := &PlanetHandler{
				service: tt.fields.service,
			}
			request := httptest.NewRequest(http.MethodGet, "/planets", nil)
			recorder := httptest.NewRecorder()

			p.FindAll(recorder, request)

			assert.Equal(t, tt.wantHttpStatusCode, recorder.Code)

			tt.fields.service.AssertExpectations(t)
		})
	}
}

func TestPlanetHandler_Update(t *testing.T) {
	id := primitive.NewObjectID()
	type fields struct {
		service *mocks.Planet
	}
	type args struct {
		id   string
		body io.Reader
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantHttpStatusCode int
		mock               func(fs *mocks.Planet)
	}{
		{name: "sucesss",
			fields: fields{
				service: new(mocks.Planet),
			},
			args: args{
				id:   (id.Hex()),
				body: strings.NewReader(`{"name":"mock", "climate":"mock", "terrain":"mock"}`),
			},
			wantHttpStatusCode: http.StatusOK,
			mock: func(fs *mocks.Planet) {
				fs.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			}},
	}
	for _, tt := range tests {
		tt.mock(tt.fields.service)
		t.Run(tt.name, func(t *testing.T) {
			p := &PlanetHandler{
				service: tt.fields.service,
			}
			request := httptest.NewRequest(http.MethodPut, "/planets/"+id.Hex(), tt.args.body)
			recorder := httptest.NewRecorder()

			p.Update(recorder, request)

			assert.Equal(t, tt.wantHttpStatusCode, recorder.Code)

			tt.fields.service.AssertExpectations(t)
		})
	}
}

func TestPlanetHandler_FindById(t *testing.T) {
	id := primitive.NewObjectID()
	planet := &model.PlanetOut{ID: id, Name: "", Terrain: "", Climate: ""}

	type fields struct {
		service *mocks.Planet
	}
	type args struct {
		id string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantHttpStatusCode int
		mock               func(fs *mocks.Planet)
	}{
		{
			name: "sucesss",
			fields: fields{
				service: new(mocks.Planet),
			},
			args: args{
				id: (id.Hex()),
			},
			wantHttpStatusCode: http.StatusOK,
			mock: func(fs *mocks.Planet) {
				fs.On("FindById", mock.Anything, mock.Anything).Return(planet, nil).Once()
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.service)
			p := &PlanetHandler{
				service: tt.fields.service,
			}
			request := httptest.NewRequest(http.MethodGet, "/planets/"+id.Hex(), nil)
			recorder := httptest.NewRecorder()

			p.FindById(recorder, request)

			assert.Equal(t, tt.wantHttpStatusCode, recorder.Code)

			tt.fields.service.AssertExpectations(t)
		})
	}
}
