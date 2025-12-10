package repository

import "errors"

var (
	ErrNotObjectID = errors.New("inserted id is not an ObjectID")
)
