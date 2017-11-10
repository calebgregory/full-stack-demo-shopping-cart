package util

import (
	"encoding/json"
	"github.com/satchelhealth/errors"
	"io/ioutil"
	"net/http"
)

// ResponseError for a given request handler. Has StatusCode to return to http
// client and Message.
type ResponseError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// ResponseErrer defines the behavior of a Response that Errs
type ResponseErrer interface {
	ResponseError() *ResponseError
}

// ErringResponse implements the ResponseErrer interface. It is to be used as
// an embedded type in each Response-type struct.
type ErringResponse struct {
	Err *ResponseError `json:"error"`
}

// ResponseError implements the ResponseErrer interface by returning r.Err.
func (r *ErringResponse) ResponseError() *ResponseError {
	return r.Err
}

func AllowCORS(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}
		if r.Method == http.MethodOptions {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		f(w, r)
	})
}

func BindJSON(req *http.Request, schema interface{}) error {
	jsonBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return errors.Wrap(err, "bind json read all")
	}

	err = json.Unmarshal(jsonBody, schema)
	if err != nil {
		return errors.Wrap(err, "bind json unmarshall")
	}

	return nil
}

func WriteResponse(w http.ResponseWriter, res ResponseErrer) {
	if err := res.ResponseError(); err != nil {
		w.WriteHeader(err.StatusCode)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	b, marshalErr := json.Marshal(res)
	if marshalErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
