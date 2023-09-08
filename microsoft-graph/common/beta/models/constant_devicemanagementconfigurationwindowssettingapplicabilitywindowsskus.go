package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus string

const (
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkusholoLens                DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "HoloLens"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkusholoLensEnterprise      DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "HoloLensEnterprise"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkusholographicForBusiness  DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "HolographicForBusiness"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkusiot                     DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "Iot"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkusiotEnterprise           DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "IotEnterprise"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkussurfaceHub              DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "SurfaceHub"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkusunknown                 DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "Unknown"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsEducation        DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "WindowsEducation"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsEnterprise       DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "WindowsEnterprise"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsHome             DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "WindowsHome"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsMobile           DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "WindowsMobile"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsMobileEnterprise DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "WindowsMobileEnterprise"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsMultiSession     DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "WindowsMultiSession"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsProfessional     DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "WindowsProfessional"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsTeamSurface      DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "WindowsTeamSurface"
)

func PossibleValuesForDeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus() []string {
	return []string{
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkusholoLens),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkusholoLensEnterprise),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkusholographicForBusiness),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkusiot),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkusiotEnterprise),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkussurfaceHub),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkusunknown),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsEducation),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsEnterprise),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsHome),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsMobile),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsMobileEnterprise),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsMultiSession),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsProfessional),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsTeamSurface),
	}
}

func parseDeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus(input string) (*DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus, error) {
	vals := map[string]DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus{
		"hololens":                DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkusholoLens,
		"hololensenterprise":      DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkusholoLensEnterprise,
		"holographicforbusiness":  DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkusholographicForBusiness,
		"iot":                     DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkusiot,
		"iotenterprise":           DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkusiotEnterprise,
		"surfacehub":              DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkussurfaceHub,
		"unknown":                 DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkusunknown,
		"windowseducation":        DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsEducation,
		"windowsenterprise":       DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsEnterprise,
		"windowshome":             DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsHome,
		"windowsmobile":           DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsMobile,
		"windowsmobileenterprise": DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsMobileEnterprise,
		"windowsmultisession":     DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsMultiSession,
		"windowsprofessional":     DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsProfessional,
		"windowsteamsurface":      DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkuswindowsTeamSurface,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus(input)
	return &out, nil
}
