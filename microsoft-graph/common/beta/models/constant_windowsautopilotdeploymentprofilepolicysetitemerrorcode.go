package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeploymentProfilePolicySetItemErrorCode string

const (
	WindowsAutopilotDeploymentProfilePolicySetItemErrorCodedeleted      WindowsAutopilotDeploymentProfilePolicySetItemErrorCode = "Deleted"
	WindowsAutopilotDeploymentProfilePolicySetItemErrorCodenoError      WindowsAutopilotDeploymentProfilePolicySetItemErrorCode = "NoError"
	WindowsAutopilotDeploymentProfilePolicySetItemErrorCodenotFound     WindowsAutopilotDeploymentProfilePolicySetItemErrorCode = "NotFound"
	WindowsAutopilotDeploymentProfilePolicySetItemErrorCodeunauthorized WindowsAutopilotDeploymentProfilePolicySetItemErrorCode = "Unauthorized"
)

func PossibleValuesForWindowsAutopilotDeploymentProfilePolicySetItemErrorCode() []string {
	return []string{
		string(WindowsAutopilotDeploymentProfilePolicySetItemErrorCodedeleted),
		string(WindowsAutopilotDeploymentProfilePolicySetItemErrorCodenoError),
		string(WindowsAutopilotDeploymentProfilePolicySetItemErrorCodenotFound),
		string(WindowsAutopilotDeploymentProfilePolicySetItemErrorCodeunauthorized),
	}
}

func parseWindowsAutopilotDeploymentProfilePolicySetItemErrorCode(input string) (*WindowsAutopilotDeploymentProfilePolicySetItemErrorCode, error) {
	vals := map[string]WindowsAutopilotDeploymentProfilePolicySetItemErrorCode{
		"deleted":      WindowsAutopilotDeploymentProfilePolicySetItemErrorCodedeleted,
		"noerror":      WindowsAutopilotDeploymentProfilePolicySetItemErrorCodenoError,
		"notfound":     WindowsAutopilotDeploymentProfilePolicySetItemErrorCodenotFound,
		"unauthorized": WindowsAutopilotDeploymentProfilePolicySetItemErrorCodeunauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeploymentProfilePolicySetItemErrorCode(input)
	return &out, nil
}
