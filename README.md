# MusGo_float64
Provides serialization of the Golang's `float64` type.

# How to use
Simply cast `float64` to `musgo_float64.Float64`:
```go
	var n float64 = 5.5
	// Marshal
	buf := make([]byte, musgo_float64.Float64(n).SizeMUS())
	musgo_float64.Float64(n).MarshalMUS(buf)
	// Unmarshal
	_, err := (*musgo_float64.Float64)(&n).UnmarshalMUS(buf)
	if err != nil {
		panic(err)
	}
```

# More info
You can find at [github.com/ymz-ncnk/musgo](https://github.com/ymz-ncnk/musgo).

