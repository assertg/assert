package assert

import "testing"

func TestThatIsEqualTo(t *testing.T) {

	t.Run("when assert equal values expect ok", func(t *testing.T) {
		ThatIsEqualTo(t, true, true)
		ThatIsEqualTo(t, false, false)
		ThatIsEqualTo(t, 1, 1)
		ThatIsEqualTo(t, 1.2, 1.2)
		ThatIsEqualTo(t, "string", "string")

		ptr := &struct{ value int }{10}
		ptrSame := ptr
		ThatIsEqualTo(t, ptr, ptrSame)
	})

}
