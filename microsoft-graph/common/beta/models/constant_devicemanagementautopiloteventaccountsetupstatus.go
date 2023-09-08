package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAutopilotEventAccountSetupStatus string

const (
	DeviceManagementAutopilotEventAccountSetupStatusdisabled           DeviceManagementAutopilotEventAccountSetupStatus = "Disabled"
	DeviceManagementAutopilotEventAccountSetupStatusfailure            DeviceManagementAutopilotEventAccountSetupStatus = "Failure"
	DeviceManagementAutopilotEventAccountSetupStatusinProgress         DeviceManagementAutopilotEventAccountSetupStatus = "InProgress"
	DeviceManagementAutopilotEventAccountSetupStatusnotAttempted       DeviceManagementAutopilotEventAccountSetupStatus = "NotAttempted"
	DeviceManagementAutopilotEventAccountSetupStatussuccess            DeviceManagementAutopilotEventAccountSetupStatus = "Success"
	DeviceManagementAutopilotEventAccountSetupStatussuccessOnRetry     DeviceManagementAutopilotEventAccountSetupStatus = "SuccessOnRetry"
	DeviceManagementAutopilotEventAccountSetupStatussuccessWithTimeout DeviceManagementAutopilotEventAccountSetupStatus = "SuccessWithTimeout"
	DeviceManagementAutopilotEventAccountSetupStatusunknown            DeviceManagementAutopilotEventAccountSetupStatus = "Unknown"
)

func PossibleValuesForDeviceManagementAutopilotEventAccountSetupStatus() []string {
	return []string{
		string(DeviceManagementAutopilotEventAccountSetupStatusdisabled),
		string(DeviceManagementAutopilotEventAccountSetupStatusfailure),
		string(DeviceManagementAutopilotEventAccountSetupStatusinProgress),
		string(DeviceManagementAutopilotEventAccountSetupStatusnotAttempted),
		string(DeviceManagementAutopilotEventAccountSetupStatussuccess),
		string(DeviceManagementAutopilotEventAccountSetupStatussuccessOnRetry),
		string(DeviceManagementAutopilotEventAccountSetupStatussuccessWithTimeout),
		string(DeviceManagementAutopilotEventAccountSetupStatusunknown),
	}
}

func parseDeviceManagementAutopilotEventAccountSetupStatus(input string) (*DeviceManagementAutopilotEventAccountSetupStatus, error) {
	vals := map[string]DeviceManagementAutopilotEventAccountSetupStatus{
		"disabled":           DeviceManagementAutopilotEventAccountSetupStatusdisabled,
		"failure":            DeviceManagementAutopilotEventAccountSetupStatusfailure,
		"inprogress":         DeviceManagementAutopilotEventAccountSetupStatusinProgress,
		"notattempted":       DeviceManagementAutopilotEventAccountSetupStatusnotAttempted,
		"success":            DeviceManagementAutopilotEventAccountSetupStatussuccess,
		"successonretry":     DeviceManagementAutopilotEventAccountSetupStatussuccessOnRetry,
		"successwithtimeout": DeviceManagementAutopilotEventAccountSetupStatussuccessWithTimeout,
		"unknown":            DeviceManagementAutopilotEventAccountSetupStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAutopilotEventAccountSetupStatus(input)
	return &out, nil
}
