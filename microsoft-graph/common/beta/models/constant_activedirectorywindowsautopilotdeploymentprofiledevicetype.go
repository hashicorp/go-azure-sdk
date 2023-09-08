package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType string

const (
	ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceTypeholoLens       ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType = "HoloLens"
	ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceTypesurfaceHub2    ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType = "SurfaceHub2"
	ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceTypesurfaceHub2S   ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType = "SurfaceHub2S"
	ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceTypevirtualMachine ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType = "VirtualMachine"
	ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceTypewindowsPc      ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType = "WindowsPc"
)

func PossibleValuesForActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType() []string {
	return []string{
		string(ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceTypeholoLens),
		string(ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceTypesurfaceHub2),
		string(ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceTypesurfaceHub2S),
		string(ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceTypevirtualMachine),
		string(ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceTypewindowsPc),
	}
}

func parseActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType(input string) (*ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType, error) {
	vals := map[string]ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType{
		"hololens":       ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceTypeholoLens,
		"surfacehub2":    ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceTypesurfaceHub2,
		"surfacehub2s":   ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceTypesurfaceHub2S,
		"virtualmachine": ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceTypevirtualMachine,
		"windowspc":      ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceTypewindowsPc,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType(input)
	return &out, nil
}
