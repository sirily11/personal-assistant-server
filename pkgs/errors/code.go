package errors

// ErrorCode is a custom error code type
// 1000 - 1999: 400 Bad Request
// 2000 - 2999: 401 Unauthorized
// 3000 - 3999: 403 Forbidden
// 4000 - 4999: 404 Not Found
type ErrorCode int

const (
	ErrTheGivenResourceWasNotFound ErrorCode = 4000
)
