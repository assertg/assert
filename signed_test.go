package assert

import "testing"

func TestSinged(t *testing.T) {

	t.Run("when assert positive vales expect ok", func(t *testing.T) {
		Positive(t, int8(8))
		Positive(t, int16(16))
		Positive(t, int32(32))
		Positive(t, int64(64))
		Positive(t, float32(0.32))
		Positive(t, float64(0.64))
	})

	t.Run("when assert negative vales expect ok", func(t *testing.T) {
		Negative(t, int8(-8))
		Negative(t, int16(-16))
		Negative(t, int32(-32))
		Negative(t, int64(-64))
		Negative(t, float32(-0.32))
		Negative(t, float64(-0.64))
	})

}
