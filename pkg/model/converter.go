package model

import (
	"fmt"
	"reflect"
	"time"
)

//convertTimestampToTime is converting the value of timestamp db-column to a Time instance
func convertTimestampToTime(value interface{}) (interface{}, error) {
	if reflect.TypeOf(value).Kind() == reflect.String {
		layout := "2006-01-02 15:04:05" //see https://golang.org/src/time/format.go
		return time.Parse(layout, value.(string))
	}
	if timeValue, ok := value.(time.Time); ok {
		return timeValue, nil
	}
	return nil, fmt.Errorf("failed to convert value '%s' (kind: %s) for field 'Created' to Time struct",
		value, reflect.TypeOf(value).Kind())
}
