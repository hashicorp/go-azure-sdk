package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeploymentProfileAssignmentSource string

const (
	WindowsAutopilotDeploymentProfileAssignmentSourcedirect     WindowsAutopilotDeploymentProfileAssignmentSource = "Direct"
	WindowsAutopilotDeploymentProfileAssignmentSourcepolicySets WindowsAutopilotDeploymentProfileAssignmentSource = "PolicySets"
)

func PossibleValuesForWindowsAutopilotDeploymentProfileAssignmentSource() []string {
	return []string{
		string(WindowsAutopilotDeploymentProfileAssignmentSourcedirect),
		string(WindowsAutopilotDeploymentProfileAssignmentSourcepolicySets),
	}
}

func parseWindowsAutopilotDeploymentProfileAssignmentSource(input string) (*WindowsAutopilotDeploymentProfileAssignmentSource, error) {
	vals := map[string]WindowsAutopilotDeploymentProfileAssignmentSource{
		"direct":     WindowsAutopilotDeploymentProfileAssignmentSourcedirect,
		"policysets": WindowsAutopilotDeploymentProfileAssignmentSourcepolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeploymentProfileAssignmentSource(input)
	return &out, nil
}
