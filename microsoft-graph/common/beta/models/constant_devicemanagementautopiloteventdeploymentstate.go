package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAutopilotEventDeploymentState string

const (
	DeviceManagementAutopilotEventDeploymentStatedisabled           DeviceManagementAutopilotEventDeploymentState = "Disabled"
	DeviceManagementAutopilotEventDeploymentStatefailure            DeviceManagementAutopilotEventDeploymentState = "Failure"
	DeviceManagementAutopilotEventDeploymentStateinProgress         DeviceManagementAutopilotEventDeploymentState = "InProgress"
	DeviceManagementAutopilotEventDeploymentStatenotAttempted       DeviceManagementAutopilotEventDeploymentState = "NotAttempted"
	DeviceManagementAutopilotEventDeploymentStatesuccess            DeviceManagementAutopilotEventDeploymentState = "Success"
	DeviceManagementAutopilotEventDeploymentStatesuccessOnRetry     DeviceManagementAutopilotEventDeploymentState = "SuccessOnRetry"
	DeviceManagementAutopilotEventDeploymentStatesuccessWithTimeout DeviceManagementAutopilotEventDeploymentState = "SuccessWithTimeout"
	DeviceManagementAutopilotEventDeploymentStateunknown            DeviceManagementAutopilotEventDeploymentState = "Unknown"
)

func PossibleValuesForDeviceManagementAutopilotEventDeploymentState() []string {
	return []string{
		string(DeviceManagementAutopilotEventDeploymentStatedisabled),
		string(DeviceManagementAutopilotEventDeploymentStatefailure),
		string(DeviceManagementAutopilotEventDeploymentStateinProgress),
		string(DeviceManagementAutopilotEventDeploymentStatenotAttempted),
		string(DeviceManagementAutopilotEventDeploymentStatesuccess),
		string(DeviceManagementAutopilotEventDeploymentStatesuccessOnRetry),
		string(DeviceManagementAutopilotEventDeploymentStatesuccessWithTimeout),
		string(DeviceManagementAutopilotEventDeploymentStateunknown),
	}
}

func parseDeviceManagementAutopilotEventDeploymentState(input string) (*DeviceManagementAutopilotEventDeploymentState, error) {
	vals := map[string]DeviceManagementAutopilotEventDeploymentState{
		"disabled":           DeviceManagementAutopilotEventDeploymentStatedisabled,
		"failure":            DeviceManagementAutopilotEventDeploymentStatefailure,
		"inprogress":         DeviceManagementAutopilotEventDeploymentStateinProgress,
		"notattempted":       DeviceManagementAutopilotEventDeploymentStatenotAttempted,
		"success":            DeviceManagementAutopilotEventDeploymentStatesuccess,
		"successonretry":     DeviceManagementAutopilotEventDeploymentStatesuccessOnRetry,
		"successwithtimeout": DeviceManagementAutopilotEventDeploymentStatesuccessWithTimeout,
		"unknown":            DeviceManagementAutopilotEventDeploymentStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAutopilotEventDeploymentState(input)
	return &out, nil
}
