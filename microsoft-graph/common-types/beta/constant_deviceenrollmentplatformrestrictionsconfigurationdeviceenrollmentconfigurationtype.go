package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType string

const (
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_DefaultLimit                                          DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "defaultLimit"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions                           DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "defaultPlatformRestrictions"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "defaultWindows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness                        DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "defaultWindowsHelloForBusiness"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration              DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "deviceComanagementAuthorityConfiguration"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration                  DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "enrollmentNotificationsConfiguration"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_Limit                                                 DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "limit"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions                                  DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "platformRestrictions"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction                             DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "singlePlatformRestriction"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_Unknown                                               DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "unknown"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration        DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "windows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness                               DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType = "windowsHelloForBusiness"
)

func PossibleValuesForDeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType() []string {
	return []string{
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_DefaultLimit),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_Limit),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_Unknown),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness),
	}
}

func (s *DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType(input string) (*DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType, error) {
	vals := map[string]DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType{
		"defaultlimit":                DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_DefaultLimit,
		"defaultplatformrestrictions": DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions,
		"defaultwindows10enrollmentcompletionpageconfiguration": DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration,
		"defaultwindowshelloforbusiness":                        DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness,
		"devicecomanagementauthorityconfiguration":              DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration,
		"enrollmentnotificationsconfiguration":                  DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration,
		"limit":                                                 DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_Limit,
		"platformrestrictions":                                  DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions,
		"singleplatformrestriction":                             DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction,
		"unknown":                                               DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_Unknown,
		"windows10enrollmentcompletionpageconfiguration":        DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration,
		"windowshelloforbusiness":                               DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType(input)
	return &out, nil
}
