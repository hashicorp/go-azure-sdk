package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeploymentProfileDeviceType string

const (
	WindowsAutopilotDeploymentProfileDeviceTypeholoLens       WindowsAutopilotDeploymentProfileDeviceType = "HoloLens"
	WindowsAutopilotDeploymentProfileDeviceTypesurfaceHub2    WindowsAutopilotDeploymentProfileDeviceType = "SurfaceHub2"
	WindowsAutopilotDeploymentProfileDeviceTypesurfaceHub2S   WindowsAutopilotDeploymentProfileDeviceType = "SurfaceHub2S"
	WindowsAutopilotDeploymentProfileDeviceTypevirtualMachine WindowsAutopilotDeploymentProfileDeviceType = "VirtualMachine"
	WindowsAutopilotDeploymentProfileDeviceTypewindowsPc      WindowsAutopilotDeploymentProfileDeviceType = "WindowsPc"
)

func PossibleValuesForWindowsAutopilotDeploymentProfileDeviceType() []string {
	return []string{
		string(WindowsAutopilotDeploymentProfileDeviceTypeholoLens),
		string(WindowsAutopilotDeploymentProfileDeviceTypesurfaceHub2),
		string(WindowsAutopilotDeploymentProfileDeviceTypesurfaceHub2S),
		string(WindowsAutopilotDeploymentProfileDeviceTypevirtualMachine),
		string(WindowsAutopilotDeploymentProfileDeviceTypewindowsPc),
	}
}

func parseWindowsAutopilotDeploymentProfileDeviceType(input string) (*WindowsAutopilotDeploymentProfileDeviceType, error) {
	vals := map[string]WindowsAutopilotDeploymentProfileDeviceType{
		"hololens":       WindowsAutopilotDeploymentProfileDeviceTypeholoLens,
		"surfacehub2":    WindowsAutopilotDeploymentProfileDeviceTypesurfaceHub2,
		"surfacehub2s":   WindowsAutopilotDeploymentProfileDeviceTypesurfaceHub2S,
		"virtualmachine": WindowsAutopilotDeploymentProfileDeviceTypevirtualMachine,
		"windowspc":      WindowsAutopilotDeploymentProfileDeviceTypewindowsPc,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeploymentProfileDeviceType(input)
	return &out, nil
}
