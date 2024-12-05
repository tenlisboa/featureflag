package contracts

type Evaluator interface {
	Elegible(operator string) bool
	Evaluate(param any, value any) bool
}
