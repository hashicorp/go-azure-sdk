package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureADWindowsAutopilotDeploymentProfileDeviceType string

const (
	AzureADWindowsAutopilotDeploymentProfileDeviceTypeholoLens       AzureADWindowsAutopilotDeploymentProfileDeviceType = "HoloLens"
	AzureADWindowsAutopilotDeploymentProfileDeviceTypesurfaceHub2    AzureADWindowsAutopilotDeploymentProfileDeviceType = "SurfaceHub2"
	AzureADWindowsAutopilotDeploymentProfileDeviceTypesurfaceHub2S   AzureADWindowsAutopilotDeploymentProfileDeviceType = "SurfaceHub2S"
	AzureADWindowsAutopilotDeploymentProfileDeviceTypevirtualMachine AzureADWindowsAutopilotDeploymentProfileDeviceType = "VirtualMachine"
	AzureADWindowsAutopilotDeploymentProfileDeviceTypewindowsPc      AzureADWindowsAutopilotDeploymentProfileDeviceType = "WindowsPc"
)

func PossibleValuesForAzureADWindowsAutopilotDeploymentProfileDeviceType() []string {
	return []string{
		string(AzureADWindowsAutopilotDeploymentProfileDeviceTypeholoLens),
		string(AzureADWindowsAutopilotDeploymentProfileDeviceTypesurfaceHub2),
		string(AzureADWindowsAutopilotDeploymentProfileDeviceTypesurfaceHub2S),
		string(AzureADWindowsAutopilotDeploymentProfileDeviceTypevirtualMachine),
		string(AzureADWindowsAutopilotDeploymentProfileDeviceTypewindowsPc),
	}
}

func parseAzureADWindowsAutopilotDeploymentProfileDeviceType(input string) (*AzureADWindowsAutopilotDeploymentProfileDeviceType, error) {
	vals := map[string]AzureADWindowsAutopilotDeploymentProfileDeviceType{
		"hololens":       AzureADWindowsAutopilotDeploymentProfileDeviceTypeholoLens,
		"surfacehub2":    AzureADWindowsAutopilotDeploymentProfileDeviceTypesurfaceHub2,
		"surfacehub2s":   AzureADWindowsAutopilotDeploymentProfileDeviceTypesurfaceHub2S,
		"virtualmachine": AzureADWindowsAutopilotDeploymentProfileDeviceTypevirtualMachine,
		"windowspc":      AzureADWindowsAutopilotDeploymentProfileDeviceTypewindowsPc,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureADWindowsAutopilotDeploymentProfileDeviceType(input)
	return &out, nil
}
