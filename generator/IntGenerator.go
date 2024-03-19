package generator

import (
	"errors"
	"math/rand"
	"time"
)

type IntGenerator struct {
}

func (i IntGenerator) Generate() interface{} {
	return generateRandomNumber(int(time.Now().Unix()))
}

func (i IntGenerator) GenerateSlice() []interface{} {
	sliceSize := 10
	sl := make([]interface{}, 0, sliceSize)
	for ind := 0; ind < sliceSize; ind++ {
		sl = append(sl, i.Generate())
	}

	return sl
}

func (i IntGenerator) GenerateWithParam(param int) (interface{}, error) {
	randomNumber := generateRandomNumber(param)
	if randomNumber > 50 {
		return randomNumber, errors.New("number more then 50")
	}
	return randomNumber, nil
}

func generateRandomNumber(maxNumber int) int {
	randNumber := rand.Intn(maxNumber + 1)
	return randNumber
}
