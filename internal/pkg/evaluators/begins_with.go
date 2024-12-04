package evaluators

import "strings"

type BeginsWithEvaluator struct{}

func (evaluator *BeginsWithEvaluator) Elegible(operator string) bool {
	return operator == "begins-with"
}

func (evaluator *BeginsWithEvaluator) Evaluate(param any, value any) bool {
	return strings.HasPrefix(param.(string), value.(string))
}
