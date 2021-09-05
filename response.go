package dragon

func (d *Dragon) ResponseStatus(statusCode int) {
	d.ResponseWriter.WriteHeader(statusCode)
}
