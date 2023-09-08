package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptDeviceHealthScriptType string

const (
	DeviceHealthScriptDeviceHealthScriptTypedeviceHealthScript     DeviceHealthScriptDeviceHealthScriptType = "DeviceHealthScript"
	DeviceHealthScriptDeviceHealthScriptTypemanagedInstallerScript DeviceHealthScriptDeviceHealthScriptType = "ManagedInstallerScript"
)

func PossibleValuesForDeviceHealthScriptDeviceHealthScriptType() []string {
	return []string{
		string(DeviceHealthScriptDeviceHealthScriptTypedeviceHealthScript),
		string(DeviceHealthScriptDeviceHealthScriptTypemanagedInstallerScript),
	}
}

func parseDeviceHealthScriptDeviceHealthScriptType(input string) (*DeviceHealthScriptDeviceHealthScriptType, error) {
	vals := map[string]DeviceHealthScriptDeviceHealthScriptType{
		"devicehealthscript":     DeviceHealthScriptDeviceHealthScriptTypedeviceHealthScript,
		"managedinstallerscript": DeviceHealthScriptDeviceHealthScriptTypemanagedInstallerScript,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthScriptDeviceHealthScriptType(input)
	return &out, nil
}
