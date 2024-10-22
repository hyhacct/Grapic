package constant

import "net/http"

const (
	CODE200 = http.StatusOK
	CODE401 = http.StatusUnauthorized
	CODE403 = http.StatusForbidden
	CODE404 = http.StatusNotFound
	CODE500 = http.StatusInternalServerError
)
