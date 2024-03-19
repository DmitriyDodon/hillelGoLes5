package main

import (
	"flag"
	"fmt"
	"les5/generator"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

const defaultParam = 10

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	param := flag.Int("param", defaultParam, "Enter int param value")

	flag.Parse()

	generatorsArray := [4]generator.Generator{
		generator.BoolGenerator{},
		generator.IntGenerator{},
		generator.StringGenerator{},
		generator.Float64Generator{},
	}

	for _, generator := range generatorsArray {
		err := checkGenerator(generator, *param)
		if err != nil {
			logger.Error(err.Error())
		}
		fmt.Print("\n\n")
	}

}

func checkGenerator(generator generator.Generator, param int) error {
	logrus.Info("New generator checking start!")

	generateRes := generator.Generate()

	switch v := generateRes.(type) {
	case int:
		fmt.Printf("generateRes is int with value equal %d\n", v)
	case bool:
		fmt.Printf("generateRes is bool with value equal %t\n", v)
	case float64:
		fmt.Printf("generateRes is bool with value equal %.2f\n", v)
	case string:
		fmt.Printf("generateRes is bool with value equal %s\n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

	fmt.Printf("GenerateSlice result: %v\n", generator.GenerateSlice())

	res, err := generator.GenerateWithParam(param)

	if err != nil {
		return err
	}

	s, ok := res.(string)
	if ok {
		fmt.Printf("Res is string: %s\n", s)
	}

	i, ok := res.(int)
	if ok {
		fmt.Printf("Res is int: %d\n", i)
	}

	b, ok := res.(bool)
	if ok {
		fmt.Printf("Res is bool: %t\n", b)
	}

	f, ok := res.(float64)
	if ok {
		fmt.Printf("Res is float64: %.1f\n", f)
	}

	return nil
}
