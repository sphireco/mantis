package helper

import (
	"strconv"
	"time"
)

// Reverse a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// StrConvParseBoolHideError same as strconv.ParseBool except hides the error (returns false)
func StrConvParseBoolHideError(boolean string) bool {
	ret, err := strconv.ParseBool(boolean)
	if err != nil {
		return false
	}
	return ret
}

// StrConvAtoiWithDefault same as strconv.Atoi except only returns the value or a default value
func StrConvAtoiWithDefault(intAsString string, defaultValue int) int {
	intFromStr, intFromStrErr := strconv.Atoi(intAsString)
	if intFromStrErr != nil {
		intFromStr = defaultValue
	}
	return intFromStr
}

// StrConvAtoiWithDefaultTimeDuration same as strconv.Atoi except only returns the value or a default value
func StrConvAtoiWithDefaultTimeDuration(intAsString string, defaultValue int) time.Duration {
	intFromStr, intFromStrErr := strconv.Atoi(intAsString)
	if intFromStrErr != nil {
		intFromStr = defaultValue
	}
	return time.Duration(intFromStr)
}

// StringWithDefault
func StringWithDefault(givenValue string, defaultValue string) string {
	if givenValue == "" {
		return defaultValue
	}
	return givenValue
}

// IntWithDefault
func IntWithDefault(givenValue int, defaultValue int) int {
	if givenValue == 0 {
		return defaultValue
	}
	return givenValue
}
