package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus string

const (
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_AssignedInSync          WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus = "assignedInSync"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_AssignedOutOfSync       WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus = "assignedOutOfSync"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_AssignedUnkownSyncState WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus = "assignedUnkownSyncState"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_Failed                  WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus = "failed"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_NotAssigned             WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus = "notAssigned"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_Pending                 WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus = "pending"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_Unknown                 WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus = "unknown"
)

func PossibleValuesForWindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus() []string {
	return []string{
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_AssignedInSync),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_AssignedOutOfSync),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_AssignedUnkownSyncState),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_Failed),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_NotAssigned),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_Pending),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_Unknown),
	}
}

func (s *WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus(input string) (*WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus, error) {
	vals := map[string]WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus{
		"assignedinsync":          WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_AssignedInSync,
		"assignedoutofsync":       WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_AssignedOutOfSync,
		"assignedunkownsyncstate": WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_AssignedUnkownSyncState,
		"failed":                  WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_Failed,
		"notassigned":             WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_NotAssigned,
		"pending":                 WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_Pending,
		"unknown":                 WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus(input)
	return &out, nil
}
