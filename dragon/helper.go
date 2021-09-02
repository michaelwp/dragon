package dragon

import "reflect"

func isMatching(expect string, actual string) bool {
	return reflect.DeepEqual(expect, actual)
}
