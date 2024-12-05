package usecase

import (
	"featureflag/internal/pkg/contracts"
	"featureflag/pkg/evaluators"
	"fmt"
)

type ValidateFlag struct {
	repository contracts.Repository
}

func (action *ValidateFlag) Execute(flagKey string) (bool, error) {
	flag, err := action.repository.Get(flagKey)
	if err != nil {
		return false, err
	}

	if !flag.Enabled {
		return false, nil
	}

	return evaluateConditions(flag.Conditions, map[string]any{}), nil
}

func evaluateConditions(conditions []map[string]any, params map[string]any) bool {
	var result bool
	for _, condition := range conditions {
		// Recupera os dados da condição
		paramKey := condition["parameter"].(string)
		operator := condition["operator"].(string)
		order := int(condition["order"].(int)) // Ordem é um número
		logical := "and"                       // Valor padrão para lógica
		if val, ok := condition["logical"].(string); ok {
			logical = val
		}

		// Recupera o valor do parâmetro
		paramValue, paramExists := params[paramKey]
		if !paramExists {
			paramValue = "" // Caso o parâmetro não esteja presente
		}

		// Avalia o operador
		evaluatorManager := evaluators.NewEvaluatorsManager([]contracts.Evaluator{
			&evaluators.InEvaluator{},
			&evaluators.BeginsWithEvaluator{},
		})
		evaluator := evaluatorManager.Pick(operator)
		if evaluator == nil {
			fmt.Printf("Operador nao suportado: %s\n", operator)
			continue
		}
		conditionResult := evaluator.Evaluate(paramValue, condition["value"])

		// Aplica a lógica booleana (and/or)
		if order == 0 {
			result = conditionResult
		} else if logical == "and" {
			result = result && conditionResult
		} else if logical == "or" {
			result = result || conditionResult
		}
	}

	return result
}
