package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDeviceOperationOperationType string

const (
	TeamworkDeviceOperationOperationType_ConfigUpdate                      TeamworkDeviceOperationOperationType = "configUpdate"
	TeamworkDeviceOperationOperationType_DeviceDiagnostics                 TeamworkDeviceOperationOperationType = "deviceDiagnostics"
	TeamworkDeviceOperationOperationType_DeviceManagementAgentConfigUpdate TeamworkDeviceOperationOperationType = "deviceManagementAgentConfigUpdate"
	TeamworkDeviceOperationOperationType_DeviceRestart                     TeamworkDeviceOperationOperationType = "deviceRestart"
	TeamworkDeviceOperationOperationType_RemoteLogin                       TeamworkDeviceOperationOperationType = "remoteLogin"
	TeamworkDeviceOperationOperationType_RemoteLogout                      TeamworkDeviceOperationOperationType = "remoteLogout"
	TeamworkDeviceOperationOperationType_SoftwareUpdate                    TeamworkDeviceOperationOperationType = "softwareUpdate"
)

func PossibleValuesForTeamworkDeviceOperationOperationType() []string {
	return []string{
		string(TeamworkDeviceOperationOperationType_ConfigUpdate),
		string(TeamworkDeviceOperationOperationType_DeviceDiagnostics),
		string(TeamworkDeviceOperationOperationType_DeviceManagementAgentConfigUpdate),
		string(TeamworkDeviceOperationOperationType_DeviceRestart),
		string(TeamworkDeviceOperationOperationType_RemoteLogin),
		string(TeamworkDeviceOperationOperationType_RemoteLogout),
		string(TeamworkDeviceOperationOperationType_SoftwareUpdate),
	}
}

func (s *TeamworkDeviceOperationOperationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkDeviceOperationOperationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkDeviceOperationOperationType(input string) (*TeamworkDeviceOperationOperationType, error) {
	vals := map[string]TeamworkDeviceOperationOperationType{
		"configupdate":                      TeamworkDeviceOperationOperationType_ConfigUpdate,
		"devicediagnostics":                 TeamworkDeviceOperationOperationType_DeviceDiagnostics,
		"devicemanagementagentconfigupdate": TeamworkDeviceOperationOperationType_DeviceManagementAgentConfigUpdate,
		"devicerestart":                     TeamworkDeviceOperationOperationType_DeviceRestart,
		"remotelogin":                       TeamworkDeviceOperationOperationType_RemoteLogin,
		"remotelogout":                      TeamworkDeviceOperationOperationType_RemoteLogout,
		"softwareupdate":                    TeamworkDeviceOperationOperationType_SoftwareUpdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkDeviceOperationOperationType(input)
	return &out, nil
}
