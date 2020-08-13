package planet

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"testing"
)

func TestPlanetDocument_ToPlanetOut(t *testing.T) {

	id := primitive.NewObjectID()

	type fields struct {
		PlanetDocument PlanetDocument
	}

	tests := []struct {
		name   string
		fields fields
		want   *PlanetOut
	}{
		{
			name: "parse planet document to planet out",
			fields: fields{PlanetDocument: PlanetDocument{
			ID:                      id,
			Name:                    "Terra",
			Climate:                 "frio",
			Terrain:                 "sei laoq",
			NumberOfFilmAppearances: 8,
		}},
		want: &PlanetOut{
			ID:                      id,
			Name:                    "Terra",
			Climate:                 "frio",
			Terrain:                 "sei laoq",
			NumberOfFilmAppearances: 8,
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.fields.PlanetDocument

			if got := p.ToPlanetOut(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToPlanetOut() = %v, want %v", got, tt.want)
			}
		})
	}
}