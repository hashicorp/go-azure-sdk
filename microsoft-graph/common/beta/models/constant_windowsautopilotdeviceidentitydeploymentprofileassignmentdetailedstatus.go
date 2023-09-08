package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus string

const (
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatushardwareRequirementsNotMet      WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus = "HardwareRequirementsNotMet"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatusholoLensProfileNotSupported     WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus = "HoloLensProfileNotSupported"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatusnone                            WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus = "None"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatussurfaceHub2SProfileNotSupported WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus = "SurfaceHub2SProfileNotSupported"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatussurfaceHubProfileNotSupported   WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus = "SurfaceHubProfileNotSupported"
	WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatuswindowsPcProfileNotSupported    WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus = "WindowsPcProfileNotSupported"
)

func PossibleValuesForWindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus() []string {
	return []string{
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatushardwareRequirementsNotMet),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatusholoLensProfileNotSupported),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatusnone),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatussurfaceHub2SProfileNotSupported),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatussurfaceHubProfileNotSupported),
		string(WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatuswindowsPcProfileNotSupported),
	}
}

func parseWindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus(input string) (*WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus, error) {
	vals := map[string]WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus{
		"hardwarerequirementsnotmet":      WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatushardwareRequirementsNotMet,
		"hololensprofilenotsupported":     WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatusholoLensProfileNotSupported,
		"none":                            WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatusnone,
		"surfacehub2sprofilenotsupported": WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatussurfaceHub2SProfileNotSupported,
		"surfacehubprofilenotsupported":   WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatussurfaceHubProfileNotSupported,
		"windowspcprofilenotsupported":    WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatuswindowsPcProfileNotSupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeviceIdentityDeploymentProfileAssignmentDetailedStatus(input)
	return &out, nil
}
