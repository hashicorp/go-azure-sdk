package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceRunWorkflowExecutionType string

const (
	IdentityGovernanceRunWorkflowExecutionType_OnDemand  IdentityGovernanceRunWorkflowExecutionType = "onDemand"
	IdentityGovernanceRunWorkflowExecutionType_Scheduled IdentityGovernanceRunWorkflowExecutionType = "scheduled"
)

func PossibleValuesForIdentityGovernanceRunWorkflowExecutionType() []string {
	return []string{
		string(IdentityGovernanceRunWorkflowExecutionType_OnDemand),
		string(IdentityGovernanceRunWorkflowExecutionType_Scheduled),
	}
}

func (s *IdentityGovernanceRunWorkflowExecutionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceRunWorkflowExecutionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceRunWorkflowExecutionType(input string) (*IdentityGovernanceRunWorkflowExecutionType, error) {
	vals := map[string]IdentityGovernanceRunWorkflowExecutionType{
		"ondemand":  IdentityGovernanceRunWorkflowExecutionType_OnDemand,
		"scheduled": IdentityGovernanceRunWorkflowExecutionType_Scheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceRunWorkflowExecutionType(input)
	return &out, nil
}
