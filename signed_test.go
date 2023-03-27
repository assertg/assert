package assert

import "testing"

func TestSinged(t *testing.T) {

	t.Run("when assert positive vales expect ok", func(t *testing.T) {
		ThatIsPositive(t, int8(8))
		ThatIsPositive(t, int16(16))
		ThatIsPositive(t, int32(32))
		ThatIsPositive(t, int64(64))
		ThatIsPositive(t, float32(0.32))
		ThatIsPositive(t, float64(0.64))
	})

	t.Run("when assert negative vales expect ok", func(t *testing.T) {
		ThatIsNegative(t, int8(-8))
		ThatIsNegative(t, int16(-16))
		ThatIsNegative(t, int32(-32))
		ThatIsNegative(t, int64(-64))
		ThatIsNegative(t, float32(-0.32))
		ThatIsNegative(t, float64(-0.64))
	})

}
