package dragon

import (
	"encoding/json"
	"encoding/xml"
)

func (d *Dragon) ResponseJSON(statusCode int, v interface{}) error {
	d.ResponseHeader.Set("Content-Type", "application/json")
	d.ResponseStatus(statusCode)
	err := json.NewEncoder(d.ResponseWriter).Encode(v)
	if err != nil {
		return err
	}

	return nil
}

func (d *Dragon) ResponseXML(statusCode int, v interface{}) error {
	d.ResponseHeader.Set("Content-Type", "application/xml")
	d.ResponseStatus(statusCode)
	err := xml.NewEncoder(d.ResponseWriter).Encode(v)
	if err != nil {
		return err
	}

	return nil
}

func (d *Dragon) ResponseStatus(statusCode int) {
	d.ResponseWriter.WriteHeader(statusCode)
}
