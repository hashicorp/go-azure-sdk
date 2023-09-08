package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAutopilotEventDeviceSetupStatus string

const (
	DeviceManagementAutopilotEventDeviceSetupStatusdisabled           DeviceManagementAutopilotEventDeviceSetupStatus = "Disabled"
	DeviceManagementAutopilotEventDeviceSetupStatusfailure            DeviceManagementAutopilotEventDeviceSetupStatus = "Failure"
	DeviceManagementAutopilotEventDeviceSetupStatusinProgress         DeviceManagementAutopilotEventDeviceSetupStatus = "InProgress"
	DeviceManagementAutopilotEventDeviceSetupStatusnotAttempted       DeviceManagementAutopilotEventDeviceSetupStatus = "NotAttempted"
	DeviceManagementAutopilotEventDeviceSetupStatussuccess            DeviceManagementAutopilotEventDeviceSetupStatus = "Success"
	DeviceManagementAutopilotEventDeviceSetupStatussuccessOnRetry     DeviceManagementAutopilotEventDeviceSetupStatus = "SuccessOnRetry"
	DeviceManagementAutopilotEventDeviceSetupStatussuccessWithTimeout DeviceManagementAutopilotEventDeviceSetupStatus = "SuccessWithTimeout"
	DeviceManagementAutopilotEventDeviceSetupStatusunknown            DeviceManagementAutopilotEventDeviceSetupStatus = "Unknown"
)

func PossibleValuesForDeviceManagementAutopilotEventDeviceSetupStatus() []string {
	return []string{
		string(DeviceManagementAutopilotEventDeviceSetupStatusdisabled),
		string(DeviceManagementAutopilotEventDeviceSetupStatusfailure),
		string(DeviceManagementAutopilotEventDeviceSetupStatusinProgress),
		string(DeviceManagementAutopilotEventDeviceSetupStatusnotAttempted),
		string(DeviceManagementAutopilotEventDeviceSetupStatussuccess),
		string(DeviceManagementAutopilotEventDeviceSetupStatussuccessOnRetry),
		string(DeviceManagementAutopilotEventDeviceSetupStatussuccessWithTimeout),
		string(DeviceManagementAutopilotEventDeviceSetupStatusunknown),
	}
}

func parseDeviceManagementAutopilotEventDeviceSetupStatus(input string) (*DeviceManagementAutopilotEventDeviceSetupStatus, error) {
	vals := map[string]DeviceManagementAutopilotEventDeviceSetupStatus{
		"disabled":           DeviceManagementAutopilotEventDeviceSetupStatusdisabled,
		"failure":            DeviceManagementAutopilotEventDeviceSetupStatusfailure,
		"inprogress":         DeviceManagementAutopilotEventDeviceSetupStatusinProgress,
		"notattempted":       DeviceManagementAutopilotEventDeviceSetupStatusnotAttempted,
		"success":            DeviceManagementAutopilotEventDeviceSetupStatussuccess,
		"successonretry":     DeviceManagementAutopilotEventDeviceSetupStatussuccessOnRetry,
		"successwithtimeout": DeviceManagementAutopilotEventDeviceSetupStatussuccessWithTimeout,
		"unknown":            DeviceManagementAutopilotEventDeviceSetupStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAutopilotEventDeviceSetupStatus(input)
	return &out, nil
}
