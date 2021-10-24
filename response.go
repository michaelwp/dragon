package dragon

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
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

func (d *Dragon) ResponseHTML(statusCode int, v interface{}) error {
	d.ResponseHeader.Set("Content-Type", "text/html; charset=utf-8")
	d.ResponseStatus(statusCode)
	_, err := fmt.Fprintf(d.ResponseWriter, v.(string))
	if err != nil {
		return err
	}

	return nil
}

func (d *Dragon) ResponseStatus(statusCode int) {
	d.ResponseWriter.WriteHeader(statusCode)
}
