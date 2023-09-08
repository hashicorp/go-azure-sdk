package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerBucketPropertyRuleRuleKind string

const (
	PlannerBucketPropertyRuleRuleKindbucketRule PlannerBucketPropertyRuleRuleKind = "BucketRule"
	PlannerBucketPropertyRuleRuleKindplanRule   PlannerBucketPropertyRuleRuleKind = "PlanRule"
	PlannerBucketPropertyRuleRuleKindtaskRule   PlannerBucketPropertyRuleRuleKind = "TaskRule"
)

func PossibleValuesForPlannerBucketPropertyRuleRuleKind() []string {
	return []string{
		string(PlannerBucketPropertyRuleRuleKindbucketRule),
		string(PlannerBucketPropertyRuleRuleKindplanRule),
		string(PlannerBucketPropertyRuleRuleKindtaskRule),
	}
}

func parsePlannerBucketPropertyRuleRuleKind(input string) (*PlannerBucketPropertyRuleRuleKind, error) {
	vals := map[string]PlannerBucketPropertyRuleRuleKind{
		"bucketrule": PlannerBucketPropertyRuleRuleKindbucketRule,
		"planrule":   PlannerBucketPropertyRuleRuleKindplanRule,
		"taskrule":   PlannerBucketPropertyRuleRuleKindtaskRule,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerBucketPropertyRuleRuleKind(input)
	return &out, nil
}
