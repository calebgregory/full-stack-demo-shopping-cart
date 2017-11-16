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

// AllowCORS middleware, allows Cross-Origin Resource Sharing. If request
// Method = OPTIONS, sets headers then returns.
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
		// TODO: refactor content-type setting into own decorator
		w.Header().Set("Content-Type", "application/json")
		f(w, r)
	})
}

// BindJSON binds the request data incoming from HTTP Request Body to the
// schema provided as argument.
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

// WriteResponse writes the response struct provided as argument to the
// http.ResponseWriter. If res is ErringResponse, writes err.StatusCode and
// marshaled JSON of error message. Otherwise writes StatusOK and marshaled
// JSON of response struct.
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
