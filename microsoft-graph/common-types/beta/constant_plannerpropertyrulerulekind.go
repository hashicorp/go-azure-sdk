package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPropertyRuleRuleKind string

const (
	PlannerPropertyRuleRuleKind_BucketRule PlannerPropertyRuleRuleKind = "bucketRule"
	PlannerPropertyRuleRuleKind_PlanRule   PlannerPropertyRuleRuleKind = "planRule"
	PlannerPropertyRuleRuleKind_TaskRule   PlannerPropertyRuleRuleKind = "taskRule"
)

func PossibleValuesForPlannerPropertyRuleRuleKind() []string {
	return []string{
		string(PlannerPropertyRuleRuleKind_BucketRule),
		string(PlannerPropertyRuleRuleKind_PlanRule),
		string(PlannerPropertyRuleRuleKind_TaskRule),
	}
}

func (s *PlannerPropertyRuleRuleKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerPropertyRuleRuleKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerPropertyRuleRuleKind(input string) (*PlannerPropertyRuleRuleKind, error) {
	vals := map[string]PlannerPropertyRuleRuleKind{
		"bucketrule": PlannerPropertyRuleRuleKind_BucketRule,
		"planrule":   PlannerPropertyRuleRuleKind_PlanRule,
		"taskrule":   PlannerPropertyRuleRuleKind_TaskRule,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPropertyRuleRuleKind(input)
	return &out, nil
}
