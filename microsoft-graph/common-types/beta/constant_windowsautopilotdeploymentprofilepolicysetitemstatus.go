package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeploymentProfilePolicySetItemStatus string

const (
	WindowsAutopilotDeploymentProfilePolicySetItemStatus_Error          WindowsAutopilotDeploymentProfilePolicySetItemStatus = "error"
	WindowsAutopilotDeploymentProfilePolicySetItemStatus_NotAssigned    WindowsAutopilotDeploymentProfilePolicySetItemStatus = "notAssigned"
	WindowsAutopilotDeploymentProfilePolicySetItemStatus_PartialSuccess WindowsAutopilotDeploymentProfilePolicySetItemStatus = "partialSuccess"
	WindowsAutopilotDeploymentProfilePolicySetItemStatus_Success        WindowsAutopilotDeploymentProfilePolicySetItemStatus = "success"
	WindowsAutopilotDeploymentProfilePolicySetItemStatus_Unknown        WindowsAutopilotDeploymentProfilePolicySetItemStatus = "unknown"
	WindowsAutopilotDeploymentProfilePolicySetItemStatus_Validating     WindowsAutopilotDeploymentProfilePolicySetItemStatus = "validating"
)

func PossibleValuesForWindowsAutopilotDeploymentProfilePolicySetItemStatus() []string {
	return []string{
		string(WindowsAutopilotDeploymentProfilePolicySetItemStatus_Error),
		string(WindowsAutopilotDeploymentProfilePolicySetItemStatus_NotAssigned),
		string(WindowsAutopilotDeploymentProfilePolicySetItemStatus_PartialSuccess),
		string(WindowsAutopilotDeploymentProfilePolicySetItemStatus_Success),
		string(WindowsAutopilotDeploymentProfilePolicySetItemStatus_Unknown),
		string(WindowsAutopilotDeploymentProfilePolicySetItemStatus_Validating),
	}
}

func (s *WindowsAutopilotDeploymentProfilePolicySetItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAutopilotDeploymentProfilePolicySetItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAutopilotDeploymentProfilePolicySetItemStatus(input string) (*WindowsAutopilotDeploymentProfilePolicySetItemStatus, error) {
	vals := map[string]WindowsAutopilotDeploymentProfilePolicySetItemStatus{
		"error":          WindowsAutopilotDeploymentProfilePolicySetItemStatus_Error,
		"notassigned":    WindowsAutopilotDeploymentProfilePolicySetItemStatus_NotAssigned,
		"partialsuccess": WindowsAutopilotDeploymentProfilePolicySetItemStatus_PartialSuccess,
		"success":        WindowsAutopilotDeploymentProfilePolicySetItemStatus_Success,
		"unknown":        WindowsAutopilotDeploymentProfilePolicySetItemStatus_Unknown,
		"validating":     WindowsAutopilotDeploymentProfilePolicySetItemStatus_Validating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeploymentProfilePolicySetItemStatus(input)
	return &out, nil
}
