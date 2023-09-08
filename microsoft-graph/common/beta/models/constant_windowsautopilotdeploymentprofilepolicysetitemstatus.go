package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeploymentProfilePolicySetItemStatus string

const (
	WindowsAutopilotDeploymentProfilePolicySetItemStatuserror          WindowsAutopilotDeploymentProfilePolicySetItemStatus = "Error"
	WindowsAutopilotDeploymentProfilePolicySetItemStatusnotAssigned    WindowsAutopilotDeploymentProfilePolicySetItemStatus = "NotAssigned"
	WindowsAutopilotDeploymentProfilePolicySetItemStatuspartialSuccess WindowsAutopilotDeploymentProfilePolicySetItemStatus = "PartialSuccess"
	WindowsAutopilotDeploymentProfilePolicySetItemStatussuccess        WindowsAutopilotDeploymentProfilePolicySetItemStatus = "Success"
	WindowsAutopilotDeploymentProfilePolicySetItemStatusunknown        WindowsAutopilotDeploymentProfilePolicySetItemStatus = "Unknown"
	WindowsAutopilotDeploymentProfilePolicySetItemStatusvalidating     WindowsAutopilotDeploymentProfilePolicySetItemStatus = "Validating"
)

func PossibleValuesForWindowsAutopilotDeploymentProfilePolicySetItemStatus() []string {
	return []string{
		string(WindowsAutopilotDeploymentProfilePolicySetItemStatuserror),
		string(WindowsAutopilotDeploymentProfilePolicySetItemStatusnotAssigned),
		string(WindowsAutopilotDeploymentProfilePolicySetItemStatuspartialSuccess),
		string(WindowsAutopilotDeploymentProfilePolicySetItemStatussuccess),
		string(WindowsAutopilotDeploymentProfilePolicySetItemStatusunknown),
		string(WindowsAutopilotDeploymentProfilePolicySetItemStatusvalidating),
	}
}

func parseWindowsAutopilotDeploymentProfilePolicySetItemStatus(input string) (*WindowsAutopilotDeploymentProfilePolicySetItemStatus, error) {
	vals := map[string]WindowsAutopilotDeploymentProfilePolicySetItemStatus{
		"error":          WindowsAutopilotDeploymentProfilePolicySetItemStatuserror,
		"notassigned":    WindowsAutopilotDeploymentProfilePolicySetItemStatusnotAssigned,
		"partialsuccess": WindowsAutopilotDeploymentProfilePolicySetItemStatuspartialSuccess,
		"success":        WindowsAutopilotDeploymentProfilePolicySetItemStatussuccess,
		"unknown":        WindowsAutopilotDeploymentProfilePolicySetItemStatusunknown,
		"validating":     WindowsAutopilotDeploymentProfilePolicySetItemStatusvalidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeploymentProfilePolicySetItemStatus(input)
	return &out, nil
}
