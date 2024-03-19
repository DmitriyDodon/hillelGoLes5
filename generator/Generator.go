package generator

type Generator interface {
	Generate() interface{}
	GenerateSlice() []interface{}
	GenerateWithParam(param int) (interface{}, error)
}
