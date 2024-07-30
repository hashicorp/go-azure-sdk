package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureADWindowsAutopilotDeploymentProfileDeviceType string

const (
	AzureADWindowsAutopilotDeploymentProfileDeviceType_HoloLens       AzureADWindowsAutopilotDeploymentProfileDeviceType = "holoLens"
	AzureADWindowsAutopilotDeploymentProfileDeviceType_SurfaceHub2    AzureADWindowsAutopilotDeploymentProfileDeviceType = "surfaceHub2"
	AzureADWindowsAutopilotDeploymentProfileDeviceType_SurfaceHub2S   AzureADWindowsAutopilotDeploymentProfileDeviceType = "surfaceHub2S"
	AzureADWindowsAutopilotDeploymentProfileDeviceType_VirtualMachine AzureADWindowsAutopilotDeploymentProfileDeviceType = "virtualMachine"
	AzureADWindowsAutopilotDeploymentProfileDeviceType_WindowsPc      AzureADWindowsAutopilotDeploymentProfileDeviceType = "windowsPc"
)

func PossibleValuesForAzureADWindowsAutopilotDeploymentProfileDeviceType() []string {
	return []string{
		string(AzureADWindowsAutopilotDeploymentProfileDeviceType_HoloLens),
		string(AzureADWindowsAutopilotDeploymentProfileDeviceType_SurfaceHub2),
		string(AzureADWindowsAutopilotDeploymentProfileDeviceType_SurfaceHub2S),
		string(AzureADWindowsAutopilotDeploymentProfileDeviceType_VirtualMachine),
		string(AzureADWindowsAutopilotDeploymentProfileDeviceType_WindowsPc),
	}
}

func (s *AzureADWindowsAutopilotDeploymentProfileDeviceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAzureADWindowsAutopilotDeploymentProfileDeviceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAzureADWindowsAutopilotDeploymentProfileDeviceType(input string) (*AzureADWindowsAutopilotDeploymentProfileDeviceType, error) {
	vals := map[string]AzureADWindowsAutopilotDeploymentProfileDeviceType{
		"hololens":       AzureADWindowsAutopilotDeploymentProfileDeviceType_HoloLens,
		"surfacehub2":    AzureADWindowsAutopilotDeploymentProfileDeviceType_SurfaceHub2,
		"surfacehub2s":   AzureADWindowsAutopilotDeploymentProfileDeviceType_SurfaceHub2S,
		"virtualmachine": AzureADWindowsAutopilotDeploymentProfileDeviceType_VirtualMachine,
		"windowspc":      AzureADWindowsAutopilotDeploymentProfileDeviceType_WindowsPc,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureADWindowsAutopilotDeploymentProfileDeviceType(input)
	return &out, nil
}
