package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeploymentProfileDeviceType string

const (
	WindowsAutopilotDeploymentProfileDeviceType_HoloLens       WindowsAutopilotDeploymentProfileDeviceType = "holoLens"
	WindowsAutopilotDeploymentProfileDeviceType_SurfaceHub2    WindowsAutopilotDeploymentProfileDeviceType = "surfaceHub2"
	WindowsAutopilotDeploymentProfileDeviceType_SurfaceHub2S   WindowsAutopilotDeploymentProfileDeviceType = "surfaceHub2S"
	WindowsAutopilotDeploymentProfileDeviceType_VirtualMachine WindowsAutopilotDeploymentProfileDeviceType = "virtualMachine"
	WindowsAutopilotDeploymentProfileDeviceType_WindowsPc      WindowsAutopilotDeploymentProfileDeviceType = "windowsPc"
)

func PossibleValuesForWindowsAutopilotDeploymentProfileDeviceType() []string {
	return []string{
		string(WindowsAutopilotDeploymentProfileDeviceType_HoloLens),
		string(WindowsAutopilotDeploymentProfileDeviceType_SurfaceHub2),
		string(WindowsAutopilotDeploymentProfileDeviceType_SurfaceHub2S),
		string(WindowsAutopilotDeploymentProfileDeviceType_VirtualMachine),
		string(WindowsAutopilotDeploymentProfileDeviceType_WindowsPc),
	}
}

func (s *WindowsAutopilotDeploymentProfileDeviceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAutopilotDeploymentProfileDeviceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAutopilotDeploymentProfileDeviceType(input string) (*WindowsAutopilotDeploymentProfileDeviceType, error) {
	vals := map[string]WindowsAutopilotDeploymentProfileDeviceType{
		"hololens":       WindowsAutopilotDeploymentProfileDeviceType_HoloLens,
		"surfacehub2":    WindowsAutopilotDeploymentProfileDeviceType_SurfaceHub2,
		"surfacehub2s":   WindowsAutopilotDeploymentProfileDeviceType_SurfaceHub2S,
		"virtualmachine": WindowsAutopilotDeploymentProfileDeviceType_VirtualMachine,
		"windowspc":      WindowsAutopilotDeploymentProfileDeviceType_WindowsPc,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeploymentProfileDeviceType(input)
	return &out, nil
}
