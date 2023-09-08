package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskPropertyRuleRuleKind string

const (
	PlannerTaskPropertyRuleRuleKindbucketRule PlannerTaskPropertyRuleRuleKind = "BucketRule"
	PlannerTaskPropertyRuleRuleKindplanRule   PlannerTaskPropertyRuleRuleKind = "PlanRule"
	PlannerTaskPropertyRuleRuleKindtaskRule   PlannerTaskPropertyRuleRuleKind = "TaskRule"
)

func PossibleValuesForPlannerTaskPropertyRuleRuleKind() []string {
	return []string{
		string(PlannerTaskPropertyRuleRuleKindbucketRule),
		string(PlannerTaskPropertyRuleRuleKindplanRule),
		string(PlannerTaskPropertyRuleRuleKindtaskRule),
	}
}

func parsePlannerTaskPropertyRuleRuleKind(input string) (*PlannerTaskPropertyRuleRuleKind, error) {
	vals := map[string]PlannerTaskPropertyRuleRuleKind{
		"bucketrule": PlannerTaskPropertyRuleRuleKindbucketRule,
		"planrule":   PlannerTaskPropertyRuleRuleKindplanRule,
		"taskrule":   PlannerTaskPropertyRuleRuleKindtaskRule,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTaskPropertyRuleRuleKind(input)
	return &out, nil
}
