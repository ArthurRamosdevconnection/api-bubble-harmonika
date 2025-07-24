package crud

import (
	"reflect"
)

func GetJSONFieldNames[T any]() []string {
	var names []string
	t := reflect.TypeOf((*T)(nil)).Elem() // pega o tipo da struct

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Ignora campos embutidos ou privados
		if field.PkgPath != "" {
			continue
		}

		tag := field.Tag.Get("json")
		if tag == "-" {
			continue
		}

		// Se a tag tiver vírgula (ex: "name,omitempty"), pega só a parte antes da vírgula
		name := tag
		if name == "" {
			name = field.Name // fallback: nome do campo
		} else if comma := indexComma(name); comma >= 0 {
			name = name[:comma]
		}
		if name == "-" {
			continue
		}
		if name == "BaseModel" {
			continue
		}
		names = append(names, name)
	}
	return append(names, "id")
}

func indexComma(s string) int {
	for i, r := range s {
		if r == ',' {
			return i
		}
	}
	return -1
}
