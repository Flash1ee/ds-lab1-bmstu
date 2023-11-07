package domain

import "errors"

type Error struct {
	Message error `json:"message"`
}

var (
	RecordNotFound = errors.New("record is not found")
	UnknownError   = errors.New("unknown error")
	InvalidRequest = errors.New("invalid request data")
	InvalidPerson  = errors.New("invalid person data")
)
