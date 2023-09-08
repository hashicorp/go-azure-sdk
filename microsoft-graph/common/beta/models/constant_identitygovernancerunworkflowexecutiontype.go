package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceRunWorkflowExecutionType string

const (
	IdentityGovernanceRunWorkflowExecutionTypeonDemand  IdentityGovernanceRunWorkflowExecutionType = "OnDemand"
	IdentityGovernanceRunWorkflowExecutionTypescheduled IdentityGovernanceRunWorkflowExecutionType = "Scheduled"
)

func PossibleValuesForIdentityGovernanceRunWorkflowExecutionType() []string {
	return []string{
		string(IdentityGovernanceRunWorkflowExecutionTypeonDemand),
		string(IdentityGovernanceRunWorkflowExecutionTypescheduled),
	}
}

func parseIdentityGovernanceRunWorkflowExecutionType(input string) (*IdentityGovernanceRunWorkflowExecutionType, error) {
	vals := map[string]IdentityGovernanceRunWorkflowExecutionType{
		"ondemand":  IdentityGovernanceRunWorkflowExecutionTypeonDemand,
		"scheduled": IdentityGovernanceRunWorkflowExecutionTypescheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceRunWorkflowExecutionType(input)
	return &out, nil
}
