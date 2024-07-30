package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerBucketPropertyRuleRuleKind string

const (
	PlannerBucketPropertyRuleRuleKind_BucketRule PlannerBucketPropertyRuleRuleKind = "bucketRule"
	PlannerBucketPropertyRuleRuleKind_PlanRule   PlannerBucketPropertyRuleRuleKind = "planRule"
	PlannerBucketPropertyRuleRuleKind_TaskRule   PlannerBucketPropertyRuleRuleKind = "taskRule"
)

func PossibleValuesForPlannerBucketPropertyRuleRuleKind() []string {
	return []string{
		string(PlannerBucketPropertyRuleRuleKind_BucketRule),
		string(PlannerBucketPropertyRuleRuleKind_PlanRule),
		string(PlannerBucketPropertyRuleRuleKind_TaskRule),
	}
}

func (s *PlannerBucketPropertyRuleRuleKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerBucketPropertyRuleRuleKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerBucketPropertyRuleRuleKind(input string) (*PlannerBucketPropertyRuleRuleKind, error) {
	vals := map[string]PlannerBucketPropertyRuleRuleKind{
		"bucketrule": PlannerBucketPropertyRuleRuleKind_BucketRule,
		"planrule":   PlannerBucketPropertyRuleRuleKind_PlanRule,
		"taskrule":   PlannerBucketPropertyRuleRuleKind_TaskRule,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerBucketPropertyRuleRuleKind(input)
	return &out, nil
}
