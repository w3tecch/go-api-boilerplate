package testhelpers

import (
	"encoding/json"
	"net/http/httptest"
)

// JSON reads the json response and decodes it
func JSON(w httptest.ResponseRecorder, res interface{}) {
	str := w.Body.String()
	json.Unmarshal([]byte(str), &res)
}
