package entities

type Flag struct {
	Key         string
	Enabled     bool
	Conditions  []map[string]any
	Environment string
	Description string
}
