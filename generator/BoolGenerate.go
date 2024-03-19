package generator

import (
	"errors"
	"math/rand"
)

type BoolGenerator struct {
}

func (b BoolGenerator) Generate() interface{} {
	return rand.Intn(2) == 1
}

func (b BoolGenerator) GenerateSlice() []interface{} {
	sliceSize := 10
	sl := make([]interface{}, 0, sliceSize)
	for i := 0; i < sliceSize; i++ {
		sl = append(sl, b.Generate())
	}
	return sl
}

func (b BoolGenerator) GenerateWithParam(param int) (interface{}, error) {
	res := param == 1
	if !res {
		return res, errors.New("false value")
	}
	return res, nil
}
