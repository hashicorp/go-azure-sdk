package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskPropertyRuleRuleKind string

const (
	PlannerTaskPropertyRuleRuleKind_BucketRule PlannerTaskPropertyRuleRuleKind = "bucketRule"
	PlannerTaskPropertyRuleRuleKind_PlanRule   PlannerTaskPropertyRuleRuleKind = "planRule"
	PlannerTaskPropertyRuleRuleKind_TaskRule   PlannerTaskPropertyRuleRuleKind = "taskRule"
)

func PossibleValuesForPlannerTaskPropertyRuleRuleKind() []string {
	return []string{
		string(PlannerTaskPropertyRuleRuleKind_BucketRule),
		string(PlannerTaskPropertyRuleRuleKind_PlanRule),
		string(PlannerTaskPropertyRuleRuleKind_TaskRule),
	}
}

func (s *PlannerTaskPropertyRuleRuleKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerTaskPropertyRuleRuleKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerTaskPropertyRuleRuleKind(input string) (*PlannerTaskPropertyRuleRuleKind, error) {
	vals := map[string]PlannerTaskPropertyRuleRuleKind{
		"bucketrule": PlannerTaskPropertyRuleRuleKind_BucketRule,
		"planrule":   PlannerTaskPropertyRuleRuleKind_PlanRule,
		"taskrule":   PlannerTaskPropertyRuleRuleKind_TaskRule,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTaskPropertyRuleRuleKind(input)
	return &out, nil
}
