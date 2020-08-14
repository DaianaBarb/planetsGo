package planet

//import (
//	"reflect"
//	"testing"
//
//	"go.mongodb.org/mongo-driver/bson/primitive"
//)
//
//func TestPlanetDocument_ToPlanetOut(t *testing.T) {
//
//	id := primitive.NewObjectID()
//
//	type fields struct {
//		PlanetDocument PlanetDocument
//	}
//
//	tests := []struct {
//		name   string
//		fields fields
//		want   *PlanetOut
//	}{
//		{
//			name: "parse planet document to planet out",
//			fields: fields{PlanetDocument: PlanetDocument{
//				ID:      id,
//				Name:    "Terra",
//				Climate: "frio",
//				Terrain: "sei laoq",
//			}},
//			want: &PlanetOut{
//				ID:                      id,
//				Name:                    "Terra",
//				Climate:                 "frio",
//				Terrain:                 "sei laoq",
//				NumberOfFilmAppearances: 0,
//			}},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := tt.fields.PlanetDocument
//
//			if got := p.ToPlanetOut(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ToPlanetOut() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestPlanetIn_ToDocument(t *testing.T) {
//
//	type fields struct {
//		PlanetIn PlanetIn
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   *PlanetDocument
//	}{
//		{
//			name: "parse planetIn  to planetDocument",
//			fields: fields{PlanetIn: PlanetIn{
//				Name:    "Terra",
//				Climate: "frio",
//				Terrain: "arid",
//			}},
//			want: &PlanetDocument{
//				ID:      primitive.ObjectID{},
//				Name:    "Terra",
//				Climate: "frio",
//				Terrain: "arid",
//			}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := tt.fields.PlanetIn
//			if got := p.ToDocument(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ToDocument() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
