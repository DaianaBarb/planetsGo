package router

import (
	"encoding/json"
	"net/http"
	"projeto-star-wars-api-go/internal/service"
	"testing"
)

func TestPlanetHandler_SavePlanet(t *testing.T) {
	type fields struct {
		service service.Planet
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    json.Decoder
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PlanetHandler{
				service: tt.fields.service,
			}

		})
	}
}
