package commonreflect

import (
	"reflect"
	"strings"
)

// GetAllUnignoredNames will return all tags which have not been defined as upsert_ignore=true
func GetAllUnignoredNames(tag string, obj any) []string {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		return nil
	}

	names := make([]string, 0, t.NumField())

FieldsLoop:
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tags := strings.Split(field.Tag.Get(tag), ";")
		if len(tags) == 0 {
			continue
		}
		name := tags[0]
		for j := 1; j < len(tags); j++ {
			if tags[j] == "upsert_ignore=true" {
				continue FieldsLoop
			}
		}
		names = append(names, name)
	}
	return names
}
