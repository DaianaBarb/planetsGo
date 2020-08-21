package router

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"projeto-star-wars-api-go/internal/service/mocks"
	"strings"
	"testing"

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
