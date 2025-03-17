package helper

import (
	"net/http"

	"github.com/alielmi98/go-ecommerce-api/pkg/service_errors"
)

var StatusCodeMapping = map[string]int{
	// User
	service_errors.EmailExists:               409,
	service_errors.UsernameExists:            409,
	service_errors.RecordNotFound:            404,
	service_errors.PermissionDenied:          403,
	service_errors.UsernameOrPasswordInvalid: 401,
	// Token
	service_errors.InvalidRefreshToken: 401,
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return value
}
