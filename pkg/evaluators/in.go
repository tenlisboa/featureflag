package evaluators

type InEvaluator struct {
}

func (evaluator *InEvaluator) Elegible(operator string) bool {
	return operator == "in"
}

func (evaluator *InEvaluator) Evaluate(param any, value any) bool {
	values, ok := value.([]any)
	if !ok {
		return false
	}
	for _, v := range values {
		if param == v {
			return true
		}
	}
	return false
}
