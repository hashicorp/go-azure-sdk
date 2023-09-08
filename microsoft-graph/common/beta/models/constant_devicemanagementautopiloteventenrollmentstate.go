package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAutopilotEventEnrollmentState string

const (
	DeviceManagementAutopilotEventEnrollmentStateblocked      DeviceManagementAutopilotEventEnrollmentState = "Blocked"
	DeviceManagementAutopilotEventEnrollmentStateenrolled     DeviceManagementAutopilotEventEnrollmentState = "Enrolled"
	DeviceManagementAutopilotEventEnrollmentStatefailed       DeviceManagementAutopilotEventEnrollmentState = "Failed"
	DeviceManagementAutopilotEventEnrollmentStatenotContacted DeviceManagementAutopilotEventEnrollmentState = "NotContacted"
	DeviceManagementAutopilotEventEnrollmentStatependingReset DeviceManagementAutopilotEventEnrollmentState = "PendingReset"
	DeviceManagementAutopilotEventEnrollmentStateunknown      DeviceManagementAutopilotEventEnrollmentState = "Unknown"
)

func PossibleValuesForDeviceManagementAutopilotEventEnrollmentState() []string {
	return []string{
		string(DeviceManagementAutopilotEventEnrollmentStateblocked),
		string(DeviceManagementAutopilotEventEnrollmentStateenrolled),
		string(DeviceManagementAutopilotEventEnrollmentStatefailed),
		string(DeviceManagementAutopilotEventEnrollmentStatenotContacted),
		string(DeviceManagementAutopilotEventEnrollmentStatependingReset),
		string(DeviceManagementAutopilotEventEnrollmentStateunknown),
	}
}

func parseDeviceManagementAutopilotEventEnrollmentState(input string) (*DeviceManagementAutopilotEventEnrollmentState, error) {
	vals := map[string]DeviceManagementAutopilotEventEnrollmentState{
		"blocked":      DeviceManagementAutopilotEventEnrollmentStateblocked,
		"enrolled":     DeviceManagementAutopilotEventEnrollmentStateenrolled,
		"failed":       DeviceManagementAutopilotEventEnrollmentStatefailed,
		"notcontacted": DeviceManagementAutopilotEventEnrollmentStatenotContacted,
		"pendingreset": DeviceManagementAutopilotEventEnrollmentStatependingReset,
		"unknown":      DeviceManagementAutopilotEventEnrollmentStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAutopilotEventEnrollmentState(input)
	return &out, nil
}
