package serializers

import (
	"context"
	"errors"
	"Company import path goes here"

)

var (
	subjects            = global.Subjects
	UndefinedEventError = errors.New("event type is not defined")
)

func Serialize(_ context.Context, eventType string, data interface{}) (binary []byte, err error) {

	if _, ok := subjects[eventType]; !ok {
		err = UndefinedEventError
		return
	}

	binary, err = global.Registry.WithSchema(subjects[eventType]).Encode(data)
	if err != nil {
		log.Error("Error in serialization the event ", eventType, err)
		return
	}
	return
}
