package assert

import "testing"

func TestError(t *testing.T) {
	NoError(t, nil)
}
