package mongodbrepository

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"

	mongodbcustomtype "github.com/thnthien/impa/custom-type/mongodb"
)

func getValueAndType(obj any) (reflect.Value, reflect.Type) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)

	if t.Kind() == reflect.Pointer {
		v = reflect.Indirect(v)
		t = v.Type()
	} else if t.Kind() != reflect.Struct {
		panic("not handled this kind")
	}

	return v, t
}

func getID(obj any) any {
	v, t := getValueAndType(obj)

	for i := 0; i < v.NumField(); i++ {
		name := t.Field(i).Name
		if name == "ID" {
			return v.Field(i).Interface()
		}
	}
	return nil
}

func beforeCreate(obj any) {
	v, t := getValueAndType(obj)

	now := time.Now()

	for i := 0; i < v.NumField(); i++ {
		name := t.Field(i).Name
		if name == "CreatedAt" || name == "UpdatedAt" {
			v.Field(i).Set(reflect.ValueOf(now))
		}
	}
}

func beforeUpdate(obj any) {
	v, t := getValueAndType(obj)

	now := time.Now()

	for i := 0; i < v.NumField(); i++ {
		name := t.Field(i).Name
		if name == "UpdatedAt" {
			v.Field(i).Set(reflect.ValueOf(now))
		}
	}
}

func getIdFilter(id any) bson.M {
	var filter bson.M

	switch id.(type) {
	case mongodbcustomtype.UUID:
		i := id.(mongodbcustomtype.UUID)
		_, b, _ := i.MarshalBSONValue()
		filter = bson.M{"_id": bson.RawValue{
			Type:  bsontype.Binary,
			Value: b,
		}}
	default:
		filter = bson.M{"_id": id}
	}

	return filter
}
