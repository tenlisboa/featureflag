package evaluators

import "featureflag/internal/pkg/contracts"

type Evaluators struct {
	evaluators []contracts.Evaluator
}

func NewEvaluatorsManager(evaluators []contracts.Evaluator) *Evaluators {
	return &Evaluators{
		evaluators: evaluators,
	}
}

func (e *Evaluators) Pick(operator string) contracts.Evaluator {
	for _, evaluator := range e.evaluators {
		if evaluator.Elegible(operator) {
			return evaluator
		}
	}
	return nil
}
