package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDeviceOperationOperationType string

const (
	TeamworkDeviceOperationOperationTypeconfigUpdate                      TeamworkDeviceOperationOperationType = "ConfigUpdate"
	TeamworkDeviceOperationOperationTypedeviceDiagnostics                 TeamworkDeviceOperationOperationType = "DeviceDiagnostics"
	TeamworkDeviceOperationOperationTypedeviceManagementAgentConfigUpdate TeamworkDeviceOperationOperationType = "DeviceManagementAgentConfigUpdate"
	TeamworkDeviceOperationOperationTypedeviceRestart                     TeamworkDeviceOperationOperationType = "DeviceRestart"
	TeamworkDeviceOperationOperationTyperemoteLogin                       TeamworkDeviceOperationOperationType = "RemoteLogin"
	TeamworkDeviceOperationOperationTyperemoteLogout                      TeamworkDeviceOperationOperationType = "RemoteLogout"
	TeamworkDeviceOperationOperationTypesoftwareUpdate                    TeamworkDeviceOperationOperationType = "SoftwareUpdate"
)

func PossibleValuesForTeamworkDeviceOperationOperationType() []string {
	return []string{
		string(TeamworkDeviceOperationOperationTypeconfigUpdate),
		string(TeamworkDeviceOperationOperationTypedeviceDiagnostics),
		string(TeamworkDeviceOperationOperationTypedeviceManagementAgentConfigUpdate),
		string(TeamworkDeviceOperationOperationTypedeviceRestart),
		string(TeamworkDeviceOperationOperationTyperemoteLogin),
		string(TeamworkDeviceOperationOperationTyperemoteLogout),
		string(TeamworkDeviceOperationOperationTypesoftwareUpdate),
	}
}

func parseTeamworkDeviceOperationOperationType(input string) (*TeamworkDeviceOperationOperationType, error) {
	vals := map[string]TeamworkDeviceOperationOperationType{
		"configupdate":                      TeamworkDeviceOperationOperationTypeconfigUpdate,
		"devicediagnostics":                 TeamworkDeviceOperationOperationTypedeviceDiagnostics,
		"devicemanagementagentconfigupdate": TeamworkDeviceOperationOperationTypedeviceManagementAgentConfigUpdate,
		"devicerestart":                     TeamworkDeviceOperationOperationTypedeviceRestart,
		"remotelogin":                       TeamworkDeviceOperationOperationTyperemoteLogin,
		"remotelogout":                      TeamworkDeviceOperationOperationTyperemoteLogout,
		"softwareupdate":                    TeamworkDeviceOperationOperationTypesoftwareUpdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkDeviceOperationOperationType(input)
	return &out, nil
}
