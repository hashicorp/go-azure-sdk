package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus string

const (
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus_HardwareRequirementsNotMet      WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus = "hardwareRequirementsNotMet"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus_HoloLensProfileNotSupported     WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus = "holoLensProfileNotSupported"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus_None                            WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus = "none"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus_SurfaceHub2SProfileNotSupported WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus = "surfaceHub2SProfileNotSupported"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus_SurfaceHubProfileNotSupported   WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus = "surfaceHubProfileNotSupported"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus_WindowsPcProfileNotSupported    WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus = "windowsPcProfileNotSupported"
)

func PossibleValuesForWindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus() []string {
	return []string{
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus_HardwareRequirementsNotMet),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus_HoloLensProfileNotSupported),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus_None),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus_SurfaceHub2SProfileNotSupported),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus_SurfaceHubProfileNotSupported),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus_WindowsPcProfileNotSupported),
	}
}

func (s *WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus(input string) (*WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus, error) {
	vals := map[string]WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus{
		"hardwarerequirementsnotmet":      WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus_HardwareRequirementsNotMet,
		"hololensprofilenotsupported":     WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus_HoloLensProfileNotSupported,
		"none":                            WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus_None,
		"surfacehub2sprofilenotsupported": WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus_SurfaceHub2SProfileNotSupported,
		"surfacehubprofilenotsupported":   WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus_SurfaceHubProfileNotSupported,
		"windowspcprofilenotsupported":    WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus_WindowsPcProfileNotSupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus(input)
	return &out, nil
}
