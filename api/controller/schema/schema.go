package schema

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
