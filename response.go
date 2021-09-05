package dragon

import (
	"encoding/json"
)


func (d *Dragon) ResponseJSON(statusCode int, v interface{}) error {
	d.ResponseWriter.Header().Set("Content-Type", "application/json")
	d.ResponseWriter.WriteHeader(statusCode)
	err := json.NewEncoder(d.ResponseWriter).Encode(v)
	if err != nil {
		return err
	}

	return nil
}
