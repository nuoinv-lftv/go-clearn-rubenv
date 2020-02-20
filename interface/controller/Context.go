package controller

// Context ...
type Context interface {
	// JSON(code int, i interface{}) error
	Param(key string) string
	BindJSON(obj interface{}) error
	JSON(code int, i interface{})
	String(code int, format string, values ...interface{})
}
