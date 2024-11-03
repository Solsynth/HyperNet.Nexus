package sec

import (
	"fmt"
	"reflect"
	"strconv"
)

func (v UserInfo) HasPermNode(requiredKey string, requiredValue any) bool {
	if heldValue, ok := v.PermNodes[requiredKey]; ok {
		return comparePermNode(heldValue, requiredValue)
	}
	return false
}

func (v UserInfo) HasPermNodeWithDefault(requiredKey string, requiredValue any, defaultValue any) bool {
	if heldValue, ok := v.PermNodes[requiredKey]; ok {
		return comparePermNode(heldValue, requiredValue)
	}
	return comparePermNode(defaultValue, requiredValue)
}

func comparePermNode(held any, required any) bool {
	isNumeric := func(val reflect.Value) bool {
		kind := val.Kind()
		return kind >= reflect.Int && kind <= reflect.Uint64 || kind >= reflect.Float32 && kind <= reflect.Float64
	}

	toFloat64 := func(val reflect.Value) float64 {
		switch val.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return float64(val.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return float64(val.Uint())
		case reflect.Float32, reflect.Float64:
			return val.Float()
		default:
			panic(fmt.Sprintf("non-numeric value of kind %s", val.Kind()))
		}
	}

	heldValue := reflect.ValueOf(held)
	requiredValue := reflect.ValueOf(required)

	if isNumeric(requiredValue) && heldValue.Kind() == reflect.String {
		numericValue, _ := strconv.ParseFloat(heldValue.String(), 64)
		return numericValue >= toFloat64(requiredValue)
	}

	switch heldValue.Kind() {
	case reflect.String:
		if heldValue.String() == requiredValue.String() {
			return true
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < heldValue.Len(); i++ {
			if reflect.DeepEqual(heldValue.Index(i).Interface(), required) {
				return true
			}
		}
	default:
		if isNumeric(heldValue) && isNumeric(requiredValue) {
			return toFloat64(heldValue) >= toFloat64(requiredValue)
		}

		if reflect.DeepEqual(held, required) {
			return true
		}
	}

	return false
}
