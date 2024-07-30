package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType string

const (
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_DefaultLimit                                          DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "defaultLimit"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions                           DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "defaultPlatformRestrictions"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "defaultWindows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness                        DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "defaultWindowsHelloForBusiness"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration              DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "deviceComanagementAuthorityConfiguration"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration                  DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "enrollmentNotificationsConfiguration"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_Limit                                                 DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "limit"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions                                  DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "platformRestrictions"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction                             DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "singlePlatformRestriction"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_Unknown                                               DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "unknown"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration        DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "windows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness                               DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType = "windowsHelloForBusiness"
)

func PossibleValuesForDeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType() []string {
	return []string{
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_DefaultLimit),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_Limit),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_Unknown),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness),
	}
}

func (s *DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType(input string) (*DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType, error) {
	vals := map[string]DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType{
		"defaultlimit":                DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_DefaultLimit,
		"defaultplatformrestrictions": DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions,
		"defaultwindows10enrollmentcompletionpageconfiguration": DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration,
		"defaultwindowshelloforbusiness":                        DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness,
		"devicecomanagementauthorityconfiguration":              DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration,
		"enrollmentnotificationsconfiguration":                  DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration,
		"limit":                                                 DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_Limit,
		"platformrestrictions":                                  DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions,
		"singleplatformrestriction":                             DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction,
		"unknown":                                               DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_Unknown,
		"windows10enrollmentcompletionpageconfiguration":        DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration,
		"windowshelloforbusiness":                               DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentPlatformRestrictionConfigurationDeviceEnrollmentConfigurationType(input)
	return &out, nil
}
