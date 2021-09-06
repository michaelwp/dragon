package dragon

import "context"

func (d *Dragon) SetContext(key interface{}, val interface{}) {
	d.Context = context.WithValue(d.Context, key, val)
}

func (d *Dragon) GetContext(key interface{}) interface{}{
	return d.Context.Value(key)
}
