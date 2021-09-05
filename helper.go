package dragon

import "reflect"

func isMatching(expect interface{}, actual interface{}) bool {
	return reflect.DeepEqual(expect, actual)
}


