package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceUserProcessingResultWorkflowExecutionType string

const (
	IdentityGovernanceUserProcessingResultWorkflowExecutionTypeonDemand  IdentityGovernanceUserProcessingResultWorkflowExecutionType = "OnDemand"
	IdentityGovernanceUserProcessingResultWorkflowExecutionTypescheduled IdentityGovernanceUserProcessingResultWorkflowExecutionType = "Scheduled"
)

func PossibleValuesForIdentityGovernanceUserProcessingResultWorkflowExecutionType() []string {
	return []string{
		string(IdentityGovernanceUserProcessingResultWorkflowExecutionTypeonDemand),
		string(IdentityGovernanceUserProcessingResultWorkflowExecutionTypescheduled),
	}
}

func parseIdentityGovernanceUserProcessingResultWorkflowExecutionType(input string) (*IdentityGovernanceUserProcessingResultWorkflowExecutionType, error) {
	vals := map[string]IdentityGovernanceUserProcessingResultWorkflowExecutionType{
		"ondemand":  IdentityGovernanceUserProcessingResultWorkflowExecutionTypeonDemand,
		"scheduled": IdentityGovernanceUserProcessingResultWorkflowExecutionTypescheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceUserProcessingResultWorkflowExecutionType(input)
	return &out, nil
}
