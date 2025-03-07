package rules

// Rule is an interface that all Rule
// types implement. This give consistency
// for the documentation generation.
type Rule interface {
	Name() string
	Definition() string
	Identifier() string
}

// RuleCategory holds documentation and
// inventory for each of the rule types
type RuleCategory struct {
	Name       string
	Definition string
	Rules      []Rule
}

// Rules is a type to hold all existing
// rules. Exposed for documentation generator
type Rules struct {
	Categories []RuleCategory
}

var fieldRules RuleCategory = RuleCategory{
	Name:       "",
	Definition: "",
}

// GetRules returns a list of all rules for
// documentation generation
func GetRules() *Rules {
	categories := []RuleCategory{
		{
			Name:       "Provider Level Breakages",
			Definition: "Top level behavior such as provider configuration changes and authentication.",
			Rules:      nil,
		},
		{
			Name:       "Resource Inventory Level Breakages",
			Definition: "Resource/datasource naming conventions and entry differences.",
			Rules:      resourceInventoryRulesToRuleArray(ResourceInventoryRules),
		},
		{
			Name:       "Resource Level Breakages",
			Definition: "Individual resource breakages like field entry removals or behavior within a resource.",
			Rules:      resourceSchemaRulesToRuleArray(ResourceSchemaRules),
		},
		{
			Name:       "Field Level Breakages",
			Definition: "Field level conventions like attribute changes and naming conventions.",
			Rules:      fieldRulesToRuleArray(FieldRules),
		},
	}

	return &Rules{
		Categories: categories,
	}
}
