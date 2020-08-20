package model

import (
	"projeto-star-wars-api-go/internal/provider/mongo/document"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestPlanetIn_ToPlanet(t *testing.T) {
	//
	type fields struct {
		PlanetIn PlanetIn
	}
	tests := []struct {
		name   string
		fields fields
		want   *Planet
	}{
		{
			name: " parse planetIn to planet",
			fields: fields{
				PlanetIn: PlanetIn{
					Name:    "Terra",
					Climate: "frio",
					Terrain: "arid",
				}},
			want: &Planet{

				Name:    "Terra",
				Climate: "frio",
				Terrain: "arid",
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.fields.PlanetIn
			if got := p.ToPlanet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToPlanet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlanet_ToPlanetOut(t *testing.T) {
	id := primitive.NewObjectID()
	type fields struct {
		Planet Planet
	}
	tests := []struct {
		name   string
		fields fields
		want   *PlanetOut
	}{{
		name: " parse planet to planetOut",
		fields: fields{
			Planet: Planet{
				ID:      id,
				Name:    "Terra",
				Climate: "frio",
				Terrain: "arid",
			}},
		want: &PlanetOut{
			ID:      id,
			Name:    "Terra",
			Climate: "frio",
			Terrain: "arid",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.fields.Planet
			if got := p.ToPlanetOut(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToPlanetOut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlanet_ToDocument(t *testing.T) {
	id := primitive.NewObjectID()
	type fields struct {
		Planet Planet
	}
	type args struct {
		id primitive.ObjectID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *document.Planet
	}{
		{name: " parse planet to document.plenet",
			fields: fields{
				Planet: Planet{

					Name:    "Terra",
					Climate: "frio",
					Terrain: "arid",
				}},
			args: args{id: id},
			want: &document.Planet{
				ID:      id,
				Name:    "Terra",
				Climate: "frio",
				Terrain: "arid",
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.fields.Planet
			if got := p.ToDocument(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}
