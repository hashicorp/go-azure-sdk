package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType string

const (
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_DefaultLimit                                          DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "defaultLimit"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions                           DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "defaultPlatformRestrictions"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "defaultWindows10EnrollmentCompletionPageConfiguration"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness                        DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "defaultWindowsHelloForBusiness"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration              DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "deviceComanagementAuthorityConfiguration"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration                  DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "enrollmentNotificationsConfiguration"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_Limit                                                 DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "limit"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions                                  DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "platformRestrictions"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction                             DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "singlePlatformRestriction"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_Unknown                                               DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "unknown"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration        DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "windows10EnrollmentCompletionPageConfiguration"
	DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness                               DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType = "windowsHelloForBusiness"
)

func PossibleValuesForDeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType() []string {
	return []string{
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_DefaultLimit),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_Limit),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_Unknown),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration),
		string(DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness),
	}
}

func (s *DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType(input string) (*DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType, error) {
	vals := map[string]DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType{
		"defaultlimit":                DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_DefaultLimit,
		"defaultplatformrestrictions": DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_DefaultPlatformRestrictions,
		"defaultwindows10enrollmentcompletionpageconfiguration": DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration,
		"defaultwindowshelloforbusiness":                        DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness,
		"devicecomanagementauthorityconfiguration":              DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration,
		"enrollmentnotificationsconfiguration":                  DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration,
		"limit":                                                 DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_Limit,
		"platformrestrictions":                                  DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_PlatformRestrictions,
		"singleplatformrestriction":                             DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_SinglePlatformRestriction,
		"unknown":                                               DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_Unknown,
		"windows10enrollmentcompletionpageconfiguration":        DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration,
		"windowshelloforbusiness":                               DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType_WindowsHelloForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComanagementAuthorityConfigurationDeviceEnrollmentConfigurationType(input)
	return &out, nil
}
