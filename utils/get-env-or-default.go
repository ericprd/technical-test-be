package utils

import (
	"os"
	"reflect"
	"strconv"
)

func GetEnvOrDefault[T any](key string, defaultValue T) T {
	value, exist := os.LookupEnv(key)

	if !exist {
		return defaultValue
	}

	var valueInstance T
	kind := reflect.ValueOf(valueInstance).Kind()
	switch kind {
	case reflect.String:
		return any(value).(T)
	case reflect.Int:
		if val, err := strconv.Atoi(value); err == nil {
			return any(val).(T)
		}
		return any(defaultValue).(T)
	case reflect.Bool:
		if val, err := strconv.ParseBool(value); err == nil {
			return any(val).(T)
		}
		return any(defaultValue).(T)
	default:
		return any(defaultValue).(T)
	}
}
