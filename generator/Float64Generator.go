package generator

import (
	"errors"
	"math"
	"math/rand"
)

type Float64Generator struct {
}

func (f Float64Generator) Generate() interface{} {
	return rand.Float64()
}

func (f Float64Generator) GenerateSlice() []interface{} {
	sliceSize := 10
	sl := make([]interface{}, 0, sliceSize)
	for i := 0; i < sliceSize; i++ {
		sl = append(sl, f.Generate())
	}
	return sl
}

func (f Float64Generator) GenerateWithParam(param int) (interface{}, error) {
	param++
	r := roundFloat(0+rand.Float64()*float64((param-0)), 1)
	if r < 10 {
		return r, errors.New("float64 less than 10")
	}
	return r, nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
