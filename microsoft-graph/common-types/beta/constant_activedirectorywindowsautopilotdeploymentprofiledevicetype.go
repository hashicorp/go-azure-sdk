package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType string

const (
	ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType_HoloLens       ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType = "holoLens"
	ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType_SurfaceHub2    ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType = "surfaceHub2"
	ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType_SurfaceHub2S   ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType = "surfaceHub2S"
	ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType_VirtualMachine ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType = "virtualMachine"
	ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType_WindowsPc      ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType = "windowsPc"
)

func PossibleValuesForActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType() []string {
	return []string{
		string(ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType_HoloLens),
		string(ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType_SurfaceHub2),
		string(ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType_SurfaceHub2S),
		string(ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType_VirtualMachine),
		string(ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType_WindowsPc),
	}
}

func (s *ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType(input string) (*ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType, error) {
	vals := map[string]ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType{
		"hololens":       ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType_HoloLens,
		"surfacehub2":    ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType_SurfaceHub2,
		"surfacehub2s":   ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType_SurfaceHub2S,
		"virtualmachine": ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType_VirtualMachine,
		"windowspc":      ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType_WindowsPc,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActiveDirectoryWindowsAutopilotDeploymentProfileDeviceType(input)
	return &out, nil
}
