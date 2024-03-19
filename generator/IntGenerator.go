package generator

import (
	"errors"
	"math/rand"
	"time"
)

type IntGenerator struct {
}

func (i IntGenerator) Generate() interface{} {
	return time.Now().Unix()
}

func (i IntGenerator) GenerateSlice() []interface{} {
	sliceSize := 10
	sl := make([]interface{}, 0, sliceSize)
	for i := 0; i < sliceSize; i++ {
		randomValue := generateRandomNumber(int(time.Now().Unix()))
		sl = append(sl, randomValue)
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
