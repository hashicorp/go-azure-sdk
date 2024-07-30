package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus string

const (
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_HoloLens                DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "holoLens"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_HoloLensEnterprise      DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "holoLensEnterprise"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_HolographicForBusiness  DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "holographicForBusiness"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_Iot                     DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "iot"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_IotEnterprise           DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "iotEnterprise"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_SurfaceHub              DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "surfaceHub"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_Unknown                 DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "unknown"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsEducation        DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "windowsEducation"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsEnterprise       DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "windowsEnterprise"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsHome             DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "windowsHome"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsMobile           DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "windowsMobile"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsMobileEnterprise DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "windowsMobileEnterprise"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsMultiSession     DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "windowsMultiSession"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsProfessional     DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "windowsProfessional"
	DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsTeamSurface      DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus = "windowsTeamSurface"
)

func PossibleValuesForDeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus() []string {
	return []string{
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_HoloLens),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_HoloLensEnterprise),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_HolographicForBusiness),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_Iot),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_IotEnterprise),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_SurfaceHub),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_Unknown),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsEducation),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsEnterprise),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsHome),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsMobile),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsMobileEnterprise),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsMultiSession),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsProfessional),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsTeamSurface),
	}
}

func (s *DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus(input string) (*DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus, error) {
	vals := map[string]DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus{
		"hololens":                DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_HoloLens,
		"hololensenterprise":      DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_HoloLensEnterprise,
		"holographicforbusiness":  DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_HolographicForBusiness,
		"iot":                     DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_Iot,
		"iotenterprise":           DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_IotEnterprise,
		"surfacehub":              DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_SurfaceHub,
		"unknown":                 DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_Unknown,
		"windowseducation":        DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsEducation,
		"windowsenterprise":       DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsEnterprise,
		"windowshome":             DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsHome,
		"windowsmobile":           DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsMobile,
		"windowsmobileenterprise": DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsMobileEnterprise,
		"windowsmultisession":     DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsMultiSession,
		"windowsprofessional":     DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsProfessional,
		"windowsteamsurface":      DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus_WindowsTeamSurface,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus(input)
	return &out, nil
}
