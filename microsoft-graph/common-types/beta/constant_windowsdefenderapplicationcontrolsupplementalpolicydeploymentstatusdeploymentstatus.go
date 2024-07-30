package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus string

const (
	WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus_NotAuthorizedByToken WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus = "notAuthorizedByToken"
	WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus_PolicyNotFound       WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus = "policyNotFound"
	WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus_Success              WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus = "success"
	WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus_TokenError           WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus = "tokenError"
	WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus_Unknown              WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus = "unknown"
)

func PossibleValuesForWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus() []string {
	return []string{
		string(WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus_NotAuthorizedByToken),
		string(WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus_PolicyNotFound),
		string(WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus_Success),
		string(WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus_TokenError),
		string(WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus_Unknown),
	}
}

func (s *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus(input string) (*WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus, error) {
	vals := map[string]WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus{
		"notauthorizedbytoken": WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus_NotAuthorizedByToken,
		"policynotfound":       WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus_PolicyNotFound,
		"success":              WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus_Success,
		"tokenerror":           WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus_TokenError,
		"unknown":              WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusDeploymentStatus(input)
	return &out, nil
}
