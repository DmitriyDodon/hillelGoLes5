package generator

import (
	"math/rand"

	faker "github.com/bxcodec/faker/v4"
)

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type StringGenerator struct {
}

func (s StringGenerator) Generate() interface{} {
	return faker.Word()
}

func (s StringGenerator) GenerateSlice() []interface{} {
	sliceSize := 10
	sl := make([]interface{}, 0, sliceSize)
	for i := 0; i < sliceSize; i++ {
		sl = append(sl, s.Generate())
	}

	return sl
}

func (s StringGenerator) GenerateWithParam(param int) (interface{}, error) {
	resString := make([]rune, 0, param)
	for i := 0; i < param; i++ {
		resString = append(resString, alphabet[rand.Intn(len(alphabet))])
	}
	return string(resString), nil
}
