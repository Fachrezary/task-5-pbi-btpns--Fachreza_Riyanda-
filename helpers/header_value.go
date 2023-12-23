// helpers/header_value.go

package helpers

import (
	"net/http"
)

// GetAuthorizationHeader retrieves the value of the Authorization header from the HTTP request.
func GetAuthorizationHeader(r *http.Request) string {
	return r.Header.Get("Authorization")
}
