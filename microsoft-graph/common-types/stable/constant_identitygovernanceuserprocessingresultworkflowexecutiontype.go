package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceUserProcessingResultWorkflowExecutionType string

const (
	IdentityGovernanceUserProcessingResultWorkflowExecutionType_OnDemand  IdentityGovernanceUserProcessingResultWorkflowExecutionType = "onDemand"
	IdentityGovernanceUserProcessingResultWorkflowExecutionType_Scheduled IdentityGovernanceUserProcessingResultWorkflowExecutionType = "scheduled"
)

func PossibleValuesForIdentityGovernanceUserProcessingResultWorkflowExecutionType() []string {
	return []string{
		string(IdentityGovernanceUserProcessingResultWorkflowExecutionType_OnDemand),
		string(IdentityGovernanceUserProcessingResultWorkflowExecutionType_Scheduled),
	}
}

func (s *IdentityGovernanceUserProcessingResultWorkflowExecutionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceUserProcessingResultWorkflowExecutionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceUserProcessingResultWorkflowExecutionType(input string) (*IdentityGovernanceUserProcessingResultWorkflowExecutionType, error) {
	vals := map[string]IdentityGovernanceUserProcessingResultWorkflowExecutionType{
		"ondemand":  IdentityGovernanceUserProcessingResultWorkflowExecutionType_OnDemand,
		"scheduled": IdentityGovernanceUserProcessingResultWorkflowExecutionType_Scheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceUserProcessingResultWorkflowExecutionType(input)
	return &out, nil
}
