package schema

import "time"

type ResSchema struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type NotFound struct {
	Message string
}

type BadRequest struct {
	Message string
}

type InternalServerError struct {
	Err     error
	Message string
}
