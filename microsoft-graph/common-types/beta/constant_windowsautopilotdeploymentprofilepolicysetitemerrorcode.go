package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeploymentProfilePolicySetItemErrorCode string

const (
	WindowsAutopilotDeploymentProfilePolicySetItemErrorCode_Deleted      WindowsAutopilotDeploymentProfilePolicySetItemErrorCode = "deleted"
	WindowsAutopilotDeploymentProfilePolicySetItemErrorCode_NoError      WindowsAutopilotDeploymentProfilePolicySetItemErrorCode = "noError"
	WindowsAutopilotDeploymentProfilePolicySetItemErrorCode_NotFound     WindowsAutopilotDeploymentProfilePolicySetItemErrorCode = "notFound"
	WindowsAutopilotDeploymentProfilePolicySetItemErrorCode_Unauthorized WindowsAutopilotDeploymentProfilePolicySetItemErrorCode = "unauthorized"
)

func PossibleValuesForWindowsAutopilotDeploymentProfilePolicySetItemErrorCode() []string {
	return []string{
		string(WindowsAutopilotDeploymentProfilePolicySetItemErrorCode_Deleted),
		string(WindowsAutopilotDeploymentProfilePolicySetItemErrorCode_NoError),
		string(WindowsAutopilotDeploymentProfilePolicySetItemErrorCode_NotFound),
		string(WindowsAutopilotDeploymentProfilePolicySetItemErrorCode_Unauthorized),
	}
}

func (s *WindowsAutopilotDeploymentProfilePolicySetItemErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAutopilotDeploymentProfilePolicySetItemErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAutopilotDeploymentProfilePolicySetItemErrorCode(input string) (*WindowsAutopilotDeploymentProfilePolicySetItemErrorCode, error) {
	vals := map[string]WindowsAutopilotDeploymentProfilePolicySetItemErrorCode{
		"deleted":      WindowsAutopilotDeploymentProfilePolicySetItemErrorCode_Deleted,
		"noerror":      WindowsAutopilotDeploymentProfilePolicySetItemErrorCode_NoError,
		"notfound":     WindowsAutopilotDeploymentProfilePolicySetItemErrorCode_NotFound,
		"unauthorized": WindowsAutopilotDeploymentProfilePolicySetItemErrorCode_Unauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeploymentProfilePolicySetItemErrorCode(input)
	return &out, nil
}
