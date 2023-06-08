package env

import (
	"os"
	"strconv"
)

// GetBool gets the env var as a boolean
func GetBool(envVarName string, defaultValue bool) bool {
	stringValue, ok := os.LookupEnv(envVarName)
	if !ok {
		return defaultValue
	}

	switch stringValue {
	case "true", "TRUE", "True", "yes", "Yes", "YES","1", "t", "T":
		return true
	case "false", "FALSE", "False", "no", "No", "NO","0", "f", "F":
		return false
	}

	return defaultValue
}

// GetInt64 gets the env var as an int
func GetInt64(envVarName string, defaultValue int64) int64 {
	val, ok := os.LookupEnv(envVarName)
	if !ok {
		return defaultValue
	}

	integerValue, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return defaultValue
	}

	return integerValue
}

// GetFloat64 gets the env var a float
func GetFloat64(envVarName string, defaultValue float64) float64 {
	val, ok := os.LookupEnv(envVarName)
	if !ok {
		return defaultValue
	}

	floatVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return defaultValue
	}

	return floatVal
}

// GetString gets the env var as a string
func GetString(envVarName string, defaultValue string) string {
	stringValue, _ := os.LookupEnv(envVarName)
	if stringValue == "" {
		return defaultValue
	}

	return stringValue
}