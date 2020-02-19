package controller

// Context ...
type Context interface {
	JSON(code int, i interface{}) error
}
