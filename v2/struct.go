package v2

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"unicode"
)

// get all fields
func Fields(object any) (fields []string) {
	if reflect.ValueOf(object).Kind() == reflect.Ptr {
		ptr := reflect.ValueOf(object).Elem()
		object = ptr.Interface()
	}
	typ := reflect.TypeOf(object)
	// fmt.Println("Type:", typ.Name(), "Kind:", typ.Kind())
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fields = append(fields, field.Name)
	}
	return
}

// get only exported field value
func FieldValue(object any, field string) (any, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Cannot get field '%s' in struct '%v'\n", field, reflect.TypeOf(object))
		}
	}()
	if reflect.ValueOf(object).Kind() == reflect.Ptr {
		ptr := reflect.ValueOf(object).Elem()
		object = ptr.Interface()
	}
	if len(field) > 0 && unicode.IsUpper(rune(field[0])) {
		r := reflect.ValueOf(object)
		f := reflect.Indirect(r).FieldByName(field)
		if !f.IsValid() || f.Kind() == 0 {
			return nil, fmt.Errorf("field '%s' is NOT in struct '%v'", field, reflect.TypeOf(object))
		}
		return f.Interface(), nil
	}
	return nil, fmt.Errorf("'%v' field '%s' is NOT exported", reflect.TypeOf(object), field)
}

func PathValue(object any, path string) (v any, err error) {
	for _, seg := range strings.Split(path, ".") {
		if v, err = FieldValue(object, seg); err != nil {
			return nil, err
		}
		object = v
	}
	return v, err
}

// field must be exported, AND param value type must be identical to field' value type.
func SetFieldValue(object any, field string, value any) (ok bool) {
	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()
	if reflect.ValueOf(object).Kind() != reflect.Ptr {
		return false
	}
	if len(field) > 0 && unicode.IsUpper(rune(field[0])) {
		v := reflect.ValueOf(object)
		if f := v.Elem().FieldByName(field); f.IsValid() {
			f.Set(reflect.ValueOf(value))
			return true
		}
	}
	return false
}

// TODO:
// func SetPathValue(object any, path string, value any) (ok bool) {
// }

// func PartialAsMap(object any, fields ...string) (any, error) {
// 	part := make(map[string]any)
// 	for _, field := range fields {
// 		v, err := FieldValue(object, field)
// 		if err != nil {
// 			return nil, err
// 		}
// 		part[field] = v
// 	}
// 	return part, nil
// }

func PartialAsMap(object any, paths ...string) (any, error) {
	mFlat := make(map[string]any)
	for _, path := range paths {
		v, err := PathValue(object, path)
		if err != nil {
			return nil, err
		}
		mFlat[path] = v
	}
	return MapFlatToNested(mFlat), nil
}

// get all tags
func Tags(object any, tag string, exclTags ...string) (tags []string) {
	if NotIn(tag, "json", "validate") {
		panic("tag must be [json, validate]")
	}
	if reflect.ValueOf(object).Kind() == reflect.Ptr {
		ptr := reflect.ValueOf(object).Elem()
		object = ptr.Interface()
	}
	typ := reflect.TypeOf(object)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get(tag)
		// fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
		tags = append(tags, strings.Split(tag, ",")...)
	}
	tags = Settify(tags...)
	FilterFast(&tags,
		func(i int, e string) bool {
			return len(e) > 0 && NotIn(e, exclTags...)
		},
	)
	return
}

// get all validator tags
func ValidatorTags(object any, exclTags ...string) (tags []string) {
	return Tags(object, "validate", exclTags...)
}

// get all json tags
func JsonTags(object any, exclTags ...string) (tags []string) {
	return Tags(object, "json", exclTags...)
}
