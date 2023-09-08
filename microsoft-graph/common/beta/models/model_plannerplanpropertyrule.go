package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanPropertyRule struct {
	Buckets              *[]string                        `json:"buckets,omitempty"`
	CategoryDescriptions *PlannerFieldRules               `json:"categoryDescriptions,omitempty"`
	ODataType            *string                          `json:"@odata.type,omitempty"`
	RuleKind             *PlannerPlanPropertyRuleRuleKind `json:"ruleKind,omitempty"`
	Tasks                *[]string                        `json:"tasks,omitempty"`
	Title                *PlannerFieldRules               `json:"title,omitempty"`
}
