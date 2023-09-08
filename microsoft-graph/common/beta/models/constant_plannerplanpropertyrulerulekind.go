package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanPropertyRuleRuleKind string

const (
	PlannerPlanPropertyRuleRuleKindbucketRule PlannerPlanPropertyRuleRuleKind = "BucketRule"
	PlannerPlanPropertyRuleRuleKindplanRule   PlannerPlanPropertyRuleRuleKind = "PlanRule"
	PlannerPlanPropertyRuleRuleKindtaskRule   PlannerPlanPropertyRuleRuleKind = "TaskRule"
)

func PossibleValuesForPlannerPlanPropertyRuleRuleKind() []string {
	return []string{
		string(PlannerPlanPropertyRuleRuleKindbucketRule),
		string(PlannerPlanPropertyRuleRuleKindplanRule),
		string(PlannerPlanPropertyRuleRuleKindtaskRule),
	}
}

func parsePlannerPlanPropertyRuleRuleKind(input string) (*PlannerPlanPropertyRuleRuleKind, error) {
	vals := map[string]PlannerPlanPropertyRuleRuleKind{
		"bucketrule": PlannerPlanPropertyRuleRuleKindbucketRule,
		"planrule":   PlannerPlanPropertyRuleRuleKindplanRule,
		"taskrule":   PlannerPlanPropertyRuleRuleKindtaskRule,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPlanPropertyRuleRuleKind(input)
	return &out, nil
}
