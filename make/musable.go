//go:build ignore

package main

import (
	"reflect"

	mgf "github.com/ymz-ncnk/musgo_float64"

	"github.com/ymz-ncnk/dotmusgo"
)

func main() {
	musGo, err := dotmusgo.New()
	if err != nil {
		panic(err)
	}
	unsafe := false
	err = musGo.Generate(reflect.TypeOf((*mgf.Float64)(nil)).Elem(), unsafe)
	if err != nil {
		panic(err)
	}
}
