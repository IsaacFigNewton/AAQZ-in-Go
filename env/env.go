package env

type Binding struct {
	Value interface{}
}

// let the top environment be a hashmap that maps strings to functions
var TopEnv = make(map[string]Binding)

func init() {
	TopEnv["+"] = Binding{Value: func(a, b float64) float64 {
		return a + b
	}}
	TopEnv["-"] = Binding{Value: func(a, b float64) float64 {
		return a - b
	}}
	TopEnv["*"] = Binding{Value: func(a, b float64) float64 {
		return a * b
	}}
	TopEnv["/"] = Binding{Value: func(a, b float64) float64 {
		return a / b
	}}
	TopEnv["<="] = Binding{Value: func(a, b float64) bool {
		return a <= b
	}}
	TopEnv["equal?"] = Binding{Value: func(a, b string) bool {
		return a == b
	}}
	TopEnv["true"] = Binding{Value: true}
	TopEnv["false"] = Binding{Value: false}
}
