package internal

import "net/http"

func IsBadStatusCode(responseCode int) bool {
	return responseCode < http.StatusOK || responseCode >= http.StatusBadRequest
}
