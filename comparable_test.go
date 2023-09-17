package assert

import "testing"

func TestEquals(t *testing.T) {

	t.Run("when assert equal values expect ok", func(t *testing.T) {
		Equals(t, true, true)
		Equals(t, false, false)
		Equals(t, 1, 1)
		Equals(t, 1.2, 1.2)
		Equals(t, "string", "string")

		ptr := &struct{ value int }{10}
		ptrSame := ptr
		Equals(t, ptr, ptrSame)
	})

}
