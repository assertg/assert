package assert

import "testing"

func TestThatIsPositive(t *testing.T) {

	t.Run("when assert positive vales expect ok", func(t *testing.T) {
		ThatIsPositive(t, int8(8))
		ThatIsPositive(t, int16(16))
		ThatIsPositive(t, int32(32))
		ThatIsPositive(t, int64(64))
		ThatIsPositive(t, float32(0.32))
		ThatIsPositive(t, float64(0.64))
	})

}
