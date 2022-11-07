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
	if obj == nil || (v.Kind() == reflect.Ptr && v.IsNil()) {
		return v, nil
	}
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
	v, _ := getValueAndType(obj)

	id := v.FieldByName("ID")
	if !id.IsZero() {
		return id.Interface()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() != reflect.Struct {
			continue
		}
		id = field.FieldByName("ID")
		if !id.IsZero() {
			return id.Interface()
		}
	}

	return nil
}

func structName(obj any) string {
	_, t := getValueAndType(obj)
	return t.Name()
}

func beforeCreate(obj any) {
	v, t := getValueAndType(obj)
	if t == nil {
		return
	}

	var now *time.Time

	for i := 0; i < v.NumField(); i++ {
		if !t.Field(i).IsExported() {
			continue
		}
		name := t.Field(i).Name
		switch v.Field(i).Kind() {
		case reflect.Struct:
			sName := structName(v.Field(i).Interface())
			if sName == "Time" && (name == "CreatedAt" || name == "UpdatedAt") {
				if now == nil {
					n := time.Now()
					now = &n
				}
				v.Field(i).Set(reflect.ValueOf(*now))
			} else {
				beforeCreate(v.Field(i).Addr().Interface())
			}
		case reflect.Pointer:
			beforeCreate(v.Field(i).Interface())
		}
	}
}

func beforeUpdate(obj any) {
	v, t := getValueAndType(obj)
	if t == nil {
		return
	}

	var now *time.Time

	for i := 0; i < v.NumField(); i++ {
		if !t.Field(i).IsExported() {
			continue
		}
		name := t.Field(i).Name
		switch v.Field(i).Kind() {
		case reflect.Struct:
			sName := structName(v.Field(i).Interface())
			if sName == "Time" && name == "UpdatedAt" {
				if now == nil {
					n := time.Now()
					now = &n
				}
				v.Field(i).Set(reflect.ValueOf(*now))
			} else {
				beforeUpdate(v.Field(i).Addr().Interface())
			}
		case reflect.Pointer:
			beforeUpdate(v.Field(i).Interface())
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
