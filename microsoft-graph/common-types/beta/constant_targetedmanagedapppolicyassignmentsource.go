package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppPolicyAssignmentSource string

const (
	TargetedManagedAppPolicyAssignmentSource_Direct     TargetedManagedAppPolicyAssignmentSource = "direct"
	TargetedManagedAppPolicyAssignmentSource_PolicySets TargetedManagedAppPolicyAssignmentSource = "policySets"
)

func PossibleValuesForTargetedManagedAppPolicyAssignmentSource() []string {
	return []string{
		string(TargetedManagedAppPolicyAssignmentSource_Direct),
		string(TargetedManagedAppPolicyAssignmentSource_PolicySets),
	}
}

func (s *TargetedManagedAppPolicyAssignmentSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppPolicyAssignmentSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppPolicyAssignmentSource(input string) (*TargetedManagedAppPolicyAssignmentSource, error) {
	vals := map[string]TargetedManagedAppPolicyAssignmentSource{
		"direct":     TargetedManagedAppPolicyAssignmentSource_Direct,
		"policysets": TargetedManagedAppPolicyAssignmentSource_PolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppPolicyAssignmentSource(input)
	return &out, nil
}
