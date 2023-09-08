package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceInstallStateInstallState string

const (
	DeviceInstallStateInstallStatefailed          DeviceInstallStateInstallState = "Failed"
	DeviceInstallStateInstallStateinstalled       DeviceInstallStateInstallState = "Installed"
	DeviceInstallStateInstallStatenotApplicable   DeviceInstallStateInstallState = "NotApplicable"
	DeviceInstallStateInstallStatenotInstalled    DeviceInstallStateInstallState = "NotInstalled"
	DeviceInstallStateInstallStateuninstallFailed DeviceInstallStateInstallState = "UninstallFailed"
	DeviceInstallStateInstallStateunknown         DeviceInstallStateInstallState = "Unknown"
)

func PossibleValuesForDeviceInstallStateInstallState() []string {
	return []string{
		string(DeviceInstallStateInstallStatefailed),
		string(DeviceInstallStateInstallStateinstalled),
		string(DeviceInstallStateInstallStatenotApplicable),
		string(DeviceInstallStateInstallStatenotInstalled),
		string(DeviceInstallStateInstallStateuninstallFailed),
		string(DeviceInstallStateInstallStateunknown),
	}
}

func parseDeviceInstallStateInstallState(input string) (*DeviceInstallStateInstallState, error) {
	vals := map[string]DeviceInstallStateInstallState{
		"failed":          DeviceInstallStateInstallStatefailed,
		"installed":       DeviceInstallStateInstallStateinstalled,
		"notapplicable":   DeviceInstallStateInstallStatenotApplicable,
		"notinstalled":    DeviceInstallStateInstallStatenotInstalled,
		"uninstallfailed": DeviceInstallStateInstallStateuninstallFailed,
		"unknown":         DeviceInstallStateInstallStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceInstallStateInstallState(input)
	return &out, nil
}
