package v2

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"
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

	if IsArrOrSlc(object) {
		arr := SlcToAnys(object)
		idx, ok := AnyTryToType[int](field)
		if !ok {
			return nil, fmt.Errorf("field for array must be number")
		}
		if idx >= len(arr) {
			return nil, fmt.Errorf("field(idx) is out of range of object(array/slice)")
		}
		return arr[idx], nil
	}

	if len(field) > 0 && unicode.IsUpper(rune(field[0])) {
		ov := reflect.ValueOf(object)
		f := reflect.Indirect(ov).FieldByName(field)
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

// field must be exported, AND param value type can be converted to field value type.
func SetField(object any, field string, value any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	ov := reflect.ValueOf(object)

	if IsArrOrSlc(object) {
		idx, ok := AnyTryToType[int](field)
		if !ok {
			return fmt.Errorf("field for array must be number")
		}
		if idx >= len(SlcToAnys(object)) {
			return fmt.Errorf("field(idx) is out of range of object(array/slice)")
		}
		if e := ov.Index(idx); e.IsValid() {
			e.Set(reflect.ValueOf(value))
			return nil
		}
		goto ERR
	}

	if len(field) > 0 && unicode.IsUpper(rune(field[0])) {

		if ov.Kind() != reflect.Ptr {
			return fmt.Errorf("object need to pass its address")
		}

		if f := reflect.Indirect(ov).FieldByName(field); f.IsValid() {

			fKind, fType := f.Kind().String(), f.Type().String()

			switch fKind {
			case "struct", "slice", "array", "map", "ptr", "interface":
				if fType != "time.Time" {
					f.Set(reflect.ValueOf(value))
					return nil
				}
			}

			///////////////////////////////////////////////////

			switch fType {

			case "string":
				if val, ok := AnyTryToType[string](value); ok {
					f.Set(reflect.ValueOf(val))
					return nil
				}
				goto ERR

			case "int":
				if val, ok := AnyTryToType[int](value); ok {
					f.Set(reflect.ValueOf(val))
					return nil
				}
				goto ERR

			case "int8":
				if val, ok := AnyTryToType[int8](value); ok {
					f.Set(reflect.ValueOf(val))
					return nil
				}
				goto ERR

			case "int16":
				if val, ok := AnyTryToType[int16](value); ok {
					f.Set(reflect.ValueOf(val))
					return nil
				}
				goto ERR

			case "int32":
				if val, ok := AnyTryToType[int32](value); ok {
					f.Set(reflect.ValueOf(val))
					return nil
				}
				goto ERR

			case "int64":
				if val, ok := AnyTryToType[int64](value); ok {
					f.Set(reflect.ValueOf(val))
					return nil
				}
				goto ERR

			case "uint":
				if val, ok := AnyTryToType[uint](value); ok {
					f.Set(reflect.ValueOf(val))
					return nil
				}
				goto ERR

			case "uint8":
				if val, ok := AnyTryToType[uint8](value); ok {
					f.Set(reflect.ValueOf(val))
					return nil
				}
				goto ERR

			case "uint16":
				if val, ok := AnyTryToType[uint16](value); ok {
					f.Set(reflect.ValueOf(val))
					return nil
				}
				goto ERR

			case "uint32":
				if val, ok := AnyTryToType[uint32](value); ok {
					f.Set(reflect.ValueOf(val))
					return nil
				}
				goto ERR

			case "uint64":
				if val, ok := AnyTryToType[uint64](value); ok {
					f.Set(reflect.ValueOf(val))
					return nil
				}
				goto ERR

			case "float64":
				if val, ok := AnyTryToType[float64](value); ok {
					f.Set(reflect.ValueOf(val))
					return nil
				}
				goto ERR

			case "float32":
				if val, ok := AnyTryToType[float32](value); ok {
					f.Set(reflect.ValueOf(val))
					return nil
				}
				goto ERR

			case "bool":
				if val, ok := AnyTryToType[bool](value); ok {
					f.Set(reflect.ValueOf(val))
					return nil
				}
				goto ERR

			case "time.Time":
				if val, ok := AnyTryToType[time.Time](value); ok {
					f.Set(reflect.ValueOf(val))
					return nil
				}
				goto ERR

			default:
				log.Fatalf("need type [%v] for setting '%v' value @ [%v]", fType, field, value)
			}
		}
	}

ERR:
	return fmt.Errorf("field '%v' failed to set value @ [%v]", field, value)
}

// set simple primitive
func FlatMapSetField[T any](fm map[string]any, object any, field string) error {
	if v, ok := FlatMapValTryToType[T](fm, field); ok {
		if err := SetField(object, field, v); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("[%v] cannot be found or set as [%T]", field, *new(T))
	}
	return nil
}

// set primitive array or slice
func FlatMapSetFieldAsSlc[T any](fm map[string]any, object any, field string) error {
	if v, ok := FlatMapValsTryToTypes[T](fm, field); ok {
		if err := SetField(object, field, v); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("[%v] cannot be found or set as [[]%T]", field, *new(T))
	}
	return nil
}

// set map of primitive key & value
func FlatMapSetFieldAsMap[T1 comparable, T2 any](fm map[string]any, object any, field string) error {
	m := make(map[T1]T2)
	for _, k := range FlatMapSubKeys(fm, field) {
		pk := field + "." + k
		if v, ok := FlatMapValTryToType[T2](fm, pk); ok {
			if key, ok := AnyTryToType[T1](k); ok {
				m[key] = v
			} else {
				return fmt.Errorf("map key cannot be converted to type %T", *new(T1))
			}
		} else {
			return fmt.Errorf("[%v] cannot be found or set as [map[%T]%T]", field, *new(T1), *new(T2))
		}
	}
	if len(m) > 0 {
		return SetField(object, field, m)
	}
	return nil
}

// set map of primitive key & slice of primitive value
func FlatMapSetFieldAsSlcValMap[T1 comparable, T2 any](fm map[string]any, object any, field string) error {
	m := make(map[T1][]T2)
	for _, k := range FlatMapSubKeys(fm, field) {
		pk := field + "." + k
		if v, ok := FlatMapValsTryToTypes[T2](fm, pk); ok {
			if key, ok := AnyTryToType[T1](k); ok {
				m[key] = v
			} else {
				return fmt.Errorf("map key cannot be converted to type %T", *new(T1))
			}
		} else {
			return fmt.Errorf("[%v] cannot be found or set as [map[%T][]%T]", field, *new(T1), *new(T2))
		}
	}
	if len(m) > 0 {
		return SetField(object, field, m)
	}
	return nil
}

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
	return MapFlatToNested(mFlat, nil), nil
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
