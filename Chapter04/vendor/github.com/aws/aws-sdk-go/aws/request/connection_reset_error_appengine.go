// +build appengine

package request

import (
	"strings"
)

func isErrConnectionReset(err error) bool {
	return strings.Contains(err.Error(), "connection reset")
}
