package util

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

// Iterates over all parameters and tries to set them
func DecorateType(serviceType interface{}, inputs []Input) (interface{}, error) {
	for _, input := range inputs {
		_, err := setParameter(serviceType, input)
		if err != nil {
			return nil, err
		}
	}

	return serviceType, nil
}

func getJsonFieldName(f reflect.StructField) string {
	jsonParameters := strings.Split(f.Tag.Get("json"), ",")
	if len(jsonParameters) == 0 || jsonParameters[0] == "" {
		return f.Name
	}
	return jsonParameters[0]
}

func getFieldNameByParameterName(v reflect.Type, parameterName string) string {
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Anonymous {
			// we're dealing with an anonymous/embedded struct
			fieldName := getFieldNameByParameterName(v.Field(i).Type, parameterName)
			if fieldName != "" {
				return fieldName
			}
		}
		if strings.EqualFold(getJsonFieldName(v.Field(i)), parameterName) {
			return v.Field(i).Name
		}
	}
	return ""
}

// Iterates through service Type to find and set the required parameters
// logs error and exits if parameter does not match any in Type specified field names
func setParameter(service interface{}, input Input) (interface{}, error) {
	reflectedServiceValue := reflect.ValueOf(service).Elem()
	reflectedServiceType := reflect.TypeOf(service).Elem()
	var parameterName string
	var err error
	for _, parameterName = range input.ParameterHierarchy {
		if reflectedServiceValue.Kind() == reflect.Pointer {
			if !reflectedServiceValue.Elem().IsValid() {
				// points to nothing, we need to create an instance
				reflectedServiceValue.Set(reflect.New(reflectedServiceType.Elem()))
			}

			// continue working with the struct the pointer points to (instead of working with the pointer itself)
			reflectedServiceType = reflectedServiceType.Elem()
			reflectedServiceValue = reflectedServiceValue.Elem()
		}

		fieldName := getFieldNameByParameterName(reflectedServiceType, parameterName)
		if fieldName == "" {
			err = fmt.Errorf("%s does not have field with name '%s'",
				reflectedServiceType,
				parameterName,
			)
			return nil, err
		}

		reflectedServiceValue = reflectedServiceValue.FieldByName(fieldName)
		x, _ := reflectedServiceType.FieldByName(fieldName)
		reflectedServiceType = x.Type
	}

	err = SetFields(reflectedServiceValue, input)
	if err != nil {
		err = fmt.Errorf(
			"%w cannot assign value %s to field %s with field Type %s",
			err,
			input.Value,
			strings.Join(input.ParameterHierarchy, HIERARCHY_DELIMITER),
			reflectedServiceValue.Kind(),
		)
		err = fmt.Errorf("%w %s contains field with name %s : %t",
			err,
			reflectedServiceValue.Type().Name(),
			parameterName,
			reflectedServiceValue.IsValid(),
		)
		return nil, err
	}
	info := fmt.Sprintf("setting field: '%s' value: '%s'", strings.Join(input.ParameterHierarchy, HIERARCHY_DELIMITER), input.Value)
	logrus.SetOutput(os.Stderr)
	logrus.Info(info)

	return service, nil
}

// Sets value of reflected field with type checking
func SetFields(field reflect.Value, input Input) error {
	if input.Unset {
		field.Set(reflect.Zero(field.Type()))
	} else if field.Kind() == reflect.String {
		field.SetString(input.Value)
	} else if isJson(input.Value) {
		field.Set(reflect.Zero(field.Type()))
		err := json.Unmarshal([]byte(input.Value), field.Addr().Interface())
		if err != nil {
			return fmt.Errorf("Json value could not be Unmarshalled: %s", err)
		}
	} else if field.Kind() >= reflect.Int && field.Kind() <= reflect.Int64 {
		intValue, err := strconv.ParseInt(input.Value, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(intValue)
	} else if field.Kind() >= reflect.Uint && field.Kind() <= reflect.Uint64 {
		intValue, err := strconv.ParseUint(input.Value, 10, 64)
		if err != nil {
			return err
		}
		field.SetUint(intValue)
	} else if field.Kind() == reflect.Float32 || field.Kind() == reflect.Float64 {
		floatValue, err := strconv.ParseFloat(input.Value, 64)
		if err != nil {
			return err
		}
		field.SetFloat(floatValue)
	} else if field.Kind() == reflect.Bool {
		boolValue, err := strconv.ParseBool(input.Value)
		if err != nil {
			return err
		}
		field.SetBool(boolValue)
	} else {
		return fmt.Errorf("setFields failed with field Type: %T and value: %s", reflect.TypeOf(field), input.Value)
	}
	return nil
}

// Checks if the argument is a json map and returns if it is valid json
func isJson(arg string) bool {
	if strings.HasPrefix(arg, "{") && strings.HasSuffix(arg, "}") {
		return json.Valid([]byte(arg))
	}
	return false
}
