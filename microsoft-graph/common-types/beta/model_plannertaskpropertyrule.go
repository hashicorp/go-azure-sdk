package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskPropertyRule struct {
	AppliedCategories      *PlannerFieldRules               `json:"appliedCategories,omitempty"`
	Assignments            *PlannerFieldRules               `json:"assignments,omitempty"`
	CheckLists             *PlannerFieldRules               `json:"checkLists,omitempty"`
	CompletionRequirements *[]string                        `json:"completionRequirements,omitempty"`
	Delete                 *[]string                        `json:"delete,omitempty"`
	DueDate                *[]string                        `json:"dueDate,omitempty"`
	Move                   *[]string                        `json:"move,omitempty"`
	Notes                  *[]string                        `json:"notes,omitempty"`
	ODataType              *string                          `json:"@odata.type,omitempty"`
	Order                  *[]string                        `json:"order,omitempty"`
	PercentComplete        *[]string                        `json:"percentComplete,omitempty"`
	PreviewType            *[]string                        `json:"previewType,omitempty"`
	Priority               *[]string                        `json:"priority,omitempty"`
	References             *PlannerFieldRules               `json:"references,omitempty"`
	RuleKind               *PlannerTaskPropertyRuleRuleKind `json:"ruleKind,omitempty"`
	StartDate              *[]string                        `json:"startDate,omitempty"`
	Title                  *[]string                        `json:"title,omitempty"`
}
