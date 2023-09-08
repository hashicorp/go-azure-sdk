package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPropertyRuleRuleKind string

const (
	PlannerPropertyRuleRuleKindbucketRule PlannerPropertyRuleRuleKind = "BucketRule"
	PlannerPropertyRuleRuleKindplanRule   PlannerPropertyRuleRuleKind = "PlanRule"
	PlannerPropertyRuleRuleKindtaskRule   PlannerPropertyRuleRuleKind = "TaskRule"
)

func PossibleValuesForPlannerPropertyRuleRuleKind() []string {
	return []string{
		string(PlannerPropertyRuleRuleKindbucketRule),
		string(PlannerPropertyRuleRuleKindplanRule),
		string(PlannerPropertyRuleRuleKindtaskRule),
	}
}

func parsePlannerPropertyRuleRuleKind(input string) (*PlannerPropertyRuleRuleKind, error) {
	vals := map[string]PlannerPropertyRuleRuleKind{
		"bucketrule": PlannerPropertyRuleRuleKindbucketRule,
		"planrule":   PlannerPropertyRuleRuleKindplanRule,
		"taskrule":   PlannerPropertyRuleRuleKindtaskRule,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPropertyRuleRuleKind(input)
	return &out, nil
}
