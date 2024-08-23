package httputil

import "time"

type HttpError struct {
	Error     string `json:"error"`
	TimeStamp string `json:"timestamp"`
}

func NewHttpError(message string) *HttpError {
	return &HttpError{
		Error:     message,
		TimeStamp: time.Now().Format(time.RFC3339),
	}
}
