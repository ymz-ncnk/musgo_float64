package musgo_float64

import "testing"

func TestMusgoFloat64(t *testing.T) {
	var n float64 = 5.5
	buf := make([]byte, Float64(n).SizeMUS())
	Float64(n).MarshalMUS(buf)

	var an float64
	(*Float64)(&an).UnmarshalMUS(buf)

	if n != an {
		t.Fatal("something went wrong")
	}
}
