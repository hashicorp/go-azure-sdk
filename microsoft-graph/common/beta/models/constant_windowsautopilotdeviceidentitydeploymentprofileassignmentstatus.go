package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus string

const (
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatusassignedInSync          WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus = "AssignedInSync"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatusassignedOutOfSync       WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus = "AssignedOutOfSync"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatusassignedUnkownSyncState WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus = "AssignedUnkownSyncState"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatusfailed                  WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus = "Failed"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatusnotAssigned             WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus = "NotAssigned"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatuspending                 WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus = "Pending"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatusunknown                 WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus = "Unknown"
)

func PossibleValuesForWindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus() []string {
	return []string{
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatusassignedInSync),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatusassignedOutOfSync),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatusassignedUnkownSyncState),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatusfailed),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatusnotAssigned),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatuspending),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatusunknown),
	}
}

func parseWindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus(input string) (*WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus, error) {
	vals := map[string]WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus{
		"assignedinsync":          WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatusassignedInSync,
		"assignedoutofsync":       WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatusassignedOutOfSync,
		"assignedunkownsyncstate": WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatusassignedUnkownSyncState,
		"failed":                  WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatusfailed,
		"notassigned":             WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatusnotAssigned,
		"pending":                 WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatuspending,
		"unknown":                 WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentStatus(input)
	return &out, nil
}
