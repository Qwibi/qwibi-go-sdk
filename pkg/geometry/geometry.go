package geometry

import (
	"database/sql/driver"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/mitchellh/mapstructure"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// QGeometry ...
type geometry interface {
	Valid() error
	Pb() *proto.QPBxGeometry
	GetType() string
	//Value() (driver.Value, error)
	//UnmarshalJSON(data []byte) error
}

type QGeometry struct {
	geometry
}

func NewGeometry(g geometry) *QGeometry {
	return &QGeometry{g}
}

// NewGeometryPb ...
func NewGeometryPb(in *proto.QPBxGeometry) (*QGeometry, error) {
	if in == nil {
		return nil, qlog.Error("Wrong geometry parameter type nil")
	}

	switch v := in.Type.(type) {
	case *proto.QPBxGeometry_Point:
		point, err := NewPointPb(v.Point)
		if err != nil {
			return nil, qlog.Error(err)
		}
		return &QGeometry{point}, nil
	default:
		return nil, qlog.Errorf("Unknown geometry type: %s", v)
	}
}

func NewGeometryStruct(v map[string]interface{}) (*QGeometry, error) {
	if v == nil {
		return nil, qlog.Error("Bad parameter type nil")
	}

	switch t := v["type"]; t {
	case QPointGeometryType:
		var o QPoint
		err := mapstructure.Decode(v, &o)
		if err != nil {
			return nil, qlog.Error(err)
		}
		return &QGeometry{&o}, nil
	default:
		return nil, qlog.Errorf("Unknown gemetry type: %s", t)
	}

}

func (c *QGeometry) Value() (driver.Value, error) {
	qlog.Infof("=> Value() %+v", c)
	//b, err := protobuf.Marshal(c.Pb())
	//if err != nil {
	//	return nil, qlog.Error(err)
	//}

	////todo()
	//pb := proto.QPBxGeometry{}
	//
	//err = protobuf.Unmarshal(b, &pb)
	//if err != nil {
	//	return nil, qlog.Error(err)
	//}
	//
	//qlog.Debugf("[QGeometry] Value: %+v => %+v  => %+v", c.geometry, b, pb)
	return driver.Value(c.String()), nil
}

func (c *QGeometry) Scan(src any) error {
	qlog.Infof("==> Scan() %+v", c)
	if src == nil {
		return nil
	}

	pb := &proto.QPBxGeometry{}

	if b, ok := src.([]byte); ok {
		if err := protobuf.Unmarshal(b, pb); err != nil {
			return qlog.Error(err)
		}
	}

	c, err := NewGeometryPb(pb)
	//qlog.Infof("????? %+v", c)
	if err != nil {
		qlog.Error(err)
	}

	return nil
}

func (c *QGeometry) String() string {
	return fmt.Sprint(c.geometry)
}

func (c *QGeometry) FormatParam(placeholder string, info *sql.StmtInfo) string {
	if info.Dialect == dialect.Postgres {
		return "ST_GeomFromGeoJSON(" + placeholder + ")"
	}

	return placeholder
}

func (c *QGeometry) SchemaType() map[string]string {
	return map[string]string{
		dialect.Postgres: "geometry",
	}
}
