package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/svbh/middleware/types"
)

func MakeHTTPHandler(handlerFunc types.HTTPAPIFunc) http.HandlerFunc {
	// Create a closure and return an anonymous function
	return func(w http.ResponseWriter, r *http.Request) {
		// Call the handler
		resp, err := handlerFunc(r)
		if err != nil {
			// Log error
			log.Errorf("Handler for %s %s returned error: %s", r.Method, r.URL, err)

			// Send HTTP response
			if resp != nil {
				err = WriteJSON(w, http.StatusInternalServerError, resp)
				if err != nil {
					log.Errorf("Error generating json. Err: %v", err)
				}
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			// Send HTTP response as Json
			err = WriteJSON(w, http.StatusOK, resp)
			if err != nil {
				log.Errorf("Error generating json. Err: %v", err)
			}
		}
	}
}

// WriteJSON writes the value v to the http response stream as json with standard
// json encoding.
func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	// Set content type as json
	w.Header().Set("Content-Type", "application/json")

	// write the HTTP status code
	w.WriteHeader(code)

	// Write the Json output
	return json.NewEncoder(w).Encode(v)
}

// ReadJSON is a utility function to parse JSON from http body
func ReadJSON(r *http.Request, v interface{}) error {
	// read http request body
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("Failed to read request: %v", err)
	}

	// parse request args and net conf
	if err := json.Unmarshal(content, v); err != nil {
		return fmt.Errorf("failed to parse http req: %v", err)
	}

	return nil
}
