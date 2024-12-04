package evaluators

type Evaluator interface {
	Elegible(operator string) bool
	Evaluate(param any, value any) bool
}

type Evaluators struct {
	evaluators []Evaluator
}

func NewEvaluatorsManager(evaluators []Evaluator) *Evaluators {
	return &Evaluators{
		evaluators: evaluators,
	}
}

func (e *Evaluators) Pick(operator string) Evaluator {
	for _, evaluator := range e.evaluators {
		if evaluator.Elegible(operator) {
			return evaluator
		}
	}
	return nil
}
