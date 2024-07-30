package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType string

const (
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_DefaultLimit                                          DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "defaultLimit"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions                           DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "defaultPlatformRestrictions"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "defaultWindows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness                        DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "defaultWindowsHelloForBusiness"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration              DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "deviceComanagementAuthorityConfiguration"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration                  DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "enrollmentNotificationsConfiguration"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_Limit                                                 DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "limit"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions                                  DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "platformRestrictions"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction                             DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "singlePlatformRestriction"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_Unknown                                               DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "unknown"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration        DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "windows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness                               DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "windowsHelloForBusiness"
)

func PossibleValuesForDeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType() []string {
	return []string{
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_DefaultLimit),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_Limit),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_Unknown),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness),
	}
}

func (s *DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType(input string) (*DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType, error) {
	vals := map[string]DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType{
		"defaultlimit":                DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_DefaultLimit,
		"defaultplatformrestrictions": DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions,
		"defaultwindows10enrollmentcompletionpageconfiguration": DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration,
		"defaultwindowshelloforbusiness":                        DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness,
		"devicecomanagementauthorityconfiguration":              DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration,
		"enrollmentnotificationsconfiguration":                  DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration,
		"limit":                                                 DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_Limit,
		"platformrestrictions":                                  DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions,
		"singleplatformrestriction":                             DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction,
		"unknown":                                               DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_Unknown,
		"windows10enrollmentcompletionpageconfiguration":        DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration,
		"windowshelloforbusiness":                               DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType(input)
	return &out, nil
}
