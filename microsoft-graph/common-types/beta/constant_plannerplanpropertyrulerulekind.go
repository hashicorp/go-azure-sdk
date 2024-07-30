package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanPropertyRuleRuleKind string

const (
	PlannerPlanPropertyRuleRuleKind_BucketRule PlannerPlanPropertyRuleRuleKind = "bucketRule"
	PlannerPlanPropertyRuleRuleKind_PlanRule   PlannerPlanPropertyRuleRuleKind = "planRule"
	PlannerPlanPropertyRuleRuleKind_TaskRule   PlannerPlanPropertyRuleRuleKind = "taskRule"
)

func PossibleValuesForPlannerPlanPropertyRuleRuleKind() []string {
	return []string{
		string(PlannerPlanPropertyRuleRuleKind_BucketRule),
		string(PlannerPlanPropertyRuleRuleKind_PlanRule),
		string(PlannerPlanPropertyRuleRuleKind_TaskRule),
	}
}

func (s *PlannerPlanPropertyRuleRuleKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerPlanPropertyRuleRuleKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerPlanPropertyRuleRuleKind(input string) (*PlannerPlanPropertyRuleRuleKind, error) {
	vals := map[string]PlannerPlanPropertyRuleRuleKind{
		"bucketrule": PlannerPlanPropertyRuleRuleKind_BucketRule,
		"planrule":   PlannerPlanPropertyRuleRuleKind_PlanRule,
		"taskrule":   PlannerPlanPropertyRuleRuleKind_TaskRule,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPlanPropertyRuleRuleKind(input)
	return &out, nil
}
