package geo

import (
	"github.com/qwibi/qwibi-go-sdk/geometry"
	"github.com/qwibi/qwibi-go-sdk/utils"
	"google.golang.org/protobuf/types/known/structpb"
)

// NewGeoPoint ...
func NewGeoLayer() *QGeoObject {
	return &QGeoObject{
		Gid:        utils.NewID(),
		Geometry:   geometry.NewDefaultLayer(),
		Properties: &structpb.Struct{},
	}
}
