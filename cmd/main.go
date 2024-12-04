package main

import (
	"featureflag/internal/pkg/evaluators"
	"fmt"
)

// Função para avaliar as condições
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
		evaluatorManager := evaluators.NewEvaluatorsManager([]evaluators.Evaluator{
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

// Exemplo de uso
func main() {
	conditions := []map[string]any{
		{
			"order":     0,
			"parameter": "name",
			"operator":  "begins-with",
			"value":     "Marc",
		},
		{
			"order":     2,
			"logical":   "and",
			"parameter": "status",
			"operator":  "in",
			"value":     []any{4, 5},
		},
	}

	params := map[string]any{
		"name":   "Marcus",
		"status": 2,
	}

	isActivated := evaluateConditions(conditions, params)
	fmt.Printf("Flag ativada: %t\n", isActivated)
}
